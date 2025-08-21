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
		log.Println("🔔 Tidak bisa load .env (mungkin karena running di Railway):", err)
	} else {
		log.Println("✅ .env berhasil dimuat")
	}

	// Pastikan SESSION_KEY ada
	if os.Getenv("SESSION_KEY") == "" {
		log.Fatal("❌ SESSION_KEY belum di-set! Set di .env atau di Railway environment variables.")
	}

	// Inisialisasi session
	auth.InitSession()

	// Fungsi template
	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"formatRupiah": func(n float64) string {
			return humanize.Comma(int64(n))
		},
	}).ParseGlob("templates/*.html"))

	// Koneksi DB
	config.Init()
	defer config.DB.Close()

	// Inisialisasi OAuth
	auth.InitOAuthConfig()

	// Routing
	mux := http.NewServeMux()
	auth.RegisterAuthRoutes(mux)
	routes.RegisterAppRoutes(mux, tmpl, config.DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("✅ Server berjalan di http://localhost:%s\n", port)

	err := http.ListenAndServe(":"+port, mux)
	if err != nil {
		log.Fatalf("❌ Gagal menjalankan server: %v", err)
	}
}
