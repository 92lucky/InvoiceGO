package handlers

import (
	"database/sql"
	"html/template"
	"invoice-go/internal/auth"
	"invoice-go/internal/repository"
	"net/http"
)

func HandleHome(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "home.html", nil)
	}
}



func HandleIndex(tmpl *template.Template, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := auth.GetSession(r)
		email := session.Values["email"].(string)

		_, err := repository.GetUserEmail(db, email)
		if err != nil {
			http.Redirect(w, r, "/setup", http.StatusSeeOther)
			return
		}
		tmpl.ExecuteTemplate(w, "index.html", r)
	}
}