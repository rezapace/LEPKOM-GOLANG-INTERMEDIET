package test

import (
	"database/sql"
	"fmt"
	"reza_50420900_pert4/model" // Ganti dengan nama folder kalian masing￾masing
	"testing"
)

var username, password, host, namaDB, defaultDB string

func init() {
	username = "root"
	password = ""
	host = "localhost"
	namaDB = "db_50420900_pert4" // Ganti nama db sesuai dengan NPM masing￾masing
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
	/* Syntax dibawah ini berfungsi untuk menghapus DB (Apabila ingin
	   melakukan testing,
	   silahkan di comment terlebih dahulu)
	*/
	// t.Run("Testing untuk menghapus database", func(t *testing.T) {
	// db, err := model.Connect(username, password, host, defaultDB)
	// defer db.Close()
	// if err != nil {
	// t.Fatal(err)
	// }
	// err = model.DropDB(db, namaDB)
	// if err != nil {
	// t.Fatal(err)
	// }
	// })
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

// ? output Database_test.go

// todo || go test -run TestDatabase ||

// Windows PowerShell
// Copyright (C) Microsoft Corporation. All rights reserved.

// Try the new cross-platform PowerShell https://aka.ms/pscore6

// PS C:\Program Files\Go\src\reza_50420900_pert4\test> go test -run TestDatabase
// PASS
// ok      reza_50420900_pert4/test        0.221s
// PS C:\Program Files\Go\src\reza_50420900_pert4\test>
