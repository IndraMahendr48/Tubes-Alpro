package main

import (
	"fmt"
)

// Konstanta untuk ukuran maksimum array
const MAX_PENGELUARAN = 100

// Struct Pengeluaran untuk menyimpan data kategori dan jumlah pengeluaran
type Pengeluaran struct {
	Kategori string  // Nama kategori (Transportasi, Akomodasi, Makanan, Hiburan)
	Jumlah   float64 // Nominal pengeluaran dalam Rupiah
}

// Struct untuk menyimpan daftar pengeluaran
type DaftarPengeluaran struct {
	Data [MAX_PENGELUARAN]Pengeluaran
	N    int // Jumlah elemen yang terisi
}

// Variabel global untuk melacak status pengurutan kategori (0: tidak diurutkan, 1: A-Z, 2: Z-A)
var isSortedByCategory int

// Fungsi untuk autentikasi login pengguna, mengembalikan true jika berhasil
func login() bool {
	var pengguna [2]string = [2]string{"traveler1", "traveler2"}
	var kataSandi [2]string = [2]string{"jalan123", "libur456"}

	for kesempatan := 3; kesempatan > 0; kesempatan-- {
		fmt.Println("\n=== Login Aplikasi ===")
		fmt.Print("Masukkan username: ")
		var username string
		fmt.Scan(&username)
		fmt.Print("Masukkan password: ")
		var password string
		fmt.Scan(&password)

		berhasil := false
		for i := 0; i < 2; i++ {
			if username == pengguna[i] && password == kataSandi[i] {
				berhasil = true
			}
		}
		if berhasil {
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
	fmt.Println("5. Cari Pengeluaran Ekstrem")
	fmt.Println("6. Keluar")
}

// Fungsi untuk membaca pilihan menu dari pengguna
func bacaPilihan() int {
	var pilihan int
	fmt.Print("Pilih menu (1-6): ")
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

// Fungsi untuk membaca pilihan update pengeluaran dari pengguna
func bacaUpdate() int {
	var pilihan int
	fmt.Print("Pilih menu (1-4): ")
	fmt.Scan(&pilihan)
	return pilihan
}

// Fungsi untuk validasi kategori
func isValidKategori(kategori string) bool {
	if kategori == "Transportasi" || kategori == "Akomodasi" || kategori == "Makanan" || kategori == "Hiburan" {
		return true
	}
	return false
}

// Fungsi untuk membaca kategori dari pengguna dengan validasi
func bacaKategori(prompt string) string {
	var kategori string
	for {
		fmt.Print(prompt)
		fmt.Scan(&kategori)
		if isValidKategori(kategori) {
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
		// Ambil input jumlah dari pengguna
		fmt.Scan(&jumlah)
		// Periksa apakah jumlah lebih besar atau sama dengan 0
		if jumlah >= 0 {
			return jumlah
		}
		fmt.Println("Jumlah tidak valid. Masukkan nilai positif.")
		var dummy string
		fmt.Scan(&dummy) // Bersihkan input yang salah
	}
}

// Fungsi untuk membaca nomor pengeluaran, mengembalikan -1 jika tidak valid
func bacaNomorPengeluaran(n int) int {
	var nomor int
	fmt.Print("Masukkan nomor pengeluaran: ")
	// Ambil input nomor dari pengguna
	fmt.Scan(&nomor)
	// Periksa apakah nomor valid (antara 1 dan n)
	if nomor < 1 || nomor > n {
		fmt.Println("Nomor tidak valid.")
		return -1
	}
	return nomor - 1 // Adjust to 0-based index
}

// Prosedur untuk menampilkan daftar pengeluaran
func tampilkanDaftarPengeluaran(daftar DaftarPengeluaran) {
	fmt.Println("Daftar Pengeluaran:")
	if daftar.N == 0 {
		fmt.Println("Tidak ada pengeluaran.")
		return
	}
	for i := 0; i < daftar.N; i++ {
		fmt.Printf("%d. Kategori: %s, Jumlah: Rp %.2f\n", i+1, daftar.Data[i].Kategori, daftar.Data[i].Jumlah)
	}
}

// Prosedur untuk menambahkan pengeluaran baru
func tambahPengeluaran(daftar *DaftarPengeluaran) {
	if daftar.N >= MAX_PENGELUARAN {
		fmt.Println("Kapasitas pengeluaran penuh!")
		return
	}
	kategori := bacaKategori("Masukkan kategori (Transportasi/Akomodasi/Makanan/Hiburan): ")
	jumlah := bacaJumlah()
	daftar.Data[daftar.N] = Pengeluaran{Kategori: kategori, Jumlah: jumlah}
	daftar.N++
	isSortedByCategory = 0
	fmt.Println("Pengeluaran berhasil ditambahkan!")
}

// Prosedur untuk mengedit pengeluaran berdasarkan nomor
func editPengeluaran(daftar *DaftarPengeluaran) {
	if daftar.N == 0 {
		fmt.Println("Belum ada pengeluaran untuk diedit.")
		return
	}
	tampilkanDaftarPengeluaran(*daftar)
	nomor := bacaNomorPengeluaran(daftar.N)
	if nomor == -1 {
		return
	}
	kategori := bacaKategori("Masukkan kategori baru (Transportasi/Akomodasi/Makanan/Hiburan): ")
	jumlah := bacaJumlah()
	daftar.Data[nomor] = Pengeluaran{Kategori: kategori, Jumlah: jumlah}
	isSortedByCategory = 0
	fmt.Println("Pengeluaran berhasil diedit!")
}

// Prosedur untuk menghapus pengeluaran berdasarkan nomor
func hapusPengeluaran(daftar *DaftarPengeluaran) {
	if daftar.N == 0 {
		fmt.Println("Belum ada pengeluaran untuk dihapus.")
		return
	}
	tampilkanDaftarPengeluaran(*daftar)
	nomor := bacaNomorPengeluaran(daftar.N)
	if nomor == -1 {
		return
	}
	for i := nomor; i < daftar.N-1; i++ {
		daftar.Data[i] = daftar.Data[i+1]
	}
	daftar.N--
	isSortedByCategory = 0
	fmt.Println("Pengeluaran berhasil dihapus!")
}

// Fungsi untuk case apabila memilih menu pengeluaran
func updatePengeluaran(daftar *DaftarPengeluaran) {
	for {
		menuPengeluaran()
		pilihan := bacaUpdate()

		if pilihan == 1 {
			tambahPengeluaran(daftar)
		} else if pilihan == 2 {
			editPengeluaran(daftar)
		} else if pilihan == 3 {
			hapusPengeluaran(daftar)
		} else if pilihan == 4 {
			return
		} else {
			fmt.Println("Pilihan tidak valid, silakan pilih lagi.")
		}
	}
}

// Prosedur untuk mencari pengeluaran berdasarkan kategori
func cariPengeluaran(daftar DaftarPengeluaran) {
	if daftar.N == 0 {
		fmt.Println("Belum ada pengeluaran untuk dicari.")
		return
	}
	kategori := bacaKategori("Masukkan kategori yang ingin dicari (Transportasi/Akomodasi/Makanan/Hiburan): ")
	fmt.Println("Hasil Pencarian:")
	ditemukan := false
	var indeks []int

	if isSortedByCategory == 1 || isSortedByCategory == 2 {
		fmt.Println("Menggunakan Binary Search (data telah diurutkan)")
		indeksPertama := binarySearch(daftar, kategori, isSortedByCategory == 1)
		if indeksPertama != -1 {
			for i := indeksPertama; i < daftar.N && daftar.Data[i].Kategori == kategori; i++ {
				indeks = append(indeks, i)
			}
		}
	} else {
		fmt.Println("Menggunakan Sequential Search (data belum diurutkan)")
		indeks = sequentialSearch(daftar, kategori)
	}

	for _, i := range indeks {
		fmt.Printf("%d. Kategori: %s, Jumlah: Rp %.2f\n", i+1, daftar.Data[i].Kategori, daftar.Data[i].Jumlah)
		ditemukan = true
	}
	if !ditemukan {
		fmt.Println("Tidak ada pengeluaran untuk kategori tersebut.")
	}
}

// Fungsi Sequential Search untuk mencari semua entri dengan kategori tertentu
func sequentialSearch(daftar DaftarPengeluaran, kategori string) []int {
	var indeks []int
	for i := 0; i < daftar.N; i++ {
		if daftar.Data[i].Kategori == kategori {
			indeks = append(indeks, i)
		}
	}
	return indeks
}

// Fungsi Binary Search untuk mencari indeks pertama dengan kategori tertentu
func binarySearch(daftar DaftarPengeluaran, kategori string, aToZ bool) int {
	kiri := 0               // Mulai dari indeks paling kiri
	kanan := daftar.N - 1   // Mulai dari indeks paling kanan

	for kiri <= kanan {     // Lanjutkan selama kiri tidak melewati kanan
		tengah := (kiri + kanan) / 2  // Cari tengah daftar

		// Periksa apakah kategori di tengah cocok
		if daftar.Data[tengah].Kategori == kategori {
			// Cari ke kiri untuk menemukan indeks pertama
			for tengah > 0 && daftar.Data[tengah-1].Kategori == kategori {
				tengah--
			}
			return tengah
		}

		// Tentukan arah pencarian berdasarkan urutan (A-Z atau Z-A)
		if (aToZ && daftar.Data[tengah].Kategori < kategori) || (!aToZ && daftar.Data[tengah].Kategori > kategori) {
			kiri = tengah + 1  // Pindah ke kanan
		} else {
			kanan = tengah - 1 // Pindah ke kiri
		}
	}

	return -1  // Kembali -1 jika tidak ditemukan
}

// Prosedur untuk mencari pengeluaran ekstrem (terbesar dan terkecil)
func cariPengeluaranEkstrem(daftar DaftarPengeluaran) {
	if daftar.N == 0 {
		fmt.Println("Belum ada pengeluaran untuk dicari.")
		return
	}
	var max, min Pengeluaran
	var idxMax, idxMin int
	max.Jumlah = -1
	min.Jumlah = 1e18 // Nilai awal besar untuk minimum

	for i := 0; i < daftar.N; i++ {
		if daftar.Data[i].Jumlah > max.Jumlah {
			max = daftar.Data[i]
			idxMax = i
		}
		if daftar.Data[i].Jumlah < min.Jumlah {
			min = daftar.Data[i]
			idxMin = i
		}
	}

	fmt.Println("\n=== Pengeluaran Ekstrem ===")
	fmt.Printf("Pengeluaran Terbesar: %d. Kategori: %s, Jumlah: Rp %.2f\n", idxMax+1, max.Kategori, max.Jumlah)
	fmt.Printf("Pengeluaran Terkecil: %d. Kategori: %s, Jumlah: Rp %.2f\n", idxMin+1, min.Kategori, min.Jumlah)
}

// Prosedur Bubble Sort untuk mengurutkan daftar pengeluaran
func urutkanPengeluaran(daftar *DaftarPengeluaran, kriteria int) {
	for i := 0; i < daftar.N-1; i++ {
		for j := 0; j < daftar.N-i-1; j++ {
			tukar := false
			switch kriteria {
			case 1: // Jumlah: Besar ke Kecil
				if daftar.Data[j].Jumlah < daftar.Data[j+1].Jumlah {
					tukar = true
				}
			case 2: // Jumlah: Kecil ke Besar
				if daftar.Data[j].Jumlah > daftar.Data[j+1].Jumlah {
					tukar = true
				}
			case 3: // Kategori: A-Z
				if daftar.Data[j].Kategori > daftar.Data[j+1].Kategori {
					tukar = true
				}
			case 4: // Kategori: Z-A
				if daftar.Data[j].Kategori < daftar.Data[j+1].Kategori {
					tukar = true
				}
			}
			if tukar {
				daftar.Data[j], daftar.Data[j+1] = daftar.Data[j+1], daftar.Data[j]
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
func kelolaPengurutan(daftar *DaftarPengeluaran) {
	if daftar.N == 0 {
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
	urutkanPengeluaran(daftar, pilihanUrut)
	fmt.Println("Pengeluaran berhasil diurutkan!")
	tampilkanDaftarPengeluaran(*daftar)
}

// Prosedur untuk menampilkan laporan pengeluaran
func tampilkanLaporan(daftar DaftarPengeluaran, anggaran float64) {
	if daftar.N == 0 {
		fmt.Println("Belum ada pengeluaran untuk dilaporkan.")
		return
	}
	fmt.Println("\n=== Laporan Pengeluaran ===")
	var totalTransportasi, totalAkomodasi, totalMakanan, totalHiburan float64
	var totalPengeluaran float64
	for i := 0; i < daftar.N; i++ {
		if daftar.Data[i].Kategori == "Transportasi" {
			totalTransportasi += daftar.Data[i].Jumlah
		} else if daftar.Data[i].Kategori == "Akomodasi" {
			totalAkomodasi += daftar.Data[i].Jumlah
		} else if daftar.Data[i].Kategori == "Makanan" {
			totalMakanan += daftar.Data[i].Jumlah
		} else if daftar.Data[i].Kategori == "Hiburan" {
			totalHiburan += daftar.Data[i].Jumlah
		}
		totalPengeluaran += daftar.Data[i].Jumlah
	}
	fmt.Println("Rincian per Kategori:")
	if totalTransportasi > 0 {
		fmt.Printf("- Transportasi: Rp %.2f\n", totalTransportasi)
	}
	if totalAkomodasi > 0 {
		fmt.Printf("- Akomodasi: Rp %.2f\n", totalAkomodasi)
	}
	if totalMakanan > 0 {
		fmt.Printf("- Makanan: Rp %.2f\n", totalMakanan)
	}
	if totalHiburan > 0 {
		fmt.Printf("- Hiburan: Rp %.2f\n", totalHiburan)
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
		var kategoriTerbesar string
		var jumlahTerbesar float64
		if totalTransportasi > jumlahTerbesar {
			jumlahTerbesar = totalTransportasi
			kategoriTerbesar = "Transportasi"
		}
		if totalAkomodasi > jumlahTerbesar {
			jumlahTerbesar = totalAkomodasi
			kategoriTerbesar = "Akomodasi"
		}
		if totalMakanan > jumlahTerbesar {
			jumlahTerbesar = totalMakanan
			kategoriTerbesar = "Makanan"
		}
		if totalHiburan > jumlahTerbesar {
			jumlahTerbesar = totalHiburan
			kategoriTerbesar = "Hiburan"
		}
		fmt.Printf("- Kurangi pengeluaran pada %s (Rp %.2f).\n", kategoriTerbesar, jumlahTerbesar)
	} else {
		fmt.Println("- Pertahankan pola pengeluaran Anda, masih ada sisa anggaran.")
	}
}

func main() {
	if !login() {
		return
	}

	var daftar DaftarPengeluaran
	var anggaran float64
	var inputValid bool
	for inputValid == false {
		fmt.Print("Masukkan anggaran awal (Rp): ")
		// Ambil input anggaran dari pengguna
		fmt.Scan(&anggaran)
		// Periksa apakah anggaran adalah angka positif atau nol
		if anggaran >= 0 {
			inputValid = true // Input valid, keluar dari loop
		} else {
			fmt.Println("Anggaran tidak valid. Masukkan nilai positif.")
			var dummy string
			fmt.Scan(&dummy) // Bersihkan input yang salah
			inputValid = false
		}
	}

	for {
		tampilkanMenu()
		pilihan := bacaPilihan()

		if pilihan == 1 {
			updatePengeluaran(&daftar)
		} else if pilihan == 2 {
			cariPengeluaran(daftar)
		} else if pilihan == 3 {
			kelolaPengurutan(&daftar)
		} else if pilihan == 4 {
			tampilkanLaporan(daftar, anggaran)
		} else if pilihan == 5 {
			cariPengeluaranEkstrem(daftar)
		} else if pilihan == 6 {
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		} else {
			fmt.Println("Pilihan tidak valid, silakan pilih lagi.")
		}
	}
}