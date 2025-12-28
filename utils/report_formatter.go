package utils

import (
	"mime/multipart"
	"net/http"
	"time"
)

//report parsing
type ReportForm struct {
	NamaPT string
	Bulan  string
	File   multipart.File
}

func ParseReportForm(r *http.Request) (ReportForm, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return ReportForm{}, err
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		return ReportForm{}, err
	}

	namaPT := r.FormValue("namapt")
	bulan := r.FormValue("bulan")
	if bulan == "" {
		bulan = time.Now().Format("January 2006")
	}

	return ReportForm{
		NamaPT: namaPT,
		Bulan:  bulan,
		File:   file,
	}, nil
}
