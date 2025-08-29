package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB // Global DB instance

func Init() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("❌ DATABASE_URL belum di-set di environment variables")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("❌ Gagal buka koneksi DB: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Gagal ping DB: %v", err)
	}

	log.Println("✅ Koneksi DB berhasil")
	DB = db
}
