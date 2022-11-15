package test

import (
	"database/sql"
	"fmt"
	"reza_50420900_pert3/model"
	"testing"
)

var username, password, host, namaDB, defaultDB string

func init() {
	username = "root"
	password = ""
	host = "localhost"
	namaDB = "reza_50420900_pert3"
	defaultDB = "mysql"
}
func TestDatabase(t *testing.T) {

	t.Run("Testing untuk membuat database", func(t *testing.T) {
		db, err := model.Connect(username, password, host, defaultDB)

		defer db.Close()

		if err != nil {
			t.Fatal(err)
		}

		err = model.CreateDB(db, namaDB)

		if err != nil {
			t.Fatal(err)
		}

	})
	t.Run("Testing untuk memeriksa koneksi database", func(t *testing.T) {

		db, err := model.Connect(username, password, host, defaultDB)

		defer db.Close()

		if err != nil {
			t.Fatal(err)
		}
	})
}
func initDatabase() (*sql.DB, error) {

	dbInit, err := model.Connect(username, password, host, defaultDB)

	if err != nil {
		fmt.Println("Gagal melakukan koneksi")
	}

	if err = model.DropDB(dbInit, namaDB); err != nil {
		fmt.Println("Gagal menghapus database")
		return nil, err
	}

	if err = model.CreateDB(dbInit, namaDB); err != nil {
		fmt.Println("Gagal membuat database")
		return nil, err
	}

	dbInit.Close()

	db, err := model.Connect(username, password, password, namaDB)

	if err != nil {
		fmt.Println("Gagal melakukan koneksi")
		return nil, err
	}

	if err = model.CreateTable(db, model.TabelMahasiswa); err != nil {
		fmt.Println("Gagal membuat table mahasiswa")
		return nil, err
	}
	if err = model.CreateTable(db, model.TabelMatkul); err != nil {
		fmt.Println("Gagal membuat table matkul")
		return nil, err
	}

	if err = model.CreateTable(db, model.TabelNilai); err != nil {
		fmt.Println("Gagal membuat table nilai")
		return nil, err
	}

	return db, nil
}

// > Test run finished at 11/2/2022, 10:57:59 PM <

// Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestDatabase$ reza_50420900_pert3/test

// === RUN   TestDatabase
// === RUN   TestDatabase/Testing_untuk_membuat_database
// === RUN   TestDatabase/Testing_untuk_memeriksa_koneksi_database
// --- PASS: TestDatabase (0.01s)
//     --- PASS: TestDatabase/Testing_untuk_membuat_database (0.01s)
//     --- PASS: TestDatabase/Testing_untuk_memeriksa_koneksi_database (0.00s)
// PASS
// ok      reza_50420900_pert3/test        0.243s

// > Test run finished at 11/2/2022, 10:58:23 PM <
