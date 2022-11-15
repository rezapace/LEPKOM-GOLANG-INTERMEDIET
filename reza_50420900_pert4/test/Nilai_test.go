package test

import (
	"fmt"
	"reza_50420900_pert4/model" // Ganti dengan nama folder kalian masing￾masing
	"testing"
)

func TestNilai(t *testing.T) {
	var dataInsertNilai = []model.Nilai{
		model.Nilai{
			Id_nilai: 1,
			Kd_mk:    "IA02",
			NPM:      "50420900",
			UTS:      70,
			UAS:      70,
			Total:    70,
			Grade:    "B",
		},
		model.Nilai{
			Id_nilai: 2,
			Kd_mk:    "IA03",
			NPM:      "44444444",
			UTS:      50,
			UAS:      50,
			Total:    50,
			Grade:    "D",
		},
	}
	db, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Testing insert nilai", func(t *testing.T) {
		for _, dataInsert := range dataInsertNilai {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})
	t.Run("Testing update nilai", func(t *testing.T) {
		var updateData = map[string]interface{}{
			"uts":   30,
			"uas":   30,
			"total": 30,
			"grade": "E"}
		data := dataInsertNilai[0]
		if err := data.Update(db, updateData); err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Testing GetAll nilai", func(t *testing.T) {
		// data := dataInsertNilai[0]
		_, err := model.GetNilai(db, "44444444")
		if err != nil {
			t.Fatal(err)
		}
		// fmt.Println(result)
	})
	t.Run("Testing Get nilai", func(t *testing.T) {
		result, err := model.GetAllNilai(db)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(result[0])
	})
	// t.Run("Testing delete nilai", func(t *testing.T) {
	// data := dataInsertNilai[0]
	// if err := data.Delete(db); err != nil {
	// t.Fatal(err)
	// }
	// })
	defer db.Close()
}

// ? output nilai_test.go

// todo || go test -run TestNilai ||
// todo || CD.. ||
// todo || go run main.go ||

// Windows PowerShell
// Copyright (C) Microsoft Corporation. All rights reserved.

// Try the new cross-platform PowerShell https://aka.ms/pscore6

// PS C:\Program Files\Go\src\reza_50420900_pert4\test> go test -run TestNilai
// &{1 IA02 50420900 30 30 30 E}
// PASS
// ok      reza_50420900_pert4/test        0.278s
// PS C:\Program Files\Go\src\reza_50420900_pert4\test>
