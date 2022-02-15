package main

import (
	"github.com/dnaeon/go-vcr/recorder"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestSimple(t *testing.T){

	// Start the recorder
	r, err := recorder.New("fixtures/golang-org")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Stop()

	// Create an HTTP client and inject our transport
	client := &http.Client{
		Transport: r, // Inject as transport
	}

	url := "http://golang.org/"
	resp, err := client.Get(url)
	if err != nil {
		t.Fatalf("Failed to get url %s: %s", url, err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %s", err)
	}

	wantTitle := "<title>The Go Programming Language</title>"
	bodyContent := string(body)

	if !strings.Contains(bodyContent, wantTitle){
		t.Errorf("Title %s not found in response.", wantTitle)
	}






}
