package main

import (
	"io"
	"log"
	"time"

	"github.com/deckarep/golang-set"
)

const requestTimeout = 30 * time.Second

// Crawler struct definition
type Crawler struct {
	// baseURL of the site being crawled
	baseURL string
	ToCrawl chan string
	results chan SitemapRecord
	visited mapset.Set
	sitemap *Sitemap
}

// Start crawler process
func (c *Crawler) Start() {
	log.Printf("Starting crawler")
	running := true

	for running {
		select {
		case url := <-c.ToCrawl:
			log.Printf("Enqueuing url=%v\n", url)
			go c.crawl(url)
		// Hacky - can we use a done channel instead?
		case <-time.Tick(5 * time.Second):
			log.Printf("[DEBUG] Crawled URLs: %d", c.visited.Cardinality())
			log.Printf("[DEBUG] To Crawl URLs: %d", len(c.ToCrawl))
			if len(c.ToCrawl) == 0 {
				c.sitemap.printMap()
				log.Println("No more work left")
				running = false
			}
		case result := <-c.results:
			log.Printf("Processing result")
			go c.processResults(result)
		}
	}
	// Print the sitemap
	c.sitemap.printMap()
}

func (c *Crawler) crawl(target string) {
	if c.visited.Contains(target) {
		log.Printf("[IGNORE]Ignoring previsited")
	} else {
		log.Printf("Making request to %s", target)
		req, _ := genRequest(target, "GET")
		body, _ := getURL(req)

		log.Printf("Crawled")
		c.visited.Add(target)
		c.extractLinks(target, body)
	}
}

func (c *Crawler) extractLinks(target string, body io.Reader) {
	log.Printf("Extracting links")
	extracted := parsePage(c.baseURL, body)
	result := SitemapRecord{
		parentURL: target,
		links:     extracted,
	}
	log.Printf("Adding to results queue")
	c.results <- result
}

func (c *Crawler) processResults(result SitemapRecord) {
	log.Println("Consuming results")
	c.sitemap.addRecord(result)

	// 2. Find which of the URLs haven't been visited
	nextUrls := result.links.Difference(c.visited)
	for link := range nextUrls.Iter() {
		if link, ok := link.(string); ok {
			log.Printf("Adding %s to crawler queue", link)
			c.ToCrawl <- link
		}
	}
}
