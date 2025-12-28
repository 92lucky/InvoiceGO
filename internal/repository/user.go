package repository

import (
	"database/sql"
	"invoice-go/internal/model"
)

func SaveUserProfile(db *sql.DB, p model.AppProfile) error {
	_, err := db.Exec(`
		INSERT INTO user_profile (email, nama_pt, nama_bank, no_rekening, penanggung_jawab, alamat, kabupaten)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (email) DO UPDATE SET
			nama_pt = EXCLUDED.nama_pt,
			nama_bank = EXCLUDED.nama_bank,
			no_rekening = EXCLUDED.no_rekening,
			penanggung_jawab = EXCLUDED.penanggung_jawab,
			alamat = EXCLUDED.alamat,
			kabupaten = EXCLUDED.kabupaten
	`, p.Email, p.NamaPT, p.NamaBank, p.NoRekening, p.PenanggungJawab, p.Alamat, p.Kabupaten)
	return err
}

func GetUserEmail(db *sql.DB, email string) (*model.AppProfile, error) {
	row := db.QueryRow(`
		SELECT email, nama_pt, nama_bank, no_rekening, penanggung_jawab, alamat, kabupaten
		FROM user_profile
		WHERE email = $1
	`, email)

	var profile model.AppProfile
	err := row.Scan(
		&profile.Email,
		&profile.NamaPT,
		&profile.NamaBank,
		&profile.NoRekening,
		&profile.PenanggungJawab,
		&profile.Alamat,
		&profile.Kabupaten,
	)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

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

