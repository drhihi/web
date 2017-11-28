package routes

import (
	"github.com/martini-contrib/render"
	"log"
	"gopkg.in/mgo.v2"
	"web/session"
	"web/db/documents"
)

func IndexHendler(rnd render.Render, s *session.Session, db *mgo.Database) {

	log.Println(s.Username)

	postsCollection := db.C("posts")

	postsDocuments := []documents.PostDocument{}
	postsCollection.Find(nil).All(&postsDocuments)

	rnd.HTML(200, "index", postsDocuments)
}
