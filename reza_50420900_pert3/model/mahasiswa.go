package model

import (
	"database/sql"
	"fmt"
	"strings"
)

var TabelMahasiswa string = `
	CREATE TABLE mahasiswa(
		npm VARCHAR(10) PRIMARY KEY,
		nama VARCHAR(30),
		kelas VARCHAR(10)
	);
`

type Mahasiswa struct {
	NPM   string `json:"NPM"`
	Nama  string `json:"Nama"`
	Kelas string `json:"Kelas"`
}

func (m *Mahasiswa) Fields() ([]string, []interface{}) {
	fields := []string{"npm", "nama", "kelas"}
	temp := []interface{}{&m.NPM, &m.Nama, &m.Kelas}
	return fields, temp
}

func (m *Mahasiswa) Structur() *Mahasiswa {
	return &Mahasiswa{}
}

func (m *Mahasiswa) Insert(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO %v values(?,?,?)", "mahasiswa")
	_, err := db.Exec(query, &m.NPM, &m.Nama, &m.Kelas)
	return err
}

func (m *Mahasiswa) Update(db *sql.DB, data map[string]interface{}) error {

	var kolom = []string{}
	var args []interface{}

	i := 1

	// Ini loop data untuk dimasukan kedalam set,
	for key, value := range data {
		updateData := fmt.Sprintf("%v = ?", strings.ToLower(key))
		kolom = append(kolom, updateData)
		args = append(args, value)
		i++
	}

	// Ubah array menjadi string dengan pemisah koma
	dataUpdate := strings.Join(kolom, ",")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", "mahasiswa", dataUpdate, "NPM")
	args = append(args, m.NPM)
	// Exec dengan query yang ada
	_, err := db.Exec(query, args...)
	return err
}

func (m *Mahasiswa) Delete(db *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", "mahasiswa", "NPM")
	// Exec dengan query yang ada
	_, err := db.Exec(query, m.NPM)
	return err
}

func GetMahasiswa(db *sql.DB, id string) (*Mahasiswa, error) {

	m := &Mahasiswa{}
	each := m.Structur()
	_, dst := each.Fields()
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", "mahasiswa", "NPM")

	// isinya akan dimasukan kedalam var dst yang dideklarasikan diatas
	err := db.QueryRow(query, id).Scan(dst...)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return each, nil
}

func GetAllMahasiswa(db *sql.DB) ([]*Mahasiswa, error) {
	m := &Mahasiswa{}
	query := fmt.Sprintf("SELECT * FROM %s", "mahasiswa")
	data, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer data.Close()

	var result []*Mahasiswa

	for data.Next() {
		each := m.Structur()
		_, dst := each.Fields()

		err := data.Scan(dst...)
		if err != nil {
			return nil, err
		}
		fmt.Println(each)
		result = append(result, each)
	}

	return result, nil

}
