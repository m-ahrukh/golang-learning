package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

var htmlCode = `<html>
    <head>
        <title>test</title>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-QWTKZyjpPEjISv5WaRU9OFeRpok6YctnYmDr5pNlyT2bRjXh0JMhjY6hW+ALEwIH" crossorigin="anonymous">
    </head>
    <body>
        <h1>Links</h1>
		<div>
			<a href="/dog">
			<span>Something in a span </span> 
			Text not in a span 
			<b>Bold text!</b>
			</a>
			<a href="https://www.google.com/">you can search anything on <strong>Google</strong></a>
			<a href="https://web.facebook.com/?_rdc=1&_rdr">contact me on <i class="fa fa-facebook" aria-hidden="true"></i></a>
		</div>
    </body>
</html>`

type Link struct {
	Href string
	Text string
}

func newLinkNode() Link {
	link := Link{
		Href: "",
		Text: "",
	}
	return link
}

func main() {
	doc, err := html.Parse(strings.NewReader(htmlCode))
	if err != nil {
		log.Fatal(err)
	}

	// dfs(doc, "")

	var links []Link

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			href := ""
			for _, a := range n.Attr {
				if a.Key == "href" {
					href = a.Val
					break
				}
			}
			textContent := getTextContent(n)

			link := newLinkNode()
			link.Href = href
			link.Text = textContent

			links = append(links, link)

			// fmt.Printf("href: %s, text: %s\n", href, textContent)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	fmt.Println("Links {")
	for _, link := range links {
		fmt.Printf("\tHref: %s,\n\tText: %s,\n\n", link.Href, link.Text)
	}
	fmt.Println("}")
}

func getTextContent(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	var text string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		text += getTextContent(c)
	}
	return strings.Join(strings.Fields(text), " ")
	// return text
}
