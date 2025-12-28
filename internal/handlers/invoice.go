package handlers

import (
	"html/template"
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

func InvoicePDFHandler(w http.ResponseWriter, r *http.Request) {
	download := r.URL.Query().Get("download") == "true"
	err := service.GenerateInvoicePDF(w, r, download)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
