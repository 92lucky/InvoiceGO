package service

import (
	"fmt"
	"invoice-go/config"
	"invoice-go/internal/auth"
	"invoice-go/internal/model"
	"invoice-go/internal/repository"
	"invoice-go/utils"
	"net/http"
	"strconv"
)

func ServiceInvoice(r *http.Request) (model.InvoiceData, error) {
	err := r.ParseForm()
	if err != nil {
		return model.InvoiceData{}, err
	}

	qty, _ := strconv.ParseFloat(r.FormValue("quantity_kg"), 64)

	// Kirim 0 dulu ke HitungTagihan kalau memang butuh param dpp
	displayQty, pokok, ppn, _, total := utils.HitungTagihan(qty)



	data := model.InvoiceData{
		InvoiceNumber: r.FormValue("invoice_number"),
		InvoiceDate:   r.FormValue("invoice_date"),
		Periode:       r.FormValue("periode"),
		QuantityKG:    qty,
		DisplayQty:    displayQty,
		Pokok:         pokok,
		PPN:           ppn,
		Total:         total,
	}

	return data, nil
}
func GenerateInvoicePDF(w http.ResponseWriter, r *http.Request, isDownload bool) error {
	// Ambil session
	session, _ := auth.GetSession(r)
	emailVal, ok := session.Values["email"]
	if !ok {
		return fmt.Errorf("email tidak ditemukan di session")
	}

	email, ok := emailVal.(string)
	if !ok || email == "" {
		return fmt.Errorf("email di session invalid")
	}

	// Ambil profile dari repository
	profile, err := repository.GetUserEmail(config.DB, email)
	if err != nil || profile == nil {
		return fmt.Errorf("profil belum diisi")
	}

	// Parse form
	if err := r.ParseMultipartForm(10 << 20); err != nil  {
		return fmt.Errorf("gagal parsing form: %v", err)
	}

	// Ambil data invoice dari form
	invoiceNumber := r.FormValue("invoice_number")
	invoiceDate := r.FormValue("invoice_date")
	periode := r.FormValue("periode")
	qty, _ := strconv.ParseFloat(r.FormValue("quantity_kg"), 64)
	

	// Hitung tagihan
	displayQty, pokok, ppn, dpp, total := utils.HitungTagihan(qty)

	data := model.InvoiceData{
		InvoiceNumber: invoiceNumber,
		InvoiceDate:   invoiceDate,
		Periode:       periode,
		QuantityKG:    qty,
		DisplayQty:    displayQty,
		DPP:           dpp,
		Pokok:         pokok,
		PPN:           ppn,
		Total:         total,
	}

	// Generate PDF → dereference pointer karena GeneratePDFInvoice menerima value
	pdf := utils.GeneratePDFInvoice(*profile, data)

	// Set header
	w.Header().Set("Content-Type", "application/pdf")
	if isDownload {
		w.Header().Set("Content-Disposition", "attachment; filename=invoice.pdf")
	} else {
		w.Header().Set("Content-Disposition", "inline; filename=invoice.pdf")
	}

	return pdf.Output(w)
}
