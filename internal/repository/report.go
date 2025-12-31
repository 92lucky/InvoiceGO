package repository

import (
	"database/sql"
	"invoice-go/internal/model"
)


func GetDataLo(db *sql.DB) ([]model.DataRow, error) {
	rows, err := db.Query(`
		SELECT 
			no, tanggal,no_so, no_lo, jumlah_tabung, jumlah_kg, tarif, biaya_angkut
		FROM lo_bulanan
		ORDER BY tanggal
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hasil []model.DataRow
	for rows.Next() {
		var row model.DataRow
		err := rows.Scan(&row.No, &row.Date, &row.NoSO, &row.NoLO, &row.JumlahTbg, &row.JumlahKg, &row.Tarif, &row.BiayaAngkut)
		if err != nil {
			return nil, err
		}
		hasil = append(hasil, row)
	}
	return hasil, nil
}

