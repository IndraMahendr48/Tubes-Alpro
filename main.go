package main

import (
	"fmt"
)

// Struct Pengeluaran untuk menyimpan data kategori dan jumlah pengeluaran
type Pengeluaran struct {
	Kategori string  // Nama kategori (Transportasi, Akomodasi, Makanan, Hiburan)
	Jumlah   float64 // Nominal pengeluaran dalam Rupiah
}

// Variabel global untuk melacak status pengurutan kategori (0: tidak diurutkan, 1: A-Z, 2: Z-A)
var isSortedByCategory int

// Fungsi untuk autentikasi login pengguna, mengembalikan true jika berhasil
func login() bool {
	pengguna := map[string]string{
		"traveler1": "jalan123",
		"traveler2": "libur456",
	}

	for kesempatan := 3; kesempatan > 0; kesempatan-- {
		fmt.Println("\n=== Login Aplikasi ===")
		fmt.Print("Masukkan username: ")
		var username string
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		var password string
		fmt.Scan(&password)

		if pass, ada := pengguna[username]; ada && pass == password {
			fmt.Println("Login berhasil! Selamat datang,", username)
			return true
		}
		fmt.Printf("Login gagal. Kesempatan tersisa: %d\n", kesempatan-1)
	}
	fmt.Println("Terlalu banyak percobaan. Aplikasi ditutup.")
	return false
}

// Prosedur untuk menampilkan menu aplikasi
func tampilkanMenu() {
	fmt.Println("\n=== Aplikasi Pengelolaan Anggaran Traveling ===")
	fmt.Println("1. Update Pengeluaran")
	fmt.Println("2. Cari Pengeluaran")
	fmt.Println("3. Urutkan Pengeluaran")
	fmt.Println("4. Lihat Laporan")
	fmt.Println("5. Keluar")
}

// Fungsi untuk membaca pilihan menu dari pengguna
func bacaPilihan() int {
	var pilihan int
	fmt.Print("Pilih menu (1-5): ")
	fmt.Scan(&pilihan)
	return pilihan
}

// Fungsi untuk membaca pilihan update pengeluaran dari pengguna
func bacaUpdate() int {
	var pilihan int
	fmt.Print("Pilih menu (1-4): ")
	fmt.Scan(&pilihan)
	return pilihan
}

// Prosedur untuk menampilkan menu update pengeluaran
func menuPengeluaran() {
	fmt.Println("\n=== Update Pengeluaran ===")
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Edit Pengeluaran")
	fmt.Println("3. Hapus Pengeluaran")
	fmt.Println("4. Balik ke Menu Utama")
}

// Fungsi untuk case apabila memilih menu pengeluaran
func updatePengeluaran(daftarPengeluaran *[]Pengeluaran) {
	for {
		menuPengeluaran()
		pilihan := bacaUpdate()

		switch pilihan {
		case 1:
			tambahPengeluaran(daftarPengeluaran)
		case 2:
			editPengeluaran(daftarPengeluaran)
		case 3:
			hapusPengeluaran(daftarPengeluaran)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan pilih lagi.")
		}
	}
}

// Fungsi untuk membaca nomor pengeluaran, mengembalikan -1 jika tidak valid
func bacaNomorPengeluaran(panjangDaftar int) int {
	var nomor int
	fmt.Print("Masukkan nomor pengeluaran: ")
	fmt.Scan(&nomor)
	if nomor < 1 || nomor > panjangDaftar {
		fmt.Println("Nomor tidak valid.")
		return -1
	}
	return nomor - 1 // Adjust to 0-based index
}

// Fungsi untuk membaca kategori dari pengguna dengan validasi
func bacaKategori(prompt string) string {
	validKategori := map[string]bool{"Transportasi": true, "Akomodasi": true, "Makanan": true, "Hiburan": true}
	var kategori string
	for {
		fmt.Print(prompt)
		fmt.Scan(&kategori)
		if validKategori[kategori] {
			return kategori
		}
		fmt.Println("Kategori tidak valid. Gunakan: Transportasi, Akomodasi, Makanan, atau Hiburan.")
	}
}

// Fungsi untuk membaca jumlah pengeluaran dari pengguna
func bacaJumlah() float64 {
	var jumlah float64
	for {
		fmt.Print("Masukkan jumlah (Rp): ")
		if fmt.Scan(&jumlah); jumlah >= 0 {
			return jumlah
		}
		fmt.Println("Jumlah tidak valid. Masukkan nilai positif.")
	}
}

// Prosedur untuk menampilkan daftar pengeluaran
func tampilkanDaftarPengeluaran(daftarPengeluaran []Pengeluaran) {
	fmt.Println("Daftar Pengeluaran:")
	if len(daftarPengeluaran) == 0 {
		fmt.Println("Tidak ada pengeluaran.")
		return
	}
	for i, pengeluaran := range daftarPengeluaran {
		fmt.Printf("%d. Kategori: %s, Jumlah: Rp %.2f\n", i+1, pengeluaran.Kategori, pengeluaran.Jumlah)
	}
}

// Prosedur untuk menambahkan pengeluaran baru
func tambahPengeluaran(daftarPengeluaran *[]Pengeluaran) {
	kategori := bacaKategori("Masukkan kategori (Transportasi/Akomodasi/Makanan/Hiburan): ")
	jumlah := bacaJumlah()
	*daftarPengeluaran = append(*daftarPengeluaran, Pengeluaran{Kategori: kategori, Jumlah: jumlah})
	isSortedByCategory = 0
	fmt.Println("Pengeluaran berhasil ditambahkan!")
}

// Prosedur untuk mengedit pengeluaran berdasarkan nomor
func editPengeluaran(daftarPengeluaran *[]Pengeluaran) {
	if len(*daftarPengeluaran) == 0 {
		fmt.Println("Belum ada pengeluaran untuk diedit.")
		return
	}
	tampilkanDaftarPengeluaran(*daftarPengeluaran)
	nomor := bacaNomorPengeluaran(len(*daftarPengeluaran))
	if nomor == -1 {
		return
	}
	jumlahBaru := bacaJumlah()
	(*daftarPengeluaran)[nomor].Jumlah = jumlahBaru
	isSortedByCategory = 0
	fmt.Println("Pengeluaran berhasil diedit!")
}

// Prosedur untuk menghapus pengeluaran berdasarkan nomor
func hapusPengeluaran(daftarPengeluaran *[]Pengeluaran) {
	if len(*daftarPengeluaran) == 0 {
		fmt.Println("Belum ada pengeluaran untuk dihapus.")
		return
	}
	tampilkanDaftarPengeluaran(*daftarPengeluaran)
	nomor := bacaNomorPengeluaran(len(*daftarPengeluaran))
	if nomor == -1 {
		return
	}
	*daftarPengeluaran = append((*daftarPengeluaran)[:nomor], (*daftarPengeluaran)[nomor+1:]...)
	isSortedByCategory = 0
	fmt.Println("Pengeluaran berhasil dihapus!")
}

// Fungsi Sequential Search untuk mencari semua entri dengan kategori tertentu
func sequentialSearch(daftarPengeluaran []Pengeluaran, kategori string) []int {
	var indeks []int
	for i, pengeluaran := range daftarPengeluaran {
		if pengeluaran.Kategori == kategori {
			indeks = append(indeks, i)
		}
	}
	return indeks
}

// Fungsi Binary Search untuk mencari indeks pertama dengan kategori tertentu
func binarySearch(daftarPengeluaran []Pengeluaran, kategori string, aToZ bool) int {
	kiri, kanan := 0, len(daftarPengeluaran)-1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
		if daftarPengeluaran[tengah].Kategori == kategori {
			// Cari indeks pertama dengan kategori ini
			for tengah > 0 && daftarPengeluaran[tengah-1].Kategori == kategori {
				tengah--
			}
			return tengah
		}
		if (aToZ && daftarPengeluaran[tengah].Kategori < kategori) || (!aToZ && daftarPengeluaran[tengah].Kategori > kategori) {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

// Prosedur untuk mencari pengeluaran berdasarkan kategori
func cariPengeluaran(daftarPengeluaran []Pengeluaran) {
	if len(daftarPengeluaran) == 0 {
		fmt.Println("Belum ada pengeluaran untuk dicari.")
		return
	}
	kategori := bacaKategori("Masukkan kategori yang ingin dicari (Transportasi/Akomodasi/Makanan/Hiburan): ")
	fmt.Println("Hasil Pencarian:")
	ditemukan := false
	var indeks []int

	if isSortedByCategory == 1 || isSortedByCategory == 2 {
		fmt.Println("Menggunakan Binary Search (data telah diurutkan)")
		indeksPertama := binarySearch(daftarPengeluaran, kategori, isSortedByCategory == 1)
		if indeksPertama != -1 {
			for i := indeksPertama; i < len(daftarPengeluaran) && daftarPengeluaran[i].Kategori == kategori; i++ {
				indeks = append(indeks, i)
			}
		}
	} else {
		fmt.Println("Menggunakan Sequential Search (data belum diurutkan)")
		indeks = sequentialSearch(daftarPengeluaran, kategori)
	}

	for _, i := range indeks {
		fmt.Printf("%d. Kategori: %s, Jumlah: Rp %.2f\n", i+1, daftarPengeluaran[i].Kategori, daftarPengeluaran[i].Jumlah)
		ditemukan = true
	}
	if !ditemukan {
		fmt.Println("Tidak ada pengeluaran untuk kategori tersebut.")
	}
}

// Prosedur Bubble Sort untuk mengurutkan daftar pengeluaran
func urutkanPengeluaran(daftarPengeluaran *[]Pengeluaran, kriteria int) {
	n := len(*daftarPengeluaran)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			tukar := false
			switch kriteria {
			case 1: // Jumlah: Besar ke Kecil
				if (*daftarPengeluaran)[j].Jumlah < (*daftarPengeluaran)[j+1].Jumlah {
					tukar = true
				}
			case 2: // Jumlah: Kecil ke Besar
				if (*daftarPengeluaran)[j].Jumlah > (*daftarPengeluaran)[j+1].Jumlah {
					tukar = true
				}
			case 3: // Kategori: A-Z
				if (*daftarPengeluaran)[j].Kategori > (*daftarPengeluaran)[j+1].Kategori {
					tukar = true
				}
			case 4: // Kategori: Z-A
				if (*daftarPengeluaran)[j].Kategori < (*daftarPengeluaran)[j+1].Kategori {
					tukar = true
				}
			}
			if tukar {
				(*daftarPengeluaran)[j], (*daftarPengeluaran)[j+1] = (*daftarPengeluaran)[j+1], (*daftarPengeluaran)[j]
			}
		}
	}
	if kriteria == 3 {
		isSortedByCategory = 1 // A-Z
	} else if kriteria == 4 {
		isSortedByCategory = 2 // Z-A
	} else {
		isSortedByCategory = 0 // Tidak diurutkan berdasarkan kategori
	}
}

// Prosedur untuk mengelola pengurutan pengeluaran
func kelolaPengurutan(daftarPengeluaran *[]Pengeluaran) {
	if len(*daftarPengeluaran) == 0 {
		fmt.Println("Belum ada pengeluaran untuk diurutkan.")
		return
	}
	fmt.Println("Pilih kriteria pengurutan:")
	fmt.Println("1. Jumlah (Besar ke Kecil)")
	fmt.Println("2. Jumlah (Kecil ke Besar)")
	fmt.Println("3. Kategori (A-Z)")
	fmt.Println("4. Kategori (Z-A)")
	var pilihanUrut int
	fmt.Scan(&pilihanUrut)

	if pilihanUrut < 1 || pilihanUrut > 4 {
		fmt.Println("Pilihan tidak valid.")
		return
	}
	urutkanPengeluaran(daftarPengeluaran, pilihanUrut)
	fmt.Println("Pengeluaran berhasil diurutkan!")
	tampilkanDaftarPengeluaran(*daftarPengeluaran)
}

// Prosedur untuk menampilkan laporan pengeluaran
func tampilkanLaporan(daftarPengeluaran []Pengeluaran, anggaran float64) {
	if len(daftarPengeluaran) == 0 {
		fmt.Println("Belum ada pengeluaran untuk dilaporkan.")
		return
	}
	fmt.Println("\n=== Laporan Pengeluaran ===")
	totalPerKategori := map[string]float64{
		"Transportasi": 0,
		"Akomodasi":    0,
		"Makanan":      0,
		"Hiburan":      0,
	}
	var totalPengeluaran float64
	for _, pengeluaran := range daftarPengeluaran {
		totalPerKategori[pengeluaran.Kategori] += pengeluaran.Jumlah
		totalPengeluaran += pengeluaran.Jumlah
	}
	fmt.Println("Rincian per Kategori:")
	for kategori, jumlah := range totalPerKategori {
		if jumlah > 0 {
			fmt.Printf("- %s: Rp %.2f\n", kategori, jumlah)
		}
	}
	fmt.Printf("Total Pengeluaran: Rp %.2f\n", totalPengeluaran)
	fmt.Printf("Total Anggaran: Rp %.2f\n", anggaran)
	selisih := anggaran - totalPengeluaran
	if selisih > 0 {
		fmt.Printf("Sisa Anggaran: Rp %.2f\n", selisih)
	} else if selisih < 0 {
		fmt.Printf("Kelebihan Pengeluaran: Rp %.2f\n", -selisih)
	}
	fmt.Println("Rekomendasi Penghematan:")
	if totalPengeluaran > anggaran {
		fmt.Println("- Kurangi pengeluaran pada kategori dengan jumlah terbesar.")
		kategoriTerbesar, jumlahTerbesar := "", 0.0
		for kategori, jumlah := range totalPerKategori {
			if jumlah > jumlahTerbesar {
				jumlahTerbesar = jumlah
				kategoriTerbesar = kategori
			}
		}
		fmt.Printf("- Fokus hemat pada %s (Rp %.2f).\n", kategoriTerbesar, jumlahTerbesar)
	} else {
		fmt.Println("- Pertahankan pola pengeluaran Anda, masih ada sisa anggaran.")
	}
}

func main() {
	if !login() {
		return
	}

	var daftarPengeluaran []Pengeluaran
	var anggaran float64
	fmt.Print("Masukkan anggaran awal (Rp): ")
	fmt.Scan(&anggaran)

	for {
		tampilkanMenu()
		pilihan := bacaPilihan()

		switch pilihan {
		case 1:
			updatePengeluaran(&daftarPengeluaran)
		case 2:
			cariPengeluaran(daftarPengeluaran)
		case 3:
			kelolaPengurutan(&daftarPengeluaran)
		case 4:
			tampilkanLaporan(daftarPengeluaran, anggaran)
		case 5:
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan pilih lagi.")
		}
	}
}