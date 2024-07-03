package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type Chapter struct {
	Title     string   `json:"title"`
	Paragraph []string `json:"story"`
	Options   []Option `json:"options"`
}

type Story map[string]Chapter

var defaultTemplate = `
<!DOCTYPE html>
<html>
	<head>
		<meta charset="utf-8">
		<title>Choose Your Own Adventure</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		{{range .Paragraph}}
			<p>{{.}}</p>
		{{end}}
		<ul>
			{{range .Options}}
			<li><a href="/{{.Arc}}">{{.Text}}</a></li>
			{{end}}
		</ul>
	</body>
</html>`

func main() {
	fmt.Println("Choose Your Own Adventure ")

	jsonFile, err := os.Open("story.json")
	if err != nil {
		fmt.Println("error: ", err)
	}
	defer jsonFile.Close()

	story, err := jsonParser(jsonFile)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	handler := htmlHandler(story)
	fmt.Println("Starting server at port:", 3000)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 3000), handler))

}

func jsonParser(r io.Reader) (Story, error) {
	byteValue, _ := ioutil.ReadAll(r)
	var story Story
	err := json.Unmarshal(byteValue, &story)
	if err != nil {
		fmt.Println("Failed to parse json:", err)
		return nil, err
	}

	return story, nil
}

func htmlHandler(s Story) http.Handler {
	return handler{s}
}

type handler struct {
	story Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	htmlTemplate := template.Must(template.New("").Parse(defaultTemplate))

	err := htmlTemplate.Execute(w, h.story["intro"])
	if err != nil {
		fmt.Println("Error:", err)
	}
}
