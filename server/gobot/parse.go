package gobot

import (
	"errors"
	urllib "net/url"
	"sync"

	"golang.org/x/net/html"
)

const (
	defaultHost    = "127.0.0.1"
	defaultPort    = "9050"
	defaultTimeout = 60
)

// Parses value to retrieve href
func parseHrefs(attributes []html.Attribute) []string {
	foundUrls := make([]string, 0)
	for i := 0; i < len(attributes); i++ {
		if attributes[i].Key == "href" {
			url, err := urllib.ParseRequestURI(attributes[i].Val)
			if url == nil || err != nil {
				continue
			}
			if url.Scheme != "" {
				foundUrls = append(foundUrls, url.String())
			}
		}
	}
	return foundUrls
}

// GetLinks returns a map that contains the links as keys and their statuses as values
func GetLinks(searchURL string) (map[string]bool, error) {
	// Creating new Tor connection
	client := newTorClient(&ClientConfig{timeout: defaultTimeout})
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

	// Check all links and assign their status
	linksWithStatus := make(map[string]bool)
	var wg sync.WaitGroup
	for _, url := range totalUrls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			resp, err := client.Head(url)
			linksWithStatus[url] = err == nil && resp.StatusCode < 400
		}(url)
	}
	wg.Wait()

	return linksWithStatus, nil
}
