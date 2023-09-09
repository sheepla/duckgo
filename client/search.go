package client

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

const (
	defaultReferrer  = "https://google.com"
	defaultUserAgent = "Mozilla/5.0 (Windows NT 10.0; WOW64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.5666.197 Safari/537.36"
)

type SearchParam struct {
	Query string
}

func NewSearchParam(query string) (*SearchParam, error) {
	q := strings.TrimSpace(query)
	if q == "" {
		return nil, errors.New("search query is empty")
	}

	return &SearchParam{
		Query: q,
	}, nil
}

func (param *SearchParam) buildURL() (*url.URL, error) {
	u := &url.URL{
		Scheme: "https",
		Host:   "html.duckduckgo.com",
		Path:   "html"}
	q := u.Query()
	q.Add("q", param.Query)
	q.Add("v", "1")
	q.Add("o", "json")
	q.Add("api", "/d.js")
	u.RawQuery = q.Encode()

	return u, nil
}

func (param *SearchParam) buildRequest() (*http.Request, error) {
	u, err := param.buildURL()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return req, err
	}

	req.Header.Add("Referrer", defaultReferrer)
	req.Header.Add("User-Agent", defaultUserAgent)
	req.Header.Add("Cookie", "kl=wt-wt")
	req.Header.Add("Content-Type", "x-www-form-urlencoded")

	return req, nil
}

type SearchResult struct {
	Title   string
	Link    string
	Snippet string
}

func parse(r io.Reader) (*[]SearchResult, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	var (
		result []SearchResult
		item   SearchResult
	)
	doc.Find(".result").Each(func(i int, s *goquery.Selection) {
		item.Title = s.Find(".result__title a").Text()

		item.Link = extractLink(
			s.Find(".result__url").AttrOr("href", ""),
		)

		item.Snippet = removeHtmlTagsFromText(
			s.Find(".result__snippet").Text(),
		)

		result = append(result, item)
	})

	return &result, nil
}

func removeHtmlTags(node *html.Node, buf *bytes.Buffer) {
	if node.Type == html.TextNode {
		buf.WriteString(node.Data)
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		removeHtmlTags(child, buf)
	}
}

func removeHtmlTagsFromText(text string) string {
	node, err := html.Parse(strings.NewReader(text))
	if err != nil {
		// If it cannot be parsed text as HTML, return the text as is.
		return text
	}

	buf := &bytes.Buffer{}
	removeHtmlTags(node, buf)

	return buf.String()
}

// Extract target URL from href included in search result
// e.g.
//
//	`//duckduckgo.com/l/?uddg=https%3A%2F%2Fwww.vim8.org%2Fdownload.php&amp;rut=...`
//	                          ^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^^
//	                     --> `https://www.vim8.org/download.php`
func extractLink(href string) string {
	u, err := url.Parse(fmt.Sprintf("https:%s", strings.TrimSpace(href)))
	if err != nil {
		return ""
	}

	q := u.Query()
	if !q.Has("uddg") {
		return ""
	}

	return q.Get("uddg")
}

func Search(param *SearchParam) (*[]SearchResult, error) {
	c := &http.Client{
		Timeout: 7 * time.Second,
	}
	req, err := param.buildRequest()
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, err := parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
