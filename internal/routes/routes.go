package routes

import (
	"database/sql"
	"html/template"
	"invoice-go/internal/auth"
	"invoice-go/internal/handlers"
	"net/http"
)

// regist
func RegisterAppRoutes(mux *http.ServeMux, tmpl *template.Template, db *sql.DB) {

	//layer
	mux.HandleFunc("/", handlers.HandleHome(tmpl))
	mux.HandleFunc("/dashboard", handlers.HandleDashboard(tmpl))
	


	//user
	mux.HandleFunc("/setup", auth.RequireAuth(handlers.HandleSetup(tmpl)))
	mux.HandleFunc("/dev/reset-setup", handlers.HandleResetSetup())


	//Invoice section
	mux.HandleFunc("/invoice-form", auth.RequireAuth(handlers.HandleIndex(tmpl, db)))
	mux.HandleFunc("/generate", auth.RequireAuth(handlers.HandlersInvoice(tmpl)))
	mux.HandleFunc("/generate-pdf", auth.RequireAuth(handlers.InvoicePDFHandler))
	
	//template handle
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// report section
	mux.HandleFunc("/report", auth.RequireAuth(handlers.ShowLoPage(tmpl)))
	mux.HandleFunc("/previewLo", auth.RequireAuth(handlers.PreviewLoHandler()))
	mux.HandleFunc("/downloadLo", auth.RequireAuth(handlers.DownloadLoHandler()))

}


