package model

import (
	"database/sql"
	"fmt"
	"strings"
)

var TabelNilai string = `
	CREATE TABLE nilai(
		id_nilai INT PRIMARY KEY AUTO_INCREMENT,
		kd_mk VARCHAR(10) ,
		npm VARCHAR(10) ,
		uts REAL NOT NULL,
		uas REAL NOT NULL,
		total REAL NOT NULL,
		grade VARCHAR(5) NOT NULL
	)
`

type Nilai struct {
	Id_nilai int     `json:"id_nilai"`
	Kd_mk    string  `json:"kd_mk"`
	NPM      string  `json:"npm"`
	UAS      float64 `json:"uas"`
	UTS      float64 `json:"uts"`
	Total    float64 `json:"total"`
	Grade    string  `json:"grade"`
}

func (m *Nilai) Structur() *Nilai {
	return &Nilai{}
}

func (m *Nilai) Fields() ([]string, []interface{}) {
	fields := []string{"id_nilai", "kd_mk", "npm", "uts", "uas", "total", "grade"}
	temp := []interface{}{&m.Id_nilai, &m.Kd_mk, &m.NPM, &m.UAS, &m.UTS, &m.Total, &m.Grade}
	return fields, temp
}

func (m *Nilai) Insert(db *sql.DB) error {
	query := fmt.Sprintf("INSERT INTO %v values(?,?,?,?,?,?,?)", "nilai")
	_, err := db.Exec(query, &m.Id_nilai, &m.Kd_mk, &m.NPM, &m.UAS, &m.UTS, &m.Total, &m.Grade)
	return err
}

func (m *Nilai) Update(db *sql.DB, data map[string]interface{}) error {

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

	query := fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", "nilai", dataUpdate, "npm")
	args = append(args, m.NPM)
	// Exec dengan query yang ada
	_, err := db.Exec(query, args...)
	return err
}

func (m *Nilai) Delete(db *sql.DB) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s = ?", "nilai", "npm")
	// Exec dengan query yang ada
	_, err := db.Exec(query, m.NPM)
	return err
}

func GetNilai(db *sql.DB, id string) (*Nilai, error) {

	m := &Nilai{}
	each := m.Structur()
	_, dst := each.Fields()
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s = ?", "nilai", "npm")

	// isinya akan dimasukan kedalam var dst yang dideklarasikan diatas
	err := db.QueryRow(query, id).Scan(dst...)
	if err != nil {
		return nil, err
	}
	return each, nil
}

func GetAllNilai(db *sql.DB) ([]*Nilai, error) {
	m := &Nilai{}
	query := fmt.Sprintf("SELECT * FROM %s", "nilai")
	data, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer data.Close()

	var result []*Nilai
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
