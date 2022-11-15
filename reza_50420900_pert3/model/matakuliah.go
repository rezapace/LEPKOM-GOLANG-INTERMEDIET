package model

import (
	"database/sql"
	"fmt"
	"strings"
)

var TabelMatkul string = `
	CREATE TABLE matkul(
		kd_mk VARCHAR(10) PRIMARY KEY,
		matakuliah VARCHAR(20)
	);
`

func (m *Matkul) Structur() *Matkul {
	return &Matkul{}
}

type Matkul struct {
	Kd_mk      string `json:"Kd_mk"`
	Matakuliah string `json:"Matakuliah"`
}

func (m *Matkul) Fields() ([]string, []interface{}) {
	fields := []string{"kd_mk", "matakuliah"}
	temp := []interface{}{&m.Kd_mk, &m.Matakuliah}
	return fields, temp
}

func (m *Matkul) Insert(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO %v values(?,?)", "matkul")
	_, err := db.Exec(query, &m.Kd_mk, &m.Matakuliah)
	return err
}

func (m *Matkul) Update(db *sql.DB, data map[string]interface{}) error {

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

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", "matkul", dataUpdate, "kd_mk")
	args = append(args, m.Kd_mk)
	// Exec dengan query yang ada
	_, err := db.Exec(query, args...)
	return err
}

func (m *Matkul) Delete(db *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", "matkul", "kd_mk")
	// Exec dengan query yang ada
	_, err := db.Exec(query, m.Kd_mk)
	return err
}

func GetMatkul(db *sql.DB, id string) (*Matkul, error) {

	m := &Matkul{}
	each := m.Structur()
	_, dst := each.Fields()
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", "matkul", "kd_mk")

	// isinya akan dimasukan kedalam var dst yang dideklarasikan diatas
	err := db.QueryRow(query, id).Scan(dst...)
	if err != nil {
		return nil, err
	}
	return each, nil
}

func GetAllMatkul(db *sql.DB) ([]*Matkul, error) {
	m := &Matkul{}
	query := fmt.Sprintf("SELECT * FROM %s", "matkul")
	data, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer data.Close()

	var result []*Matkul
	for data.Next() {
		each := m.Structur()
		_, dst := each.Fields()
		err := data.Scan(dst...)

		if err != nil {
			return nil, err
		}
		result = append(result, each)
	}

	return result, nil

}
