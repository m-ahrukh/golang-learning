package blogrenderer

import (
	"bytes"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
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

	// t.Run("it converts a single post into HTML", func(t *testing.T) {
	// 	buf := bytes.Buffer{}
	// 	err := Render(&buf, aPost)

	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	// 		got := buf.String()
	// 	// 		want := `<h1>Hello World</h1>
	// 	// <p>This is a description</p>
	// 	// Tags: <ul><li>go</li><li>tdd</li></ul>`
	// 	// 		want = strings.ReplaceAll(want, "\n", "")
	// 	// 		assert.Equal(t, want, got)
	// 	approvals.VerifyString(t, buf.String())
	// })

	postRenderer, err := NewPostenderer()

	if err != nil {
		t.Fatal(err)
	}
	t.Run("it converts a single post into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := postRenderer.Render(&buf, aPost); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})
	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []Post{{Title: "Hello World"}, {Title: "Hello World 2"}}
		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}
		// got := buf.String()
		// want := `<ol><li><a href="/post/hello-world">Hello World</a></li><li><a href="/post/hello-world-2">Hello World 2</a></li></ol>`

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }

		approvals.VerifyString(t, buf.String())
	})
}

// func BenchmarkRender(b *testing.B) {
// 	var (
// 		aPost = Post{
// 			Title:       "Hello World",
// 			Body:        "This is a post",
// 			Description: "This is a description",
// 			Tags:        []string{"go", "tdd"},
// 		}
// 	)

// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		Render(io.Discard, aPost)
// 	}
// }
