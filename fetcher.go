package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/sethgrid/pester"
)

// getURL makes a request to the given URL and returns the raw response
// object.
func getURL(req *http.Request) (io.Reader, error) {
	client := pester.New()
	client.Concurrency = 1
	client.MaxRetries = 5
	client.Backoff = pester.ExponentialJitterBackoff
	client.KeepLog = true

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("[ERROR] GETing %s - %s", req.URL.String(), client.LogString())
		return nil, err
	}
	defer resp.Body.Close()

	// TODO - Replace this with an io.Copy call
	body, _ := ioutil.ReadAll(resp.Body)
	r := bytes.NewReader(body)
	return r, nil
}

// genRequest generates a HTTP request object based on the requestType
func genRequest(url string, requestType string) (*http.Request, error) {
	req, err := http.NewRequest(requestType, url, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return req, nil
}

func getHeaders(url string) io.Reader {
	req, _ := genRequest(url, "HEAD")
	res, _ := getURL(req)

	return res
}

func getPage(url string) io.Reader {
	req, _ := genRequest(url, "GET")
	res, _ := getURL(req)

	return res
}
