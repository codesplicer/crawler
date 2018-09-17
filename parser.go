package main

import (
	"io"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	mapset "github.com/deckarep/golang-set"
)

func parsePage(baseURL string, body io.Reader) mapset.Set {
	document, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		log.Fatal("Error loading HTTP response", err)
	}

	links := extractLinks(baseURL, document)
	return links
}

func extractLinks(baseURL string, doc *goquery.Document) mapset.Set {
	foundUrls := mapset.NewSet()

	if doc != nil {
		doc.Find("a").Each(func(index int, element *goquery.Selection) {
			href, exists := element.Attr("href")
			if exists {
				// We only want internal inks
				if strings.HasPrefix(href, "/") {
					foundUrls.Add(ResolveAbsoluteURL(baseURL, href))
				}
			}
		})
	}
	return foundUrls
}
