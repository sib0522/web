package models

import (
	"github.com/gorilla/sessions"
	"os"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("secret-key")))

type SessionStore struct {
	Store *sessions.CookieStore
}

func NewSessionStore() *SessionStore {
	return &SessionStore{Store: sessions.NewCookieStore([]byte("secret-key"))}
}
