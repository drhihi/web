package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"web/utils"
	"gopkg.in/mgo.v2"
	"github.com/codegangsta/martini"
	"web/db/documents"
)

func WriteHendler(rnd render.Render) {
	postDocument := documents.PostDocument{}
	rnd.HTML(200, "write", postDocument)
}

func EditHendler(rnd render.Render, params martini.Params, db *mgo.Database) {
	postsCollection := db.C("posts")
	id := params["id"]
	postDocument := documents.PostDocument{}
	err := postsCollection.FindId(id).One(&postDocument)
	if err != nil {
		rnd.Redirect("/")
		return
	}

	rnd.HTML(200, "write", postDocument)
}

func SavePostHendler(rnd render.Render, r *http.Request, db *mgo.Database) {
	postsCollection := db.C("posts")
	id := r.FormValue("id")
	title := r.FormValue("title")
	contentMarkdown := r.FormValue("content")
	contentHtml := utils.ConvertMarkdownToHtml(contentMarkdown)

	postDocument := documents.PostDocument{
		id,
		title,
		contentHtml,
		contentMarkdown,
	}

	if id != "" {
		postsCollection.UpdateId(id, postDocument)
	} else {
		id = utils.GenerateId()
		postDocument.Id = id
		postsCollection.Insert(postDocument)
	}

	rnd.Redirect("/", 302)
}

func DeleteHendler(rnd render.Render, params martini.Params, db *mgo.Database) {
	postsCollection := db.C("posts")
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

func GetHtmlHendler(rnd render.Render, r *http.Request) {
	md := r.FormValue("md")
	html := utils.ConvertMarkdownToHtml(md)
	rnd.JSON(200, map[string]interface{}{"html": html})
}
