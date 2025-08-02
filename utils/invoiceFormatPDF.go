package utils

import (
	"invoice-go/model"

	"github.com/dustin/go-humanize"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePDFInvoice(profile model.AppProfile, data model.InvoiceData) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(20, 15, 20)
	pdf.AddPage()

	// ===== HEADER =====
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 7, profile.NamaPT, "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 5, "AGEN LPG PSO", "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 5, profile.Alamat+" , "+profile.Kabupaten, "", 1, "C", false, 0, "")
	pdf.Ln(8)

	// ===== JUDUL INVOICE =====
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "INVOICE", "", 1, "C", false, 0, "")
	pdf.Ln(5)

	// ===== KEPADA =====
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 6, "Kepada : PT. Pertamina Patra Niaga")
	pdf.Ln(6)
	pdf.MultiCell(0, 5, "Alamat : Gedung Wisma Tugu II Lt.2\nJl. HR Rasuna Said KAV C7-9 Setiabudi, Jakarta 12920", "", "", false)
	pdf.Ln(8)

	// ===== TANGGAL & NOMOR =====
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(95, 6, "Tanggal : "+data.InvoiceDate, "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 6, "No. Invoice : "+data.InvoiceNumber, "", 1, "R", false, 0, "")
	pdf.Ln(8)

	// ===== TABEL =====
	tableWidth := 150.0
	pageWidth := 210.0
	margin := 20.0
	startX := margin + ((pageWidth-2*margin)-tableWidth)/2

	// Header tabel
	pdf.SetFont("Arial", "B", 11)
	pdf.SetFillColor(220, 220, 220)
	pdf.SetX(startX)
	pdf.CellFormat(90, 8, "Keterangan", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 8, "Kg", "1", 0, "C", true, 0, "")
	pdf.CellFormat(30, 8, "Nilai", "1", 1, "C", true, 0, "")

	// Data tabel
	pdf.SetFont("Arial", "", 11)
	fill := false
	rows := []struct {
		Label string
		Kg    string
		Nilai string
	}{
		{"Tagihan Transport Fee LPG 3 Kg Periode " + data.Periode, humanize.Comma(int64(data.DisplayQty)), ""},
		{"Pokok", "", humanize.Comma(int64(data.Pokok))},
		{"PPN 12%", "", humanize.Comma(int64(data.PPN))},
	}

	for _, row := range rows {
		if fill {
			pdf.SetFillColor(245, 245, 245)
		} else {
			pdf.SetFillColor(255, 255, 255)
		}
		pdf.SetX(startX)
		pdf.CellFormat(90, 8, row.Label, "1", 0, "L", true, 0, "")
		pdf.CellFormat(30, 8, row.Kg, "1", 0, "C", true, 0, "")
		pdf.CellFormat(30, 8, row.Nilai, "1", 1, "R", true, 0, "")
		fill = !fill
	}

	// ===== TOTAL =====
	pdf.SetFont("Arial", "B", 11)
	pdf.SetFillColor(230, 230, 250)
	pdf.SetX(startX)
	pdf.CellFormat(120, 8, "TOTAL", "1", 0, "R", true, 0, "")
	pdf.CellFormat(30, 8, "Rp. "+humanize.Comma(int64(data.Total)), "1", 1, "R", true, 0, "")

	// ===== TERBILANG =====
	pdf.Ln(8)
	pdf.SetFont("Arial", "I", 10)
	pdf.MultiCell(0, 6, "Terbilang : "+Terbilang(int64(data.Total)), "", "", false)

	// ===== BANK =====
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 6, "Bank : "+profile.NamaBank+" - "+profile.NoRekening)
	pdf.Ln(6)
	pdf.Cell(0, 6, "a.n "+profile.NamaPT)

	// ===== TANDA TANGAN =====
	pdf.Ln(20)
	pdf.CellFormat(0, 6, profile.Kabupaten+", "+data.InvoiceDate, "", 1, "R", false, 0, "")
	pdf.Ln(20)
	pdf.CellFormat(0, 6, profile.PenanggungJawab, "", 1, "R", false, 0, "")

	return pdf
}
