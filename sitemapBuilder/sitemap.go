package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

// TODO:
//     1. get page
//     2. parse the html of that page and extract links
//     3. build proper urls
//     4. filter liks with different domains
//     5. find pages
//     6. print xml

func main() {
	fmt.Println("Sitemap Builder Problem")

	url := flag.String("url", "https://gophercises.com/", "the url that you want to build a sitemap for")
	flag.Parse()

	fmt.Println("url:", *url)

	response, err := http.Get(*url)

	if err != nil {
		fmt.Println("Error:", err)
	}

	defer response.Body.Close()
	io.Copy(os.Stdout, response.Body) //print the html of the webpage
}
