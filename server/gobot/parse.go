package gobot

import (
	"errors"
	"net/http"
	"net/url"
	"sync"

	"golang.org/x/net/html"
)

// LinkData ...
type LinkData struct {
	Name   string
	Status bool
}

// Parses value to retrieve href
func parseHrefs(attributes []html.Attribute) []string {
	links := make([]string, 0)
	for i := 0; i < len(attributes); i++ {
		if attributes[i].Key == "href" {
			link, err := url.ParseRequestURI(attributes[i].Val)
			if link == nil || err != nil {
				continue
			}
			if link.Scheme != "" {
				links = append(links, link.String())
			}
		}
	}
	return links
}

type processFunction func(*http.Response, error)

func processLinks(c Client, links []string, process processFunction) {
	var wg sync.WaitGroup
	var mux sync.RWMutex
	for _, link := range links {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			resp, err := c.Head(l)
			mux.Lock()
			process(resp, err)
			mux.Unlock()
		}(link)
	}
	wg.Wait()
}

// GetLinks returns a map that contains the links as keys and their statuses as values
func GetLinks(searchURL string) ([]LinkData, error) {
	// Creating new Tor connection
	client := newDualClient(&ClientConfig{timeout: defaultTimeout})
	resp, err := client.Get(searchURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Begin parsing HTML
	tokenizer := html.NewTokenizer(resp.Body)
	links := make([]string, 0)
	for notEnd := true; notEnd; {
		currentTokenType := tokenizer.Next()
		switch {
		case currentTokenType == html.ErrorToken:
			notEnd = false
		case currentTokenType == html.StartTagToken:
			token := tokenizer.Token()
			// If link tag is found, append it to slice
			if token.Data == "a" {
				linksFound := parseHrefs(token.Attr)
				links = append(links, linksFound...)
			}
		}
	}

	if len(links) == 0 {
		return nil, errors.New("no links found for URL")
	}

	linkDataList := make([]LinkData, 0)
	processLinks(client, links, func(r *http.Response, e error) {
		linkDataList = append(linkDataList, LinkData{
			Name:   r.Request.URL.String(),
			Status: e == nil && r.StatusCode < 400,
		})
	})
	return linkDataList, nil
}
