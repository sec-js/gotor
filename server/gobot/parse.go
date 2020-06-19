package gobot

import (
	"fmt"
	"io"
	"log"
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
func GetLinks(rootLink string) ([]LinkData, error) {
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

	linkDataList := make([]LinkData, 0)
	processLinks(client, links, func(r *http.Response, e error) {
		linkDataList = append(linkDataList, LinkData{
			Name:   r.Request.URL.String(),
			Status: e == nil && r.StatusCode < 400,
		})
	})
	return linkDataList, nil
}
