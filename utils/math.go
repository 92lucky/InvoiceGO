package utils

import (
	"math"
	"strings"
)

//invoice count
func HitungTagihan(qtyKg float64) (displayQty, pokok, ppn, dppOut, total float64) {
	const (
		hargaSatuan  = 560.0
		pengali      = 3.0
		hargaPerGram = 354.64
	)

	displayQty = qtyKg * hargaSatuan * pengali

	// Bulatkan pokok dulu
	pokok = math.Round(displayQty * hargaPerGram)

	// Hitung DPP dari pokok yang sudah bulat
	dppOut = math.Round(pokok * 11.0 / 12.0)

	// Hitung PPN dari DPP
	ppn = math.Round(dppOut * 0.12)

	total = pokok + ppn
	return
}

func Terbilang(n int64) string {
	angka := [...]string{"", "Satu", "Dua", "Tiga", "Empat", "Lima", "Enam", "Tujuh", "Delapan", "Sembilan"}
	satuan := [...]string{"", "Ribu", "Juta", "Miliar", "Triliun"}

	if n == 0 {
		return "Nol Rupiah"
	}

	var hasil string
	var i int
	for n > 0 {
		sisa := n % 1000
		if sisa > 0 {
			var segmen string
			ratus := sisa / 100
			puluhan := (sisa % 100) / 10
			satuanAngka := sisa % 10

			if ratus > 0 {
				if ratus == 1 {
					segmen += "Seratus "
				} else {
					segmen += angka[ratus] + " Ratus "
				}
			}

			if puluhan > 0 {
				if puluhan == 1 {
					if satuanAngka == 0 {
						segmen += "Sepuluh "
					} else if satuanAngka == 1 {
						segmen += "Sebelas "
					} else {
						segmen += angka[satuanAngka] + " Belas "
					}
				} else {
					segmen += angka[puluhan] + " Puluh "
					if satuanAngka > 0 {
						segmen += angka[satuanAngka] + " "
					}
				}
			} else if satuanAngka > 0 {
				if satuanAngka == 1 && i == 1 && ratus == 0 && puluhan == 0 {
					segmen += "Seribu "
				} else {
					segmen += angka[satuanAngka] + " "
				}
			}

			segmen += satuan[i] + " "
			hasil = segmen + hasil
		}
		n /= 1000
		i++
	}

	return strings.TrimSpace(hasil) + " Rupiah"
}
