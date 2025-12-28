package routes

import (
	"database/sql"
	"html/template"
	"invoice-go/internal/auth"
	"invoice-go/internal/handlers"
	"net/http"
)

// daftar semua route
func RegisterAppRoutes(mux *http.ServeMux, tmpl *template.Template, db *sql.DB) {
	mux.HandleFunc("/", handlers.HandleHome(tmpl))
	mux.HandleFunc("/index", auth.RequireAuth(handlers.HandleIndex(tmpl, db)))
	mux.HandleFunc("/setup", auth.RequireAuth(handlers.HandleSetup(tmpl)))

	//Invoice section
	mux.HandleFunc("/generate", auth.RequireAuth(handlers.HandlersInvoice(tmpl)))
	mux.HandleFunc("/generate-pdf", auth.RequireAuth(handlers.InvoicePDFHandler))
	//style
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// LO section
	mux.HandleFunc("/lo", auth.RequireAuth(handlers.ShowLoPage(tmpl)))
	mux.HandleFunc("/previewLo", auth.RequireAuth(handlers.PreviewLoHandler()))
	mux.HandleFunc("/downloadLo", auth.RequireAuth(handlers.DownloadLoHandler()))

}
