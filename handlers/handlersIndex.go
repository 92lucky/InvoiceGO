package handlers

import (
	"invoice-go/auth"
	"invoice-go/repository"
	"database/sql"
	"html/template"
	"net/http"
)

func HandleIndex(tmpl *template.Template, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := auth.GetSession(r)
		email := session.Values["email"].(string)

		_, err := repository.GetUserEmail(db, email)
		if err != nil {
			http.Redirect(w, r, "/setup", http.StatusSeeOther)
			return
		}
		tmpl.ExecuteTemplate(w, "index.html", nil)
	}
}
