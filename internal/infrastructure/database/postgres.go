package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB // database connection

// InitDB veritabanını başlatır
// InitDB initializes the database connection
func InitDB(dataSourceName string) *sql.DB {
	var err error
	db, err := sql.Open("postgres", dataSourceName) // veritabanı bağlantısı oluşturuldu
	if err != nil {
		log.Fatalf("Veritabanı bağlantısı başlatılamadı: %v", err)
	}

	// Ping the database to verify the connection
	err = db.Ping() // veritabanına ping atılır
	if err != nil {
		log.Fatalf("Veritabanına ping atılamadı: %v", err)
	}

	log.Println("Veritabanı bağlantısı başarıyla başlatıldı")
	return db
}
