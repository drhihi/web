package main

import (
	"fmt"
	"net/http"
	"html/template"

	"web/models"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"github.com/russross/blackfriday"
)

var posts map[string]*models.Post

func indexHendler(rnd render.Render) {
	rnd.HTML(200, "index", posts)
}

func writeHendler(rnd render.Render) {
	rnd.HTML(200, "write", nil)
}

func editHendler(rnd render.Render, params martini.Params) {
	id := params["id"]
	p, ok := posts[id]
	if !ok {
		rnd.Redirect("/")
		return
	}
	rnd.HTML(200, "write", p)
}

func savePostHendler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("content")
	contentHtml := string(blackfriday.MarkdownBasic([]byte(contentMarkdown)))

	var post *models.Post
	if id != "" {
		post = posts[id]
		post.Title = title
		post.ContentHtml = contentHtml
		post.ContentMarkdown = contentMarkdown
	} else {
		id = generateId()
		post := models.NewPost(id, title, contentHtml, contentMarkdown)
		posts[post.Id] = post
	}

	rnd.Redirect("/", 302)
}

func deleteHendler(rnd render.Render, params martini.Params) {
	id := params["id"]
	_, ok := posts[id]
	if !ok {
		rnd.Redirect("/", 404)
		return
	}
	delete(posts, id)

	rnd.Redirect("/", 302)
}

func getHtmlHendler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")
	htmlBytes := blackfriday.MarkdownBasic([]byte(md))
	rnd.JSON(200, map[string]interface{}{"html": string(htmlBytes)})
}

func unescape(x string) interface{} {
	return template.HTML(x)
}

func main() {

	fmt.Println("Listening on port: 3000")

	posts = make(map[string]*models.Post, 0)

	m := martini.Classic()

	unescapeFuncMap := template.FuncMap{"unescape": unescape}

	m.Use(render.Renderer(render.Options{
		Directory:  "templates",
		Layout:     "layout",
		Extensions: []string{".tmpl", ".html"},
		Funcs:      []template.FuncMap{unescapeFuncMap},
		Charset:    "UTF-8",
		IndentJSON: true,
	}))

	staticOptions := martini.StaticOptions{Prefix: "assets"}
	m.Use(martini.Static("assets", staticOptions))

	m.Get("/", indexHendler)
	m.Get("/write", writeHendler)
	m.Get("/edit/:id", editHendler)
	m.Get("/delete/:id", deleteHendler)
	m.Post("/SavePost", savePostHendler)
	m.Post("/gethtml", getHtmlHendler)

	m.Get("/test", func() string {
		return "test"
	})

	m.Run()

}
