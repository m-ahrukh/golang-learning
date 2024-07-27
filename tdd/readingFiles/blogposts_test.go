package blogposts

import (
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello-world2.md": {Data: []byte("hola")},
	}
	posts := NewPostsFromFS(fs)
	if len(posts) != len(fs) {
		t.Errorf("\n\tgot %d posts\n\twanted %d posts", len(posts), len(fs))
	}
}
