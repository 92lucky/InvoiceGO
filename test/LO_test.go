package test

import (
	"bytes"
	"invoice-go/internal/service"
	"io"
	"mime/multipart"
	"os"
	"testing"

	"github.com/xuri/excelize/v2"
)

// helper buat Excel di memory
func createTestExcel() (*bytes.Buffer, error) {
	f := excelize.NewFile()
	sheet := f.GetSheetName(0)

	// header
	f.SetCellValue(sheet, "A1", "No")
	f.SetCellValue(sheet, "B1", "Date")
	f.SetCellValue(sheet, "C1", "NoSO")
	f.SetCellValue(sheet, "D1", "NoLO")

	// data baris 1
	f.SetCellValue(sheet, "A2", 1)
	f.SetCellValue(sheet, "B2", "2025-08-29")
	f.SetCellValue(sheet, "C2", "SO123")
	f.SetCellValue(sheet, "D2", "LO456")

	var buf bytes.Buffer
	if err := f.Write(&buf); err != nil {
		return nil, err
	}
	return &buf, nil
}

// helper simulasikan multipart.File dari buffer
func bufferToMultipartFile(buf *bytes.Buffer) (multipart.File, error) {
	tmp, err := os.CreateTemp("", "test-*.xlsx")
	if err != nil {
		return nil, err
	}

	if _, err := tmp.Write(buf.Bytes()); err != nil {
		return nil, err
	}
	if _, err := tmp.Seek(0, io.SeekStart); err != nil {
		return nil, err
	}

	return tmp, nil
}

func TestParseExcelToDataRow(t *testing.T) {
	buf, err := createTestExcel()
	if err != nil {
		t.Fatalf("failed to create test excel: %v", err)
	}

	file, err := bufferToMultipartFile(buf)
	if err != nil {
		t.Fatalf("failed to create multipart file: %v", err)
	}
	defer file.Close()
	defer os.Remove(file.(*os.File).Name())

	data, err := service.ParseExcelToDataRow(file)
	if err != nil {
		t.Fatalf("ParseExcelToDataRow error: %v", err)
	}

	if len(data) != 1 {
		t.Fatalf("expected 1 row, got %d", len(data))
	}

	row := data[0]
	if row.No != 1 || row.NoSO != "SO123" || row.NoLO != "LO456" {
		t.Errorf("unexpected data: %+v", row)
	}
}

func TestParseExcelToDataRows(t *testing.T) {
	buf, err := createTestExcel()
	if err != nil {
		t.Fatalf("failed to create test excel: %v", err)
	}

	file, err := bufferToMultipartFile(buf)
	if err != nil {
		t.Fatalf("failed to create multipart file: %v", err)
	}
	defer file.Close()
	defer os.Remove(file.(*os.File).Name())

	data, err := service.ParseExcelToDataRow(file)
	if err != nil {
		t.Fatalf("ParseExcelToDataRows error: %v", err)
	}

	if len(data) != 1 {
		t.Fatalf("expected 1 row, got %d", len(data))
	}

	row := data[0]
	if row.No != 1 || row.NoSO != "SO123" || row.NoLO != "LO456" {
		t.Errorf("unexpected data: %+v", row)
	}
}
