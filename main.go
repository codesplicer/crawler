package main

import (
	"log"
	"os"

	"github.com/deckarep/golang-set"
)

// CrawledPages = Stores a set containing all seen URLs
// TODO - add a mutex lock when making this concurrent
var CrawledPages = mapset.NewSet()

// ToCrawl Stores all URLs awaiting crawling
var ToCrawl = mapset.NewSet()
var mySitemap *Sitemap

func main() {
	if len(os.Args) < 2 {
		log.Println("USAGE: ./crawler [https://monzo.com]")
		os.Exit(1)
	}
	target := os.Args[1]
	mySitemap = NewSitemap()

	//create a new instance of the crawler structure
	c := Crawler{
		target,
		make(chan string, 512),
		make(chan SitemapRecord, 512),
		mapset.NewSet(),
		mySitemap,
	}

	// Start the crawler
	c.ToCrawl <- c.baseURL
	c.Start()
}
