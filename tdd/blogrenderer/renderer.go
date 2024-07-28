package blogrenderer

import (
	"embed"
	"html/template"
	"io"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

// const (
// 	postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
// )

var (
	postTemplates embed.FS
)

type PostRenderer struct {
	templ    *template.Template
	mdParser *parser.Parser
}

func NewPostenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	parser := parser.NewWithExtensions(extensions)
	return &PostRenderer{templ: templ, mdParser: parser}, nil
}

// func Render(w io.Writer, p Post) error {
// 	// _, err := fmt.Fprintf(w, "<h1>%s</h1>", p.Title)

// 	// return err

// 	// _, err := fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", p.Title, p.Description)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// _, err = fmt.Fprintf(w, "Tags: <ul>")
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	// for _, tag := range p.Tags {
// 	// 	_, err = fmt.Fprintf(w, "<li>%s</li>", tag)
// 	// 	if err != nil {
// 	// 		return err
// 	// 	}
// 	// }
// 	// _, err = fmt.Fprintf(w, "</ul>")
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	// return nil

// 	// temp, err := template.New("blog").Parse(postTemplate)
// 	temp, err := template.ParseFS(postTemplates, "templates/*.gohtml")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return err
// 	}

// 	if err := temp.ExecuteTemplate(w, "blog.gohtml", p); err != nil {
// 		fmt.Println("Error2:", err)
// 		return err
// 	}
// 	return nil
// }

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", newPostVM(p, r))
}

// func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
// 	indexTemplate := `<ol>{{range .}}<li><a href="/post/{{sanitiseTitle .Title}}">{{.Title}}</a></li>{{end}}</ol>`

// 	templ, err := template.New("index").Funcs(template.FuncMap{
// 		"sanitiseTitle": func(title string) string {
// 			return strings.ToLower(strings.Replace(title, " ", "-", -1))
// 		},
// 	}).Parse(indexTemplate)
// 	if err != nil {
// 		return err
// 	}

// 	if err := templ.Execute(w, posts); err != nil {
// 		return err
// 	}

// 	return nil
// }

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	// indexTemplate := `<ol>{{range .}}<li><a href="/post/{{.SanitisedTitle}}">{{.Title}}</a></li>{{end}}</ol>`

	// templ, err := template.New("index").Parse(indexTemplate)
	// if err != nil {
	// 	return err
	// }

	// if err := templ.Execute(w, posts); err != nil {
	// 	return err
	// }

	// return nil

	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}

type postViewModel struct {
	Post
	HTMLBody template.HTML
}

func newPostVM(p Post, r *PostRenderer) postViewModel {
	vm := postViewModel{Post: p}
	vm.HTMLBody = template.HTML(markdown.ToHTML([]byte(p.Body), r.mdParser, nil))
	return vm
}
