package config

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB // Global DB instance

func Init() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("❌ DATABASE_URL belum di-set di environment variables")
	}

	var db *sql.DB
	var err error

	// Loop sampai DB siap
	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Println("❌ Gagal buka koneksi DB:", err)
		} else {
			err = db.Ping()
			if err == nil {
				break // DB siap
			}
			log.Println("❌ DB belum siap, retry...", err)
		}
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("❌ Tidak bisa konek ke DB setelah 10 percobaan:", err)
	}

	log.Println("✅ Koneksi DB berhasil")
	DB = db
}
