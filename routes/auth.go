package routes

import (
	"github.com/martini-contrib/render"
	"net/http"
	"log"
	"web/session"
)

func GetLoginHendler(rnd render.Render) {
	rnd.HTML(200, "login", nil)
}

func PostLoginHendler(rnd render.Render, r *http.Request, s *session.Session) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Println(username, password)

	s.Username = username

	rnd.Redirect("/")
}
