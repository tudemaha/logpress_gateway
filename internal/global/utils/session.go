package utils

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("secret-key")
	store = sessions.NewCookieStore(key)
)

func CreateSession(w http.ResponseWriter, r *http.Request, username string) {
	store.MaxAge(3600)
	session, _ := store.Get(r, "logpress")

	session.Values["autenticated"] = true
	session.Values["username"] = username
	session.Save(r, w)
}

func GetSession(w http.ResponseWriter, r *http.Request) (string, bool) {
	session, _ := store.Get(r, "logpress")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		return "", false
	}

	username, _ := session.Values["username"].(string)

	return username, true
}
