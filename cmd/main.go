package main

import (
	"html/template"
	"invoice-go/auth"
	"invoice-go/config"
	"invoice-go/routes"
	"log"
	"net/http"
	"os"

	"github.com/dustin/go-humanize"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(".env"); err != nil {
		log.Println("Could not load .env")
	}

	// Check SESSION_KEY
	if os.Getenv("SESSION_KEY") == "" {
		log.Fatal("SESSION_KEY missing")
	}

	// Init session
	auth.InitSession()

	// Templates
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"formatRupiah": func(n float64) string {
			return humanize.Comma(int64(n))
		},
	}).ParseGlob("templates/*.html"))

	// DB
	config.Init()
	defer config.DB.Close()

	// OAuth
	auth.InitOAuthConfig()

	// Routes
	mux := http.NewServeMux()
	auth.RegisterAuthRoutes(mux)
	routes.RegisterAppRoutes(mux, tmpl, config.DB)

	// Server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running at http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("Server failed:", err)
	}
}
