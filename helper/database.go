package helper

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection() *sql.DB {

	// parseTime=true ditulis agar tipedata DATE / TIME di mysql otomatis terconvert ke time golang
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/belajar_golang_database?parseTime=true")
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
