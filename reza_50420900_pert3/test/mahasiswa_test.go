package test

import (
	"reza_50420900_pert3/model" //sesuaikan dengan nama folder (case sensitive)
	"testing"
)

func TestMahasiswa(t *testing.T) {

	var dataInsertMhs = []model.Mahasiswa{
		model.Mahasiswa{
			NPM:   "50420900",
			Nama:  "REZA",
			Kelas: "3IA20",
		},
		model.Mahasiswa{
			NPM:   "19283746",
			Nama:  "Doremi Budi",
			Kelas: "4KA20",
		},
		model.Mahasiswa{
			NPM:   "44444444",
			Nama:  "DoBud",
			Kelas: "4KA21",
		},
	}

	db, err := initDatabase()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("Testing insert  mahasiswa", func(t *testing.T) {
		for _, dataInsert := range dataInsertMhs {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
		}
	})

	t.Run("Testing update  mahasiswa", func(t *testing.T) {
		var updateData = map[string]interface{}{
			"nama": "Abdi Teh Ayeuna",
		}

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

	t.Run("Testing delete mahasiswa", func(t *testing.T) {
		data := dataInsertMhs[0]
		if err := data.Delete(db); err != nil {
			t.Fatal(err)
		}
	})
	defer db.Close()
}

// Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestMahasiswa$ reza_50420900_pert3/test

// === RUN   TestMahasiswa
// === RUN   TestMahasiswa/Testing_insert__mahasiswa
// === RUN   TestMahasiswa/Testing_update__mahasiswa
// === RUN   TestMahasiswa/Testing_Get_mahasiswa
// === RUN   TestMahasiswa/Testing_Get_mahasiswa#01
// &{19283746 Doremi Budi 4KA20}
// &{44444444 DoBud 4KA21}
// &{50420900 Abdi Teh Ayeuna 3IA20}
// === RUN   TestMahasiswa/Testing_delete_mahasiswa
// --- PASS: TestMahasiswa (0.08s)
//     --- PASS: TestMahasiswa/Testing_insert__mahasiswa (0.01s)
//     --- PASS: TestMahasiswa/Testing_update__mahasiswa (0.00s)
//     --- PASS: TestMahasiswa/Testing_Get_mahasiswa (0.00s)
//     --- PASS: TestMahasiswa/Testing_Get_mahasiswa#01 (0.00s)
//     --- PASS: TestMahasiswa/Testing_delete_mahasiswa (0.00s)
// PASS
// ok      reza_50420900_pert3/test        0.334s

// > Test run finished at 11/2/2022, 11:00:29 PM <
