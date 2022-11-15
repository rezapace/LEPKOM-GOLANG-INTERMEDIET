package test

import (
	"reza_50420900_pert4/model" // Ganti dengan nama folder kalian masing￾masing
	"testing"
)

func TestMahasiswa(t *testing.T) {
	var dataInsertMhs = []model.Mahasiswa{
		model.Mahasiswa{
			NPM:   "12345678",
			Nama:  "Budi Doremi",
			Kelas: "3IA20",
		},
		model.Mahasiswa{
			NPM:   "50420900",
			Nama:  "REZA",
			Kelas: "4IA20",
		},
		model.Mahasiswa{
			NPM:   "44444444",
			Nama:  "DoBud",
			Kelas: "4IA21",
		},
	}
	db, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Testing insert mahasiswa", func(t *testing.T) {
		for _, dataInsert := range dataInsertMhs {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})
	t.Run("Testing update mahasiswa", func(t *testing.T) {
		var updateData = map[string]interface{}{
			"nama": "Abdi Teh reza"}
		data := dataInsertMhs[0]
		if err := data.Update(db, updateData); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Testing Get mahasiswa", func(t *testing.T) {
		_, err := model.GetMahasiswa(db, "44444444")
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Testing Get mahasiswa", func(t *testing.T) {
		_, err := model.GetAllMahasiswa(db)
		if err != nil {
			t.Fatal(err)
		}
	})
	// t.Run("Testing delete mahasiswa", func(t *testing.T) {
	// data := dataInsertMhs[0]
	// if err := data.Delete(db); err != nil {
	// t.Fatal(err)
	// }
	// })
	defer db.Close()
}

// ? output Mahasiswa_test.go

// todo || go test -run TestMahasiswa ||
// todo || CD.. ||
// todo || go run main.go ||

// Windows PowerShell
// Copyright (C) Microsoft Corporation. All rights reserved.

// Try the new cross-platform PowerShell https://aka.ms/pscore6

// PS C:\Program Files\Go\src\reza_50420900_pert4\test> go test -run TestMahasiswa
// &{12345678 Abdi Teh reza 3IA20}
// &{44444444 DoBud 4IA21}
// &{50420900 REZA 4IA20}
// PASS
// ok      reza_50420900_pert4/test        0.296s
// PS C:\Program Files\Go\src\reza_50420900_pert4\test>
