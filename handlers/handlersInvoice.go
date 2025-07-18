package handlers

import (
	"html/template"
	"invoice-go/service"
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

			err = tmpl.ExecuteTemplate(w, "invoice.html", data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}

		tmpl.ExecuteTemplate(w, "generate.html", nil)
	}
}
