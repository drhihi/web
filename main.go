package main

import (
	"fmt"
	"net/http"
	"html/template"

	"web/db/documents"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	_"log"
)



var postsCollection *mgo.Collection

func indexHendler(rnd render.Render) {

	postsDocuments := []documents.PostDocument{}
	postsCollection.Find(nil).All(&postsDocuments)

	rnd.HTML(200, "index", postsDocuments)
}

func writeHendler(rnd render.Render) {
	postDocument := documents.PostDocument{}
	rnd.HTML(200, "write", postDocument)
}

func editHendler(rnd render.Render, params martini.Params) {
	id := params["id"]
	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}

	rnd.HTML(200, "write", postDocument)
}

func savePostHendler(rnd render.Render, r *http.Request) {
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("content")
	contentHtml := convertMarkdownToHtml(contentMarkdown)

	postDocument := documents.PostDocument{
		id,
		title,
		contentHtml,
		contentMarkdown,
	}

	if id != "" {
		postsCollection.UpdateId(id, postDocument)
	} else {
		id = generateId()
		postDocument.Id = id
		postsCollection.Insert(postDocument)
	}

	rnd.Redirect("/", 302)
}

func deleteHendler(rnd render.Render, params martini.Params) {
	id := params["id"]
	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}

	postsCollection.RemoveId(id)

	rnd.Redirect("/")
}

func getHtmlHendler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")
	html := convertMarkdownToHtml(md)
	rnd.JSON(200, map[string]interface{}{"html": html})
}

func unescape(x string) interface{} {
	return template.HTML(x)
}

func main() {

	fmt.Println("Listening on port: 3000")

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	postsCollection = session.DB("blog").C("posts")

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
