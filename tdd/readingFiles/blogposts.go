package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	// for range dir {
	// 	posts = append(posts, Post{})
	// }

	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()

	// postData, err := io.ReadAll(postFile)
	// if err != nil {
	// 	return Post{}, err
	// }

	// post := Post{Title: string(postData)}
	// return post, nil

	return newPost(postFile)
}

func newPost(postBody io.Reader) (Post, error) {
	// postData, err := io.ReadAll(postFile)
	// if err != nil {
	// 	return Post{}, err
	// }
	// post := Post{Title: string(postData)}
	// return post, nil

	scanner := bufio.NewScanner(postBody)

	// scanner.Scan()
	// titleLine := scanner.Text()

	// scanner.Scan()
	// descriptionLine := scanner.Text()

	// return Post{Title: titleLine[7:], Description: descriptionLine[13:]}, nil

	// readLine := func() string {
	// 	scanner.Scan()
	// 	return scanner.Text()
	// }

	// title := readLine()[len(titleSeparator):]
	// description := readLine()[len(descriptionSeparator):]

	// return Post{Title: title, Description: description}, nil

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readMetaLine(titleSeparator)
	description := readMetaLine(descriptionSeparator)
	tags := strings.Split(readMetaLine(tagsSeparator), ", ")

	// scanner.Scan()
	// buf := bytes.Buffer{}
	// for scanner.Scan() {
	// 	fmt.Fprintln(&buf, scanner.Text())
	// }
	// body := strings.TrimSuffix(buf.String(), "\n")

	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return strings.TrimSuffix(buf.String(), "\n")
}
