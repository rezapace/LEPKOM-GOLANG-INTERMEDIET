package test

import (
	"fmt"
	"reza_50420900_pert3/model"
	"testing"
)

func TestMatkul(t *testing.T) {

	var dataInsertMatkul = []model.Matkul{
		model.Matkul{
			Kd_mk:      "KA01",
			Matakuliah: "MTK Dasar",
		},
		model.Matkul{
			Kd_mk:      "KA02",
			Matakuliah: "B.Indo Dasar",
		},
		model.Matkul{
			Kd_mk:      "KA3",
			Matakuliah: "IPA Dasar",
		},
	}

	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("Testing insert  mahasiswwa", func(t *testing.T) {
		for _, dataInsert := range dataInsertMatkul {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Testing update  mahasiswwa", func(t *testing.T) {
		var updateData = map[string]interface{}{
			"matakuliah": "Algoritma 1",
		}

		data := dataInsertMatkul[0]
		if err := data.Update(db, updateData); err != nil {
			t.Fatal(err)
		}

	})

	t.Run("Testing Get matkul", func(t *testing.T) {
		// data := dataInsertMatkul[0]
		_, err := model.GetMatkul(db, "KA3")
		if err != nil {
			t.Fatal(err)
		}
		// fmt.Println(result)
	})

	t.Run("Testing GetAll matkul", func(t *testing.T) {
		result, err := model.GetAllMatkul(db)
		if err != nil {
			t.Fatal(err)
		}

		fmt.Println(result[0])
	})

	t.Run("Testing delete matkul", func(t *testing.T) {
		data := dataInsertMatkul[0]
		if err := data.Delete(db); err != nil {
			t.Fatal(err)
		}
	})

	defer db.Close()
}

// Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestMatkul$ reza_50420900_pert3/test

// === RUN   TestMatkul
// === RUN   TestMatkul/Testing_insert__mahasiswwa
// === RUN   TestMatkul/Testing_update__mahasiswwa
// === RUN   TestMatkul/Testing_Get_matkul
// === RUN   TestMatkul/Testing_GetAll_matkul
// &{KA01 Algoritma 1}
// === RUN   TestMatkul/Testing_delete_matkul
// --- PASS: TestMatkul (0.07s)
//     --- PASS: TestMatkul/Testing_insert__mahasiswwa (0.00s)
//     --- PASS: TestMatkul/Testing_update__mahasiswwa (0.00s)
//     --- PASS: TestMatkul/Testing_Get_matkul (0.00s)
//     --- PASS: TestMatkul/Testing_GetAll_matkul (0.00s)
//     --- PASS: TestMatkul/Testing_delete_matkul (0.00s)
// PASS
// ok      reza_50420900_pert3/test        0.313s

// > Test run finished at 11/2/2022, 11:02:57 PM <
