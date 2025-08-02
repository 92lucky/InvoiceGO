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

	// ===== KE PADA =====
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 6, "Kepada : PT. Pertamina Patra Niaga")
	pdf.Ln(6)
	pdf.Cell(0, 6, "Alamat : Gedung Wisma Tugu II Lt.2, Jl. HR Rasuna Said KAV C7-9 Setiabudi, Jakarta 12920")
	pdf.Ln(8)

	// ===== TANGGAL & NO INVOICE =====
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(95, 6, "Tanggal : "+data.InvoiceDate, "", 0, "L", false, 0, "")
	pdf.CellFormat(0, 6, "No. Invoice : "+data.InvoiceNumber, "", 1, "R", false, 0, "")
	pdf.Ln(8)

	colWidths := []float64{10, 80, 20, 40, 40} // No | Keterangan | Kg | Harga Satuan | Jumlah

	pdf.SetFont("Arial", "B", 11)
	pdf.SetFillColor(220, 220, 220)
	headers := []string{"No", "Keterangan", "Kg", "Harga Satuan", "Jumlah"}
	for i, h := range headers {
		pdf.CellFormat(colWidths[i], 8, h, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// ===== ISI TABEL =====
	pdf.SetFont("Arial", "", 10)
	rows := []struct {
		No    string
		Ket   string
		Kg    string
		Harga string
		Jml   string
	}{
		{"1", "Tagihan Transport Fee LPG 3 Kg Periode " + data.Periode, humanize.Comma(int64(data.DisplayQty)), "-", ""},
		{"2", "Pokok", "", "Rp. " + humanize.Comma(int64(data.Pokok)), "Rp. " + humanize.Comma(int64(data.Pokok))},
		{"3", "PPN 12%", "", "-", "Rp. " + humanize.Comma(int64(data.PPN))},
	}

	for _, row := range rows {
		cells := []string{row.No, row.Ket, row.Kg, row.Harga, row.Jml}
		aligns := []string{"C", "L", "C", "R", "R"}
		for i, c := range cells {
			pdf.CellFormat(colWidths[i], 8, c, "1", 0, aligns[i], false, 0, "")
		}
		pdf.Ln(-1)
	}

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
