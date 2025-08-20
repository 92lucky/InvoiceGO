package service

import (
	"invoice-go/model"
	"io"
	"mime/multipart"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ParseExcelToDataRow(file multipart.File) ([]model.DataRow, error) {
	tmp, err := os.CreateTemp("", "upload-*.xlsx")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmp.Name())

	_, err = io.Copy(tmp, file)
	if err != nil {
		return nil, err
	}
	tmp.Close()

	f, err := excelize.OpenFile(tmp.Name())
	if err != nil {
		return nil, err
	}

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	var data []model.DataRow
	for i, row := range rows {
		if i == 0 || len(row) < 4 {
			continue
		}
		if row[2] == "" || row[3] == "" {
			continue
		}

		no, _ := strconv.Atoi(row[0])
		tanggal := row[1]
		noSO := row[2]
		noLO := row[3]

		data = append(data, model.DataRow{
			No:          no,
			Date:        tanggal,
			NoSO:        noSO,
			NoLO:        noLO,
			JumlahTbg:   560,
			JumlahKg:    1680,
			Tarif:       354.64,
			BiayaAngkut: 595.795,
		})
	}
	return data, nil
}

func ParseExcelToDataRows(file multipart.File) ([]model.DataRow, error) {
	tmp, err := os.CreateTemp("", "upload-*.xlsx")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmp.Name())

	_, err = io.Copy(tmp, file)
	if err != nil {
		return nil, err
	}
	tmp.Close()

	f, err := excelize.OpenFile(tmp.Name())
	if err != nil {
		return nil, err
	}

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		return nil, err
	}

	var data []model.DataRow
	for i, row := range rows {
		if i == 0 || len(row) < 4 {
			continue
		}
		if row[2] == "" || row[3] == "" {
			continue
		}

		no, _ := strconv.Atoi(row[0])
		tanggal := row[1]
		noSO := row[2]
		noLO := row[3]

		data = append(data, model.DataRow{
			No:          no,
			Date:        tanggal,
			NoSO:        noSO,
			NoLO:        noLO,
			JumlahTbg:   560,
			JumlahKg:    1680,
			Tarif:       354.64,
			BiayaAngkut: 595.795,
		})
	}
	return data, nil
}
