package link

import (
	"io"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/net/html"
)

// Link is an html link
type Link struct {
	Href string
	Text string
}

// Parse will return a slice of links
func Parse(r io.Reader) (links []Link, err error) {
	p := bluemonday.NewPolicy()
	p.AllowAttrs("href").OnElements("a")
	san := p.SanitizeReader(r)

	doc, err := html.Parse(san)
	if err != nil {
		return nil, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {

		if n.Type == html.ElementNode && n.Data == "a" {
			href := getHrefValue(n.Attr)

			link := Link{
				Href: href,
				Text: strings.TrimSpace(n.FirstChild.Data),
			}
			links = append(links, link)
			return
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)
	return links, nil
}

func getHrefValue(attributes []html.Attribute) string {
	for _, a := range attributes {
		if a.Key == "href" {
			return a.Val
		}
	}
	return ""
}
