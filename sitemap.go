package main

import (
	"fmt"
	"log"
	"sync"

	mapset "github.com/deckarep/golang-set"
)

// SitemapRecord stores the results of a single pagecrawl
type SitemapRecord struct {
	// parentURL URL of the crawled page
	parentURL string
	// links all internal links discovered from the parent
	links mapset.Set
}

// Sitemap s
type Sitemap struct {
	sync.RWMutex
	links map[string]mapset.Set
}

// NewSitemap - Returns an initialised instance of Sitemap
func NewSitemap() *Sitemap {
	s := Sitemap{
		links: make(map[string]mapset.Set),
	}
	return &s
}

func (s *Sitemap) addRecord(c SitemapRecord) {
	log.Printf("Adding record to sitemap")
	s.Lock()
	s.addLinks(c.parentURL, c.links)
	log.Printf("There are %d items in the sitemap", len(s.links))
	s.Unlock()
}

func (s *Sitemap) addLinks(parentURL string, links mapset.Set) {
	s.links[parentURL] = s.getLinks(parentURL).Union(links)
}

func (s *Sitemap) getLinks(parentURL string) mapset.Set {
	value, exists := s.links[parentURL]
	if exists {
		return value
	}
	return mapset.NewSet()
}

func (s *Sitemap) printMap() {
	log.Printf("Going to print sitemap")
	for k := range s.links {
		fmt.Printf("============================\n")
		fmt.Printf("%s\n", k)

		for url := range s.links[k].Iter() {
			if url, ok := url.(string); ok {
				fmt.Printf("\t |->> %s\n", url)
			}
		}
		fmt.Printf("----------------------------\n\n")
	}
}
