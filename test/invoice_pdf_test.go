package test

import (
	"bytes"
	"invoice-go/internal/handlers"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandleGeneratePDF_Integration(t *testing.T) {
	// === Persiapan file dummy.xlsx ===
	file, err := os.Open("../static/dummy.xlsx")
	require.NoError(t, err, "Gagal membuka file dummy")
	defer file.Close()

	var buf bytes.Buffer
	writer := multipart.NewWriter(&buf)

	formFile, err := writer.CreateFormFile("file", "dummy.xlsx")
	require.NoError(t, err, "Gagal membuat form file")

	_, err = io.Copy(formFile, file)
	require.NoError(t, err, "Gagal menyalin isi file")

	_ = writer.WriteField("tanggal", "2025-07-17")
	_ = writer.WriteField("berat", "100")
	writer.Close()

	// === Simulasi request ===
	req := httptest.NewRequest(http.MethodPost, "/generate", &buf)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()
	handlers.HandleGeneratePDF(rr, req, false)

	// === Assertion ===
	assert.Equal(t, http.StatusOK, rr.Code, "Status harus 200 OK")
	assert.Equal(t, "application/pdf", rr.Header().Get("Content-Type"), "Content-Type harus application/pdf")

	body := rr.Body.Bytes()
	assert.Greater(t, len(body), 100, "PDF hasil harus memiliki isi (lebih dari 100 byte)")
}
