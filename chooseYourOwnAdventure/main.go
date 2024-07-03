package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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

func main() {
	fmt.Println("Choose Your Own Adventure ")

	jsonFile, err := os.Open("story.json")
	if err != nil {
		fmt.Println("erroe: ", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var story Story

	err = json.Unmarshal(byteValue, &story)

	if err != nil {
		fmt.Println("failed to parse json:", err)
	}

	for key, value := range story {
		fmt.Println("Chapter:", key, ":")
		fmt.Println("Chapter title:", value.Title)
		for _, para := range value.Paragraph {
			fmt.Println(para)
		}
		for _, option := range value.Options {
			fmt.Println(option)
		}
	}

}
