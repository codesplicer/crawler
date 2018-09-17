package main

import (
	"fmt"
	"log"
	"net/url"
)

// ParseBaseURL parse abseUrl
func ParseBaseURL(u string) string {
	parsed, _ := url.Parse(u)
	return fmt.Sprintf("%s://%s", parsed.Scheme, parsed.Host)
}

// ResolveAbsoluteURL converts a relative path to an absolute URL
func ResolveAbsoluteURL(baseURL, target string) string {
	base, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	endpoint, err := url.Parse(target)
	if err != nil {
		log.Fatal(err)
	}

	// Strip URI fragments
	endpoint.Fragment = ""

	return base.ResolveReference(endpoint).String()
}
