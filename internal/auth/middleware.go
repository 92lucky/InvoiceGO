package auth

import (
	"net/http"

	"github.com/gorilla/sessions"
)

func RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "session")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			returnTo := r.URL.Path
			http.Redirect(w, r, "/login?returnTo="+returnTo, http.StatusSeeOther)
			return
		}
		next(w, r)
	}
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	return Store.Get(r, "session")
}


