package gobot

import (
	"fmt"
	"io"
	"log"
	"net/url"

	"golang.org/x/net/html"
)

// Link ...
type Link struct {
	Name   string
	Status bool
}

func closeConn(conn io.Closer) {
	err := conn.Close()
	if err != nil {
		log.Printf("Error: %+v", err)
	}
}

// Parses value to retrieve href
func parseLinks(attributes []html.Attribute) []string {
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

// GetLinks returns a map that contains the links as keys and their statuses as values
func GetLinks(rootLink string) (chan Link, error) {
	// Creating new Tor connection
	client := newDualClient(&ClientConfig{timeout: defaultTimeout})
	resp, err := client.Get(rootLink)
	if err != nil {
		return nil, err
	}
	defer closeConn(resp.Body)

	tokenizer := html.NewTokenizer(resp.Body)
	links := make([]string, 0)
	for notEnd := true; notEnd; {
		currentTokenType := tokenizer.Next()
		switch {
		case currentTokenType == html.ErrorToken:
			notEnd = false
		case currentTokenType == html.StartTagToken:
			token := tokenizer.Token()
			// Parsing and collecting href attribute values from anchor tags

			if token.Data == "a" {
				links = append(links, parseLinks(token.Attr)...)
			}
		}
	}

	if len(links) == 0 {
		return nil, fmt.Errorf("no links found for %s", rootLink)
	}

	linkChan := make(chan Link)
	for _, link := range links {
		go func(link string) {
			r, e := client.Head(link)
			linkChan <- Link{
				Name:   link,
				Status: e == nil && r.StatusCode < 400,
			}
		}(link)
	}
	return linkChan, nil
}
