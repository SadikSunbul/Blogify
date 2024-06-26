package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB // database connection

// InitDB veritabanını başlatır
func InitDB(dataSourceName string) {
	var err error                                  // veritabanı baglantı hatası
	DB, err = sql.Open("postgres", dataSourceName) // veritabanı baglantısı yapılır ve hata kontrol edilir
	if err != nil {
		log.Panic(err) // hata varsa loga yazılır
	}

	if err = DB.Ping(); err != nil { // baglantıyı kontrol eder ve hata varsa loga yazılır
		log.Panic(err)
	}
}
