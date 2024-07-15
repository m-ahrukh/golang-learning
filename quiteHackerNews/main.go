package main

import (
	"errors"
	"flag"
	"fmt"
	"goLangLearning/quiteHackerNews/hn"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	// parse flags
	var port, numStories int
	flag.IntVar(&port, "port", 3000, "the port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "the number of top stories to display")
	flag.Parse()

	tpl := template.Must(template.ParseFiles("./index.gohtml"))

	http.HandleFunc("/", handler(numStories, tpl))

	// Start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handler(numStories int, tpl *template.Template) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		stories, err := getTopStories(numStories)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data := templateData{
			Stories: stories,
			Time:    time.Now().Sub(start),
		}
		err = tpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Failed to process the template", http.StatusInternalServerError)
			return
		}

	})
}

func getTopStories(numStories int) ([]item, error) {
	var client hn.Client
	ids, err := client.TopItems()
	if err != nil {
		return nil, errors.New("failed to load storioes")
	}
	var stories []item
	at := 0
	for len(stories) < numStories {
		need := (numStories - len(stories))
		stories = append(stories, getStories(ids[at:at+need])...)
		at += need
	}
	return stories[:numStories], nil
}

func getStories(ids []int) []item {

	var client hn.Client

	type result struct {
		index int
		item  item
		err   error
	}
	results := make(chan result)
	for i := 0; i < len(ids); i++ {
		go func(index int, id int) {
			hnItem, err := client.GetItem(id)
			if err != nil {
				results <- result{index, item{}, err}
				return
			}
			item := parseHNItem(hnItem)
			results <- result{index, item, nil}
		}(i, ids[i])
	}
	stories := make([]item, len(ids))
	for i := 0; i < len(ids); i++ {
		res := <-results
		if res.err == nil && isStoryLink(res.item) {
			stories[i] = res.item
		}
	}
	return stories
}

func isStoryLink(item item) bool {
	return item.Type == "story" && item.URL != ""
}

func parseHNItem(hnItem hn.Item) item {
	ret := item{Item: hnItem}
	url, err := url.Parse(ret.URL)
	if err == nil {
		ret.Host = strings.TrimPrefix(url.Hostname(), "www.")
	}
	return ret
}

// item is the same as the hn.Item, but adds the Host field
type item struct {
	hn.Item
	Host string
}

type templateData struct {
	Stories []item
	Time    time.Duration
}
