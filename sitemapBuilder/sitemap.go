package main

import (
	"flag"
	"fmt"
	link "goLangLearning/sitemapBuilder/links"
	"net/http"
	"net/url"
	"strings"
)

// TODO:
//     1. get page ---> completed
//     2. parse the html of that page and extract links ---> done
//     3. build proper urls
//     4. filter liks with different domains
//     5. find pages
//     6. print xml

func main() {
	fmt.Println("Sitemap Builder Problem")

	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	flag.Parse()

	fmt.Println("url:", *urlFlag)

	response, err := http.Get(*urlFlag)

	if err != nil {
		fmt.Println("Error:", err)
	}

	defer response.Body.Close()
	// io.Copy(os.Stdout, response.Body) //print the html of the webpage

	reqUrl := response.Request.URL
	fmt.Println(reqUrl.String())

	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	links, _ := link.Parse(response.Body)
	// for _, l := range links {
	// 	fmt.Println(l)
	// }

	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	for _, href := range hrefs {
		fmt.Println(href)
	}
}
