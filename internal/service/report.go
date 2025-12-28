// service/pdf_service.go
package service

import (
	"invoice-go/internal/model"
	"invoice-go/utils"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
)

func GeneratePDFLo(data []model.DataRow, namaPT, bulan string, w http.ResponseWriter) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "B", 14)
	pdf.Cell(0, 10, namaPT)
	pdf.Ln(6)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 9, "Agen LPG 3 Kg PT Pertamina (Persero)")
	pdf.Ln(6)
	pdf.SetFont("Arial", "B", 11)
	pdf.Cell(0, 9, "Bulan : "+bulan)
	pdf.Ln(12)

	headers := []string{"NO", "Date", "No.SO", "No.LO", "Jum (tbg)", "Jum (Kg)", "TARIF", "BIAYA ANGKUT"}
	widths := []float64{8, 22, 26, 24, 20, 20, 30, 35} // ← Lebarkan tarif & biaya

	pdf.SetFont("Arial", "B", 10)
	for i, h := range headers {
		pdf.CellFormat(widths[i], 6, h, "1", 0, "C", false, 0, "")
	}
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10)
	var totalTbg, totalKg int
	for _, row := range data {
		vals := []string{
			strconv.Itoa(row.No),
			row.Date,
			row.NoSO,
			row.NoLO,
			strconv.Itoa(row.JumlahTbg),
			strconv.Itoa(row.JumlahKg),
			"Rp. " + utils.FormatRupiah(row.Tarif),
			"Rp. " + utils.FormatRupiah(row.BiayaAngkut),
		}

		for i, val := range vals {
			pdf.CellFormat(widths[i], 5, val, "1", 0, "C", false, 0, "")
		}
		pdf.Ln(-1)

		totalTbg += row.JumlahTbg
		totalKg += row.JumlahKg
	}

	// Hitung total biaya langsung dari total KG
	totalBiaya := float64(totalKg) * 354.64

	// Baris TOTAL
	pdf.SetFont("Arial", "B", 10)
	for i := 0; i < len(widths); i++ {
		switch i {
		case 3:
			pdf.CellFormat(widths[i], 5, "TOTAL", "1", 0, "R", false, 0, "")
		case 4:
			pdf.CellFormat(widths[i], 5, strconv.Itoa(totalTbg), "1", 0, "C", false, 0, "")
		case 5:
			pdf.CellFormat(widths[i], 5, strconv.Itoa(totalKg), "1", 0, "C", false, 0, "")
		case 6:
			pdf.CellFormat(widths[i], 5, "-", "1", 0, "C", false, 0, "")
		case 7:
			pdf.CellFormat(widths[i], 5, "Rp. "+utils.Formatt(totalBiaya), "1", 0, "C", false, 0, "")
		default:
			pdf.CellFormat(widths[i], 5, "", "1", 0, "C", false, 0, "")
		}
	}
	pdf.Ln(5)

	// PPN dan Grand Total
	ppn := totalBiaya * 0.12
	grandTotal := totalBiaya + ppn

	pdf.SetFont("Arial", "", 10)
	for i := 0; i < len(widths); i++ {
		switch i {
		case 6:
			pdf.CellFormat(widths[i], 5, "PPN 12%", "1", 0, "R", false, 0, "")
		case 7:
			pdf.CellFormat(widths[i], 5, "Rp. "+utils.Formatt(ppn), "1", 0, "R", false, 0, "")
		default:
			pdf.CellFormat(widths[i], 5, " ", "1", 0, "", false, 0, "")
		}
	}
	pdf.Ln(5)

	pdf.SetFont("Arial", "B", 10)
	for i := 0; i < len(widths); i++ {
		switch i {
		case 6:
			pdf.CellFormat(widths[i], 5, "Grand Total", "1", 0, "R", false, 0, "")
		case 7:
			pdf.CellFormat(widths[i], 5, "Rp. "+utils.Formatt(grandTotal), "1", 0, "R", false, 0, "")
		default:
			pdf.CellFormat(widths[i], 5, " ", "1", 0, "", false, 0, "")
		}
	}
	pdf.Ln(30)
	pdf.SetFont("Arial", "", 8)
	pdf.Cell(0, 10, namaPT)

	return pdf.Output(w)
}


//excel export
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

