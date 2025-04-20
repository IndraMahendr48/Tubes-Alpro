package main

import (
	"fmt"
)

func main() {
	// Simpan data pengeluaran sebagai slice of slice (kategori, jumlah)
	var expenses [][]string
	var budget float64

	// masukan budget awal
	fmt.Print("Masukkan budget awal (Rp): ")
	fmt.Scan(&budget)

	for {
		// Tampilan Menu
		fmt.Println("\n=== Aplikasi Pengelolaan Budget Traveling ===")
		fmt.Println("1. Tambah Pengeluaran")
		fmt.Println("2. Edit Pengeluaran")
		fmt.Println("3. Hapus Pengeluaran")
		fmt.Println("4. Cari Pengeluaran")
		fmt.Println("5. Urutkan Pengeluaran")
		fmt.Println("6. Lihat Laporan")
		fmt.Println("7. Keluar")
		fmt.Print("Pilih menu (1-7): ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1: 
			var category string
			var amount float64
			fmt.Print("Masukkan kategori (Transportasi/Akomodasi/Makanan/Hiburan): ")
			fmt.Scan(&category)
			fmt.Print("Masukkan jumlah (Rp): ")
			fmt.Scan(&amount)
			expenses = append(expenses, []string{category, fmt.Sprintf("%.2f", amount)})
			fmt.Println("Pengeluaran berhasil ditambahkan!")

		case 2, 3, 4, 5, 6, 7:
			fmt.Println("Fitur belum diimplementasikan.")

		default:
			fmt.Println("Pilihan tidak valid, silakan pilih lagi.")
		}
	}
}