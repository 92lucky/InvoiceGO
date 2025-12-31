package handlers

import (
	"database/sql"
	"html/template"
	"invoice-go/internal/auth"
	"invoice-go/internal/repository"
	"invoice-go/internal/service"
	"net/http"
)

func HandlersInvoice(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data, err := service.ServiceInvoice(r)
			if err != nil {
				http.Error(w, "Gagal memproses data", http.StatusBadRequest)
				return
			}

			tmpl.ExecuteTemplate(w, "invoice.html", data)

			return
		}

		tmpl.ExecuteTemplate(w, "generate.html", nil)
	}
}

//invoice-form
func HandleIndex(tmpl *template.Template, db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := auth.GetSession(r)
		email := session.Values["email"].(string)

		_, err := repository.GetUserEmail(db, email)
		if err != nil {
			http.Redirect(w, r, "/setup", http.StatusSeeOther)
			return
		}
		// Ambil flash message dari session store
		success, _ := session.Values["success"].(bool)
		// reset agar muncul sekali saja
		session.Values["success"] = false
		session.Save(r, w)
		tmpl.ExecuteTemplate(w, "invoice-form.html", map[string]interface{}{
			"success": success,
		})
	}
}

func InvoicePDFHandler(w http.ResponseWriter, r *http.Request) {
	download := r.URL.Query().Get("download") == "true"
	err := service.GenerateInvoicePDF(w, r, download)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

