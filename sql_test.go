package golang_database

import (
	"context"
	"fmt"
	"strconv"
	"testing"
)

func TestExecSqlParameter(t *testing.T) {
	// Dapatkan koneksi dari database
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "user"
	password := "user"

	// tanda tanya akan diisi dengan wsername dan password
	// WAJIB agar terhindar dari SQL Injection
	query := "INSERT INTO user(username, password) VALUES(?,?)"
	result, err := db.ExecContext(ctx, query, username, password)
	if err != nil {
		panic(err)
	}

	// Mendapatkan nilai balik berupa id yg diinsert
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println("Berhasil menginput data dengan ID", id)
}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin"
	password := "admin"

	query := "SELECT username from user WHERE username = ? AND password = ?"

	rows, err := db.QueryContext(ctx, query, username, password)
	defer rows.Close()

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Sukses Login", username)
	} else {
		fmt.Println("Gagal Login")
	}
}

//
func TestSqlPrepareStatement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO comments(email, comment) VALUES (?,?)"
	prepare, err := db.PrepareContext(ctx, query)
	if err != nil {
		panic(err)
	}

	var berhasil int
	var banyakData = 5
	// simulasi input 3 data sekaligus
	for i := 1; i <= banyakData; i++ {
		email := "admin" + strconv.Itoa(i) + "@gmail.com"
		comment := "Alhamdulillah"

		result, err := prepare.ExecContext(ctx, email, comment)
		if err != nil {
			panic(err)
		}

		// Mengambil jumlah row yang berhasil diinput
		affected, err := result.RowsAffected()
		if err != nil {
			panic(err)
		} else {
			if affected != 0 {
				fmt.Println("Input Berhasil")
				berhasil++
			} else {
				fmt.Println("Input Gagal")
			}
		}
	}

	fmt.Println("====================")
	if berhasil == banyakData {
		fmt.Println("Berhasil input semua data")
	} else {
		fmt.Println("Input Berhasil", berhasil)
		fmt.Println("Input Gagal", banyakData-berhasil)
	}
}

// Input data dengan memungkinkan pembatalan data yang telah diinput
func TestSqlTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	query := "INSERT INTO comments(email, comment) VALUES (?,?)"
	// simulasi input 3 data sekaligus
	for i := 1; i <= 3; i++ {
		email := "eko" + strconv.Itoa(i) + "@gmail.com"
		comment := "Hallo"

		result, err := tx.ExecContext(ctx, query, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("Comment Id", id)
	}

	err = tx.Commit() // Untuk mengukuhkan input

	// err = tx.Rollback() // Untuk membatalkan input

	if err != nil {
		panic(err)
	}
}
