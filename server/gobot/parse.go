package gobot

import (
	"errors"
	"net/url"
	"sync"

	"golang.org/x/net/html"
)

// Link ...
type Link struct {
	Name   string
	Status bool
}

// Parses value to retrieve href
func parseHrefs(attributes []html.Attribute) []string {
	foundUrls := make([]string, 0)
	for i := 0; i < len(attributes); i++ {
		if attributes[i].Key == "href" {
			u, err := url.ParseRequestURI(attributes[i].Val)
			if u == nil || err != nil {
				continue
			}
			if u.Scheme != "" {
				foundUrls = append(foundUrls, u.String())
			}
		}
	}
	return foundUrls
}

// GetLinks returns a map that contains the links as keys and their statuses as values
func GetLinks(searchURL string) ([]Link, error) {
	// Creating new Tor connection
	client := newDualClient(&ClientConfig{timeout: defaultTimeout})
	resp, err := client.Get(searchURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Begin parsing HTML
	tokenizer := html.NewTokenizer(resp.Body)
	totalUrls := make([]string, 0)
	for notEnd := true; notEnd; {
		currentTokenType := tokenizer.Next()
		switch {
		case currentTokenType == html.ErrorToken:
			notEnd = false
		case currentTokenType == html.StartTagToken:
			token := tokenizer.Token()
			// If link tag is found, append it to slice
			if token.Data == "a" {
				urlsFound := parseHrefs(token.Attr)
				totalUrls = append(totalUrls, urlsFound...)
			}
		}
	}

	if len(totalUrls) == 0 {
		return nil, errors.New("no links found for URL")
	}

	links := make([]Link, 0)
	var wg sync.WaitGroup
	var mux sync.RWMutex
	var link Link
	for _, u := range totalUrls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resp, err := client.Head(u)
			mux.Lock()
			link = Link{
				Name:   resp.Request.URL.String(),
				Status: err == nil && resp.StatusCode < 400,
			}
			links = append(links, link)
			mux.Unlock()
		}(u)
	}
	wg.Wait()

	return links, nil
}
