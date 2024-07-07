package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	link "goLangLearning/sitemapBuilder/links"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// TODO:
//     1. get page ---> completed
//     2. parse the html of that page and extract links ---> done
//     3. build proper urls ---> done
//     4. filter liks with different domains ---> done
//     5. find pages ---> done
//     6. print xml

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	fmt.Println("Sitemap Builder Problem")

	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 10, "maximum depth of tree")
	flag.Parse()

	// fmt.Println("url:", *urlFlag, "depth:", *maxDepth)

	pages := bfs(*urlFlag, *maxDepth)
	// for _, page := range pages {
	// 	fmt.Println(page)
	// }

	//XML conversion
	toXml := urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}

	fmt.Println(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", " ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}

}

func bfs(urlString string, maxDepth int) []string {
	seen := make(map[string]struct{})
	var q map[string]struct{} //queue

	nq := map[string]struct{}{ //next queue
		urlString: struct{}{},
	}

	for i := 0; i < maxDepth; i++ {
		q, nq = nq, make(map[string]struct{})
		if len(q) == 0 {
			break
		}
		for url, _ := range q {
			if _, ok := seen[url]; ok {
				continue
			}
			seen[url] = struct{}{}

			for _, link := range get(url) {
				nq[link] = struct{}{}
			}
		}
	}
	ret := make([]string, 0, len(seen))
	for url, _ := range seen {
		ret = append(ret, url)
	}
	return ret
}

func get(urlString string) []string {
	response, err := http.Get(urlString)
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer response.Body.Close()
	// io.Copy(os.Stdout, response.Body) //print the html of the webpage
	reqUrl := response.Request.URL
	// fmt.Println(reqUrl.String())

	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(base, hrefs(response.Body, base), withPrefix(base))
}

func hrefs(body io.Reader, base string) []string {
	links, _ := link.Parse(body)

	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

func filter(base string, links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
