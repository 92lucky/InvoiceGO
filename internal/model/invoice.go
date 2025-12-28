package model

type InvoiceData struct {
	InvoiceNumber string
	InvoiceDate   string
	Periode       string
	QuantityKG    float64
	DisplayQty    float64
	Pokok         float64
	DPP           float64
	PPN           float64
	Total         float64
}
