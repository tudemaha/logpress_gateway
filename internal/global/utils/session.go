package utils

import (
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

var (
	key   = securecookie.GenerateRandomKey(32)
	store = sessions.NewCookieStore(key)
)

func CreateSession(w http.ResponseWriter, r *http.Request, username string) {
	store.Options = &sessions.Options{
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
	}

	session, _ := store.Get(r, "logpress")

	session.Values["authenticated"] = true
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
