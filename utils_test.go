package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBaseURL(t *testing.T) {
	assert.Equal(t, "https://google.com", ParseBaseURL("https://google.com?abc=123"))
	assert.Equal(t, "https://monzo.com", ParseBaseURL("https://monzo.com/foobar"))
	assert.Equal(t, "https://monzo.com", ParseBaseURL("https://monzo.com/"))
}

func TestResolveAbsoluteURL(t *testing.T) {
	assert.Equal(t, "https://google.com/foo/bar", ResolveAbsoluteURL("https://google.com", "/foo/bar"))
	assert.Equal(t, "https://monzo.com/cdn-cgi/l/email-protection", ResolveAbsoluteURL("https://monzo.com", "/cdn-cgi/l/email-protection#abc3cec7dbebc6c4c5d1c485c8c4c6"))
}
