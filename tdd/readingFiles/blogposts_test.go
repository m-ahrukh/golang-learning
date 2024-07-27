package blogposts

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestNewBlogPosts(t *testing.T) {

	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}
	posts, err := NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("\n\tgot %d posts\n\twanted %d posts", len(posts), len(fs))
	}

	got := posts[0]
	want := Post{Title: "Post 1", Description: "Description 1", Tags: []string{"tdd", "go"}, Body: `Hello
World`}
	assert.Equal(t, want, got)
}
