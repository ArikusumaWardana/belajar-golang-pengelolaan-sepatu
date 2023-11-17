package main

import "fmt"

type ukuran struct {
	us int
	eu int
	id int
	jp int
}

type sepatu struct {
	merk   string
	harga  int
	margin int
	warna  []string
	ukuran ukuran
	stok   int
}

func (s sepatu) hargaJual() int {
	return s.harga + s.margin
}

func (s sepatu) totalDuit() int {
	return s.hargaJual() * s.stok
}

func main() {

	var daftarSepatu = []sepatu{
		{
			merk:   "Nike",
			harga:  250000000,
			margin: 100000,
			warna:  []string{"merah", "indigo", "putih"},
			ukuran: ukuran{
				us: 10,
				eu: 42,
				id: 42,
				jp: 23,
			},
			stok: 5,
		},
		{
			merk:   "Adidas",
			harga:  10000000,
			margin: 2000000,
			warna:  []string{"biru", "hijau", "kuning"},
			ukuran: ukuran{
				us: 11,
				eu: 43,
				id: 44,
				jp: 24,
			},
			stok: 3,
		},
	}

	for _, s := range daftarSepatu {
		fmt.Println("Merk:", s.merk)
		fmt.Println("Harga:", s.harga)
		fmt.Println("Ukuran US:", s.ukuran.us)
		fmt.Println("Ukuran EU:", s.ukuran.eu)
		fmt.Println("Ukuran ID:", s.ukuran.id)
		fmt.Println("Ukuran JP:", s.ukuran.jp)
		fmt.Println("Harga Jual:", s.hargaJual())
		fmt.Println("Total Uang:", s.totalDuit())
	}
}
