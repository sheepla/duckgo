package client

import (
	"encoding/json"
	"testing"
)

var param = &SearchParam{
	Query: "golang",
	Page:  1,
}

func TestBuildRequest(t *testing.T) {
	req, err := param.buildRequest()
	if err != nil {
		t.Fatal(err)
	}

	url := req.URL.String()
	if url != `https://html.duckduckgo.com/html?api=%2Fd.js&o=json&q=golang&s=dc&v=1` {
		t.Fatal(url)
	}
}

func TestSearch(t *testing.T) {
    result, err := Search(param)
    if err != nil {
        t.Fatal(err)
    }

    j, err := json.Marshal(result)
    if err != nil {
        panic(err)
    }

    t.Log(string(j))
}
