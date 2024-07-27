package blogrenderer

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRennder(t *testing.T) {
	var (
		aPost = Post{
			Title:       "Hello World",
			Body:        "This is a post",
			Description: "This is a description",
			Tags:        []string{"go", "tdd"},
		}
	)

	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>Hello World</h1>
<p>This is a description</p>
Tags: <ul><li>go</li><li>tdd</li></ul>`
		want = strings.ReplaceAll(want, "\n", "")
		assert.Equal(t, want, got)
	})
}
