package session

import (
	"web/utils"
	"github.com/codegangsta/martini"
	"net/http"
	"time"
)

const (
	COOKIE_NAME = "sessionId"
)

type Session struct {
	id       string
	Username string
}

type sessionStore struct {
	data map[string]*Session
}

func newSessionStore() *sessionStore {
	s := new(sessionStore)
	s.data = make(map[string]*Session)
	return s
}

func (s *sessionStore) get(sessionId string) *Session {
	session := s.data[sessionId]
	if session == nil {
		return &Session{id: sessionId}
	}
	return session
}

func (s *sessionStore) set(session *Session) {
	s.data[session.id] = session
}

func ensureCookie(r *http.Request, w http.ResponseWriter) string {
	cookie, _ := r.Cookie(COOKIE_NAME)
	if cookie != nil {
		return cookie.Value
	}

	sessionId := utils.GenerateId()
	cookie = &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   sessionId,
		Expires: time.Now().Add(5 * time.Minute),
	}
	http.SetCookie(w, cookie)

	return sessionId
}

var globSessionStore = newSessionStore()

func Middleware(ctx martini.Context, r *http.Request, w http.ResponseWriter) {
	sessionId := ensureCookie(r, w)
	session := globSessionStore.get(sessionId)
	ctx.Map(session)
	ctx.Next()
	globSessionStore.set(session)
}
