package client

import (
	"testing"
)

var param = &SearchParam{
	Query: "golang",
}

func TestBuildRequest(t *testing.T) {
	req, err := buildRequest(param, defaultClientOption)
	if err != nil {
		t.Fatal(err)
	}

	url := req.URL.String()
	if url != `https://html.duckduckgo.com/html?api=%2Fd.js&o=json&q=golang&s=dc&v=1` {
		t.Fatal(url)
	}
}

func TestSearch(t *testing.T) {
	_, err := Search(param)
	if err != nil {
		t.Fatal(err)
	}
}
