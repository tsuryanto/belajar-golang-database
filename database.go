package golang_database

import (
	"database/sql"
	"time"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database")
	if err != nil {
		panic(err)
	}

	// Koneksi Minimal
	db.SetMaxIdleConns(10)

	// Koneksi Maksimal
	db.SetMaxOpenConns(100)

	// Close koneksi jika 5 menit tidak digunakan
	db.SetConnMaxIdleTime(5 * time.Minute)

	// Jika sudah 60 menit koneksi manapun akan ditutup menjadi sejumlah koneksi minimal,
	// lalu dibuatkan koneksi baru
	db.SetConnMaxIdleTime(60 * time.Minute)

	return db
}
