package utils

import (
	"invoice-go/model"

	"github.com/dustin/go-humanize"
	"github.com/jung-kurt/gofpdf"
)

func GeneratePDFInvoice(profile model.AppProfile, data model.InvoiceData) *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(15, 15, 15)
	pdf.AddPage()

	// ===== HEADER + GARIS =====
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(0, 7, profile.NamaPT, "", 1, "C", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(0, 5, "AGEN LPG PSO", "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 5, profile.Alamat+" , "+profile.Kabupaten, "", 1, "C", false, 0, "")
	pdf.Ln(2)
	pdf.SetDrawColor(0, 0, 0)
	pdf.Line(15, pdf.GetY(), 195, pdf.GetY()) // garis pembatas header
	pdf.Ln(6)

	// ===== JUDUL =====
	pdf.SetFont("Arial", "B", 16)
	pdf.CellFormat(0, 10, "INVOICE", "", 1, "C", false, 0, "")
	pdf.Ln(6)

	// ===== KEPADA =====
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 6, "Kepada : PT. Pertamina Patra Niaga")
	pdf.Ln(6)
	pdf.MultiCell(0, 6, "Alamat : Gedung Wisma Tugu II Lt.2,\nJl. HR Rasuna Said KAV C7-9 Setiabudi,\nJakarta 12920", "", "L", false)
	pdf.Ln(8)

	// ===== TANGGAL & NOMOR =====
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(95, 6, "Tanggal : "+data.InvoiceDate, "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 6, "No. Invoice : "+data.InvoiceNumber, "", 1, "R", false, 0, "")
	pdf.Ln(8)

	// ===== TABEL =====
	colWidths := []float64{8, 95, 20, 23, 37} // No | Keterangan | Kg | Detil | Nilai

	pdf.SetFont("Arial", "B", 11)
	pdf.SetFillColor(220, 220, 220)
	headers := []string{"No", "Keterangan", "Kg", "Detil", "Nilai"}
	for i, h := range headers {
		pdf.CellFormat(colWidths[i], 8, h, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// ===== ISI: 3 baris, Detil di kolom ke-4 =====
	pdf.SetFont("Arial", "", 10)

	// Baris 1: Keterangan + KG + Pokok
	pdf.CellFormat(colWidths[0], 8, "1", "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidths[1], 8, "Tagihan Transport Fee LPG 3 Kg Periode "+data.Periode, "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidths[2], 8, humanize.Comma(int64(data.DisplayQty)), "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidths[3], 8, "Pokok", "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidths[4], 8, "Rp. "+humanize.Comma(int64(data.Pokok)), "1", 1, "R", false, 0, "")

	// Baris 2: Detil = DPP
	pdf.CellFormat(colWidths[0], 8, "", "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidths[1], 8, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidths[2], 8, "", "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidths[3], 8, "DPP", "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidths[4], 8, "Rp. "+humanize.Comma(int64(data.DPP)), "1", 1, "R", false, 0, "")

	// Baris 3: Detil = PPN
	pdf.CellFormat(colWidths[0], 8, "", "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidths[1], 8, "", "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidths[2], 8, "", "1", 0, "C", false, 0, "")
	pdf.CellFormat(colWidths[3], 8, "PPN 12%", "1", 0, "L", false, 0, "")
	pdf.CellFormat(colWidths[4], 8, "Rp. "+humanize.Comma(int64(data.PPN)), "1", 1, "R", false, 0, "")

	// ===== TOTAL =====
	pdf.SetFont("Arial", "B", 11)
	pdf.SetFillColor(230, 230, 250)
	pdf.CellFormat(colWidths[0]+colWidths[1]+colWidths[2]+colWidths[3], 8, "TOTAL", "1", 0, "R", true, 0, "")
	pdf.CellFormat(colWidths[4], 8, "Rp. "+humanize.Comma(int64(data.Total)), "1", 1, "R", true, 0, "")

	// ===== TERBILANG =====
	pdf.Ln(8)
	pdf.SetFont("Arial", "I", 10)
	pdf.MultiCell(0, 6, "Terbilang : "+Terbilang(int64(data.Total)), "", "", false)

	// ===== BANK =====
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 6, "Bank : "+profile.NamaBank)
	pdf.Ln(6)
	pdf.Cell(0, 6, "No. Rekening : "+profile.NoRekening)
	pdf.Ln(6)
	pdf.Cell(0, 6, "a.n "+profile.NamaPT)

	// ===== TANDA TANGAN =====
	pdf.Ln(20)
	pdf.CellFormat(0, 6, profile.Kabupaten+", "+data.InvoiceDate, "", 1, "R", false, 0, "")
	pdf.Ln(20)
	pdf.CellFormat(0, 6, profile.PenanggungJawab, "", 1, "R", false, 0, "")

	return pdf
}
