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
	// Simulasi database pengguna dengan map (username: password)
	pengguna := map[string]string{
		"traveler1": "jalan123",
		"traveler2": "libur456",
	}

	// Berikan 3 kesempatan login
	for kesempatan := 3; kesempatan > 0; kesempatan-- {
		fmt.Println("\n=== Login Aplikasi ===")
		fmt.Print("Masukkan username: ")
		var username string
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		var password string
		fmt.Scan(&password)

		// Cek apakah username dan password cocok
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
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Edit Pengeluaran")
	fmt.Println("3. Hapus Pengeluaran")
	fmt.Println("4. Cari Pengeluaran")
	fmt.Println("5. Urutkan Pengeluaran")
	fmt.Println("6. Lihat Laporan")
	fmt.Println("7. Keluar")
}

// Fungsi untuk membaca pilihan menu dari pengguna
func bacaPilihan() int {
	fmt.Print("Pilih menu (1-7): ")
	var pilihan int
	fmt.Scan(&pilihan)
	return pilihan
}

// Fungsi untuk membaca nomor pengeluaran, mengembalikan -1 jika tidak valid
func bacaNomorPengeluaran(panjangDaftar int) int {
	fmt.Print("Masukkan nomor pengeluaran: ")
	var nomor int
	fmt.Scan(&nomor)
	if nomor < 1 || nomor > panjangDaftar {
		fmt.Println("Nomor tidak valid.")
		return -1
	}
	return nomor
}

// Fungsi untuk membaca kategori dari pengguna
func bacaKategori(prompt string) string {
	fmt.Print(prompt)
	var kategori string
	fmt.Scan(&kategori)
	return kategori
}

// Fungsi untuk membaca jumlah pengeluaran dari pengguna
func bacaJumlah() float64 {
	fmt.Print("Masukkan jumlah (Rp): ")
	var jumlah float64
	fmt.Scan(&jumlah)
	return jumlah
}

// Prosedur untuk menampilkan daftar pengeluaran
func tampilkanDaftarPengeluaran(daftarPengeluaran []Pengeluaran) {
	fmt.Println("Daftar Pengeluaran:")
	for i, pengeluaran := range daftarPengeluaran {
		fmt.Printf("%d. Kategori: %s, Jumlah: Rp %.2f\n", i+1, pengeluaran.Kategori, pengeluaran.Jumlah)
	}
}

// Prosedur untuk menambahkan pengeluaran baru
func tambahPengeluaran(daftarPengeluaran *[]Pengeluaran) {
	// Minta kategori dan jumlah
	kategori := bacaKategori("Masukkan kategori (Transportasi/Akomodasi/Makanan/Hiburan): ")
	jumlah := bacaJumlah()
	// Tambahkan data ke daftar pengeluaran
	*daftarPengeluaran = append(*daftarPengeluaran, Pengeluaran{Kategori: kategori, Jumlah: jumlah})
	// Tandai bahwa data tidak lagi diurutkan
	isSortedByCategory = 0
	fmt.Println("Pengeluaran berhasil ditambahkan!")
}

// Prosedur untuk mengedit pengeluaran berdasarkan nomor
func editPengeluaran(daftarPengeluaran *[]Pengeluaran) {
	if len(*daftarPengeluaran) == 0 {
		fmt.Println("Belum ada pengeluaran untuk diedit.")
		return
	}
	// Tampilkan daftar pengeluaran
	tampilkanDaftarPengeluaran(*daftarPengeluaran)
	// Minta nomor pengeluaran
	nomor := bacaNomorPengeluaran(len(*daftarPengeluaran))
	if nomor == -1 {
		return
	}
	// Minta data baru
	jumlahBaru := bacaJumlah()
	// Perbarui data pengeluaran
	(*daftarPengeluaran)[nomor-1].Jumlah = jumlahBaru
	// Tandai bahwa data tidak lagi diurutkan
	isSortedByCategory = 0
	fmt.Println("Pengeluaran berhasil diedit!")
}

// Prosedur untuk menghapus pengeluaran berdasarkan nomor
func hapusPengeluaran(daftarPengeluaran *[]Pengeluaran) {
	if len(*daftarPengeluaran) == 0 {
		fmt.Println("Belum ada pengeluaran untuk dihapus.")
		return
	}
	// Tampilkan daftar pengeluaran
	tampilkanDaftarPengeluaran(*daftarPengeluaran)
	// Minta nomor pengeluaran
	nomor := bacaNomorPengeluaran(len(*daftarPengeluaran))
	if nomor == -1 {
		return
	}
	// Hapus pengeluaran dari daftar
	*daftarPengeluaran = append((*daftarPengeluaran)[:nomor-1], (*daftarPengeluaran)[nomor:]...)
	// Tandai bahwa data tidak lagi diurutkan
	isSortedByCategory = 0
	fmt.Println("Pengeluaran berhasil dihapus!")
}

// Fungsi Sequential Search untuk mencari semua entri dengan kategori tertentu
func sequentialSearch(daftarPengeluaran []Pengeluaran, kategori string) []int {
	var indeks []int
	// Periksa setiap elemen secara berurutan
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
		if (aToZ && daftarPengeluaran[tengah].Kategori == kategori) || (!aToZ && daftarPengeluaran[tengah].Kategori == kategori) {
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
	// Minta kategori yang akan dicari
	kategori := bacaKategori("Masukkan kategori yang ingin dicari (Transportasi/Akomodasi/Makanan/Hiburan): ")
	fmt.Println("Hasil Pencarian:")
	ditemukan := false
	var indeks []int

	// Pilih metode pencarian berdasarkan status pengurutan
	if isSortedByCategory == 1 || isSortedByCategory == 2 {
		// Gunakan Binary Search jika data diurutkan berdasarkan kategori
		fmt.Println("Menggunakan Binary Search (data telah diurutkan)")
		indeksPertama := binarySearch(daftarPengeluaran, kategori, isSortedByCategory == 1)
		if indeksPertama != -1 {
			// Tambahkan semua entri dengan kategori yang sama
			for i := indeksPertama; i < len(daftarPengeluaran) && daftarPengeluaran[i].Kategori == kategori; i++ {
				indeks = append(indeks, i)
			}
		}
	} else {
		// Gunakan Sequential Search jika data belum diurutkan
		fmt.Println("Menggunakan Sequential Search (data belum diurutkan)")
		indeks = sequentialSearch(daftarPengeluaran, kategori)
	}

	// Tampilkan hasil pencarian
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
	// Bubble Sort: bandingkan dan tukar elemen berpasangan
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
			// Tukar elemen jika diperlukan
			if tukar {
				(*daftarPengeluaran)[j], (*daftarPengeluaran)[j+1] = (*daftarPengeluaran)[j+1], (*daftarPengeluaran)[j]
			}
		}
	}
}

// Prosedur untuk mengelola pengurutan pengeluaran
func kelolaPengurutan(daftarPengeluaran *[]Pengeluaran) {
	if len(*daftarPengeluaran) == 0 {
		fmt.Println("Belum ada pengeluaran untuk diurutkan.")
		return
	}
	// Tampilkan opsi pengurutan
	fmt.Println("Pilih kriteria pengurutan:")
	fmt.Println("1. Jumlah (Besar ke Kecil)")
	fmt.Println("2. Jumlah (Kecil ke Besar)")
	fmt.Println("3. Kategori (A-Z)")
	fmt.Println("4. Kategori (Z-A)")
	fmt.Print("Pilih (1-4): ")
	var pilihanUrut int
	fmt.Scan(&pilihanUrut)

	// Validasi pilihan dan panggil fungsi pengurutan
	if pilihanUrut < 1 || pilihanUrut > 4 {
		fmt.Println("Pilihan tidak valid.")
		return
	}
	urutkanPengeluaran(daftarPengeluaran, pilihanUrut)
	// Perbarui status pengurutan
	if pilihanUrut == 3 {
		isSortedByCategory = 1 // A-Z
	} else if pilihanUrut == 4 {
		isSortedByCategory = 2 // Z-A
	} else {
		isSortedByCategory = 0 // Tidak diurutkan berdasarkan kategori
	}
	// Tampilkan hasil pengurutan
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
	// Inisialisasi map untuk menyimpan total per kategori
	totalPerKategori := map[string]float64{
		"Transportasi": 0,
		"Akomodasi":   0,
		"Makanan":     0,
		"Hiburan":     0,
	}
	totalPengeluaran := 0.0
	// Hitung total pengeluaran per kategori
	for _, pengeluaran := range daftarPengeluaran {
		totalPerKategori[pengeluaran.Kategori] += pengeluaran.Jumlah
		totalPengeluaran += pengeluaran.Jumlah
	}
	// Tampilkan rincian per kategori
	fmt.Println("Rincian per Kategori:")
	for kategori, jumlah := range totalPerKategori {
		if jumlah > 0 {
			fmt.Printf("- %s: Rp %.2f\n", kategori, jumlah)
		}
	}
	// Tampilkan total pengeluaran dan anggaran
	fmt.Printf("Total Pengeluaran: Rp %.2f\n", totalPengeluaran)
	fmt.Printf("Total Anggaran: Rp %.2f\n", anggaran)
	// Hitung selisih anggaran
	selisih := anggaran - totalPengeluaran
	if selisih > 0 {
		fmt.Printf("Sisa Anggaran: Rp %.2f\n", selisih)
	} else {
		fmt.Printf("Kelebihan Pengeluaran: Rp %.2f\n", -selisih)
	}
	// Berikan rekomendasi penghematan
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
	// Meminta pengguna untuk login
	if !login() {
		return
	}

	// Deklarasi variabel untuk menyimpan daftar pengeluaran dan anggaran awal
	var daftarPengeluaran []Pengeluaran
	var anggaran float64

	// Meminta pengguna memasukkan anggaran awal
	fmt.Print("Masukkan anggaran awal (Rp): ")
	fmt.Scan(&anggaran)

	// Loop utama untuk menampilkan menu dan memproses pilihan
	for {
<<<<<<< Updated upstream
		tampilkanMenu()
		pilihan := bacaPilihan()
=======
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
>>>>>>> Stashed changes

		switch pilihan {
		case 1:
			tambahPengeluaran(&daftarPengeluaran)
		case 2:
			editPengeluaran(&daftarPengeluaran)
		case 3:
			hapusPengeluaran(&daftarPengeluaran)
		case 4:
			cariPengeluaran(daftarPengeluaran)
		case 5:
			kelolaPengurutan(&daftarPengeluaran)
		case 6:
			tampilkanLaporan(daftarPengeluaran, anggaran)
		case 7:
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan pilih lagi.")
		}
	}
}
