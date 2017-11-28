package main

import (
	"fmt"
	"html/template"
	"web/routes"

	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"web/session"
)

func unescape(x string) interface{} {
	return template.HTML(x)
}

func main() {

	fmt.Println("Listening on port: 3000")

	mongoSession, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	db := mongoSession.DB("blog")

	m := martini.Classic()

	unescapeFuncMap := template.FuncMap{"unescape": unescape}

	m.Map(db)

	m.Use(session.Middleware)

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

	m.Get("/", routes.IndexHendler)
	m.Get("/write", routes.WriteHendler)
	m.Get("/edit/:id", routes.EditHendler)
	m.Get("/delete/:id", routes.DeleteHendler)
	m.Post("/SavePost", routes.SavePostHendler)
	m.Post("/gethtml", routes.GetHtmlHendler)
	m.Get("/login", routes.GetLoginHendler)
	m.Post("/login", routes.PostLoginHendler)

	m.Get("/test", func() string {
		return "test"
	})

	m.Run()

}
