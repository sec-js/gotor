package gobot

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	urllib "net/url"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/html"
)

// Client represents a GoBot HTTP client
type Client interface {
	Get(string) (*http.Response, error)
	Head(string) (*http.Response, error)
}

// ClientConfig contains configuration for a client
type ClientConfig struct {
	addr    string
	port    string
	timeout int
}

type dualClient struct {
	regClient *http.Client
	torClient *http.Client
}

const (
	defaultHost    = "127.0.0.1"
	defaultPort    = "9050"
	defaultTimeout = 60
)

// NewTorClient creates an HTTP client capable of performing TOR requests.
func newTorClient(config *ClientConfig) *dualClient {
	if config.addr == "" {
		config.addr = defaultHost
	}

	if config.port == "" {
		config.port = defaultPort
	}

	torProxyURL, err := urllib.Parse(fmt.Sprintf("socks5://%s:%s", config.addr, config.port))
	if err != nil {
		log.Fatal("Unable to parse Tor Proxy URL. Error:", err)
	}
	torTransport := &http.Transport{Proxy: http.ProxyURL(torProxyURL)}

	// Creating both clients for regular browsing and Tor
	torc := http.Client{Transport: torTransport, Timeout: time.Second * time.Duration(config.timeout)}
	regc := http.Client{Timeout: time.Second * time.Duration(config.timeout)}
	return &dualClient{regClient: &regc, torClient: &torc}

}

func (d *dualClient) Head(url string) (*http.Response, error) {
	if strings.Contains(url, ".onion") {
		return d.torClient.Head(url)
	}

	return d.regClient.Head(url)
}

func (d *dualClient) Get(url string) (*http.Response, error) {
	if strings.Contains(url, ".onion") {
		return d.torClient.Get(url)
	}

	return d.regClient.Get(url)
}

// Parses value to retrieve href
func findHrefs(attributes []html.Attribute) []string {
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
				urlsFound := findHrefs(token.Attr)
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
