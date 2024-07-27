package blogposts

import (
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}
	posts, err := NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}
	if len(posts) != len(fs) {
		t.Errorf("\n\tgot %d posts\n\twanted %d posts", len(posts), len(fs))
	}

	got := posts[0]
	want := Post{Title: "Post 1"}
	assert.Equal(t, want, got)
}
