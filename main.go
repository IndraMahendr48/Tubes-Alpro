package main

import "fmt"

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
<<<<<<< HEAD
	fmt.Println("1. Tambah Pengeluaran")
	fmt.Println("2. Edit Pengeluaran")
	fmt.Println("3. Hapus Pengeluaran")
	fmt.Println("4. Cari Pengeluaran")
	fmt.Println("5. Urutkan Pengeluaran")
	fmt.Println("6. Lihat Laporan")
	fmt.Println("7. Keluar")
	fmt.Println("8. Lihat Nilai Ekstrim (Tertinggi & Terendah)")

=======
	fmt.Println("1. Update Pengeluaran")
	fmt.Println("2. Cari Pengeluaran")
	fmt.Println("3. Urutkan Pengeluaran")
	fmt.Println("4. Lihat Laporan")
	fmt.Println("5. Cari Pengeluaran Ekstrem")
	fmt.Println("6. Keluar")
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
}

// Fungsi untuk membaca pilihan menu dari pengguna
func bacaPilihan() int {
	fmt.Print("Pilih menu (1-7): ")
	var pilihan int
<<<<<<< HEAD
=======
	fmt.Print("Pilih menu (1-6): ")
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
	fmt.Scan(&pilihan)
	return pilihan
}

<<<<<<< HEAD
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
=======
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
	validKategori := map[string]bool{
		"Transportasi": true,
		"Akomodasi":   true,
		"Makanan":     true,
		"Hiburan":     true,
	}
	return validKategori[kategori]
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
}

// Fungsi untuk membaca kategori dari pengguna
func bacaKategori(prompt string) string {
<<<<<<< HEAD
	fmt.Print(prompt)
	var kategori string
	fmt.Scan(&kategori)
	return kategori
=======
	var kategori string
	for {
		fmt.Print(prompt)
		fmt.Scan(&kategori)
		if isValidKategori(kategori) {
			return kategori
		}
		fmt.Println("Kategori tidak valid. Gunakan: Transportasi, Akomodasi, Makanan, atau Hiburan.")
	}
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
}

// Fungsi untuk membaca jumlah pengeluaran dari pengguna
func bacaJumlah() float64 {
	fmt.Print("Masukkan jumlah (Rp): ")
	var jumlah float64
<<<<<<< HEAD
	fmt.Scan(&jumlah)
	return jumlah
=======
	for {
		fmt.Print("Masukkan jumlah (Rp): ")
		if _, err := fmt.Scan(&jumlah); err == nil && jumlah >= 0 {
			return jumlah
		}
		fmt.Println("Jumlah tidak valid. Masukkan nilai positif.")
		var dummy string
		fmt.Scan(&dummy) // Clear invalid input
	}
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
}

// Fungsi untuk membaca nomor pengeluaran, mengembalikan -1 jika tidak valid
func bacaNomorPengeluaran(n int) int {
	var nomor int
	fmt.Print("Masukkan nomor pengeluaran: ")
	if _, err := fmt.Scan(&nomor); err != nil || nomor < 1 || nomor > n {
		fmt.Println("Nomor tidak valid.")
		return -1
	}
	return nomor - 1 // Adjust to 0-based index
}

// Prosedur untuk menampilkan daftar pengeluaran
func tampilkanDaftarPengeluaran(daftar DaftarPengeluaran) {
	fmt.Println("Daftar Pengeluaran:")
<<<<<<< HEAD
	for i, pengeluaran := range daftarPengeluaran {
		fmt.Printf("%d. Kategori: %s, Jumlah: Rp %.2f\n", i+1, pengeluaran.Kategori, pengeluaran.Jumlah)
=======
	if daftar.N == 0 {
		fmt.Println("Tidak ada pengeluaran.")
		return
	}
	for i := 0; i < daftar.N; i++ {
		fmt.Printf("%d. Kategori: %s, Jumlah: Rp %.2f\n", i+1, daftar.Data[i].Kategori, daftar.Data[i].Jumlah)
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
	}
}

// Prosedur untuk menambahkan pengeluaran baru
<<<<<<< HEAD
func tambahPengeluaran(daftarPengeluaran *[]Pengeluaran) {
	// Minta kategori dan jumlah
	kategori := bacaKategori("Masukkan kategori (Transportasi/Akomodasi/Makanan/Hiburan): ")
	jumlah := bacaJumlah()
	// Tambahkan data ke daftar pengeluaran
	*daftarPengeluaran = append(*daftarPengeluaran, Pengeluaran{Kategori: kategori, Jumlah: jumlah})
	// Tandai bahwa data tidak lagi diurutkan
=======
func tambahPengeluaran(daftar *DaftarPengeluaran) {
	if daftar.N >= MAX_PENGELUARAN {
		fmt.Println("Kapasitas pengeluaran penuh!")
		return
	}
	kategori := bacaKategori("Masukkan kategori (Transportasi/Akomodasi/Makanan/Hiburan): ")
	jumlah := bacaJumlah()
	daftar.Data[daftar.N] = Pengeluaran{Kategori: kategori, Jumlah: jumlah}
	daftar.N++
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
	isSortedByCategory = 0
	fmt.Println("Pengeluaran berhasil ditambahkan!")
}

// Prosedur untuk mengedit pengeluaran berdasarkan nomor
func editPengeluaran(daftar *DaftarPengeluaran) {
	if daftar.N == 0 {
		fmt.Println("Belum ada pengeluaran untuk diedit.")
		return
	}
<<<<<<< HEAD
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
=======
	tampilkanDaftarPengeluaran(*daftar)
	nomor := bacaNomorPengeluaran(daftar.N)
	if nomor == -1 {
		return
	}
	kategori := bacaKategori("Masukkan kategori baru (Transportasi/Akomodasi/Makanan/Hiburan): ")
	jumlah := bacaJumlah()
	daftar.Data[nomor] = Pengeluaran{Kategori: kategori, Jumlah: jumlah}
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
	isSortedByCategory = 0
	fmt.Println("Pengeluaran berhasil diedit!")
}

// Prosedur untuk menghapus pengeluaran berdasarkan nomor
func hapusPengeluaran(daftar *DaftarPengeluaran) {
	if daftar.N == 0 {
		fmt.Println("Belum ada pengeluaran untuk dihapus.")
		return
	}
<<<<<<< HEAD
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
=======
	tampilkanDaftarPengeluaran(*daftar)
	nomor := bacaNomorPengeluaran(daftar.N)
	if nomor == -1 {
		return
	}
	for i := nomor; i < daftar.N-1; i++ {
		daftar.Data[i] = daftar.Data[i+1]
	}
	daftar.N--
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
	isSortedByCategory = 0
	fmt.Println("Pengeluaran berhasil dihapus!")
}

// Fungsi untuk case apabila memilih menu pengeluaran
func updatePengeluaran(daftar *DaftarPengeluaran) {
	for {
		menuPengeluaran()
		pilihan := bacaUpdate()

		switch pilihan {
		case 1:
			tambahPengeluaran(daftar)
		case 2:
			editPengeluaran(daftar)
		case 3:
			hapusPengeluaran(daftar)
		case 4:
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan pilih lagi.")
		}
	}
}

// Fungsi Sequential Search untuk mencari semua entri dengan kategori tertentu
func sequentialSearch(daftar DaftarPengeluaran, kategori string) []int {
	var indeks []int
<<<<<<< HEAD
	// Periksa setiap elemen secara berurutan
	for i, pengeluaran := range daftarPengeluaran {
		if pengeluaran.Kategori == kategori {
=======
	for i := 0; i < daftar.N; i++ {
		if daftar.Data[i].Kategori == kategori {
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
			indeks = append(indeks, i)
		}
	}
	return indeks
}

// Fungsi Binary Search untuk mencari indeks pertama dengan kategori tertentu
func binarySearch(daftar DaftarPengeluaran, kategori string, aToZ bool) int {
	kiri, kanan := 0, daftar.N-1
	for kiri <= kanan {
		tengah := (kiri + kanan) / 2
<<<<<<< HEAD
		if (aToZ && daftarPengeluaran[tengah].Kategori == kategori) || (!aToZ && daftarPengeluaran[tengah].Kategori == kategori) {
			// Cari indeks pertama dengan kategori ini
			for tengah > 0 && daftarPengeluaran[tengah-1].Kategori == kategori {
=======
		if daftar.Data[tengah].Kategori == kategori {
			for tengah > 0 && daftar.Data[tengah-1].Kategori == kategori {
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
				tengah--
			}
			return tengah
		}
		if (aToZ && daftar.Data[tengah].Kategori < kategori) || (!aToZ && daftar.Data[tengah].Kategori > kategori) {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

// Prosedur untuk mencari pengeluaran berdasarkan kategori
func cariPengeluaran(daftar DaftarPengeluaran) {
	if daftar.N == 0 {
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
		indeksPertama := binarySearch(daftar, kategori, isSortedByCategory == 1)
		if indeksPertama != -1 {
<<<<<<< HEAD
			// Tambahkan semua entri dengan kategori yang sama
			for i := indeksPertama; i < len(daftarPengeluaran) && daftarPengeluaran[i].Kategori == kategori; i++ {
=======
			for i := indeksPertama; i < daftar.N && daftar.Data[i].Kategori == kategori; i++ {
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
				indeks = append(indeks, i)
			}
		}
	} else {
		// Gunakan Sequential Search jika data belum diurutkan
		fmt.Println("Menggunakan Sequential Search (data belum diurutkan)")
		indeks = sequentialSearch(daftar, kategori)
	}

	// Tampilkan hasil pencarian
	for _, i := range indeks {
		fmt.Printf("%d. Kategori: %s, Jumlah: Rp %.2f\n", i+1, daftar.Data[i].Kategori, daftar.Data[i].Jumlah)
		ditemukan = true
	}
	if !ditemukan {
		fmt.Println("Tidak ada pengeluaran untuk kategori tersebut.")
	}
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
<<<<<<< HEAD
func urutkanPengeluaran(daftarPengeluaran *[]Pengeluaran, kriteria int) {
	n := len(*daftarPengeluaran)
	// Bubble Sort: bandingkan dan tukar elemen berpasangan
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
=======
func urutkanPengeluaran(daftar *DaftarPengeluaran, kriteria int) {
	for i := 0; i < daftar.N-1; i++ {
		for j := 0; j < daftar.N-i-1; j++ {
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
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
			// Tukar elemen jika diperlukan
			if tukar {
				daftar.Data[j], daftar.Data[j+1] = daftar.Data[j+1], daftar.Data[j]
			}
		}
	}
}

// Prosedur untuk mengelola pengurutan pengeluaran
func kelolaPengurutan(daftar *DaftarPengeluaran) {
	if daftar.N == 0 {
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
<<<<<<< HEAD
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
=======
	urutkanPengeluaran(daftar, pilihanUrut)
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
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
	// Inisialisasi map untuk menyimpan total per kategori
	totalPerKategori := map[string]float64{
		"Transportasi": 0,
		"Akomodasi":    0,
		"Makanan":      0,
		"Hiburan":      0,
	}
<<<<<<< HEAD
	totalPengeluaran := 0.0
	// Hitung total pengeluaran per kategori
	for _, pengeluaran := range daftarPengeluaran {
		totalPerKategori[pengeluaran.Kategori] += pengeluaran.Jumlah
		totalPengeluaran += pengeluaran.Jumlah
=======
	var totalPengeluaran float64
	for i := 0; i < daftar.N; i++ {
		totalPerKategori[daftar.Data[i].Kategori] += daftar.Data[i].Jumlah
		totalPengeluaran += daftar.Data[i].Jumlah
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
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

// Fungsi untuk mencari pengeluaran terbesar dan terkecil
func cariNilaiEkstrim(daftarPengeluaran []Pengeluaran) {
	if len(daftarPengeluaran) == 0 {
		fmt.Println("Belum ada data pengeluaran.")
		return
	}

	// Inisialisasi nilai ekstrem
	maxIndex, minIndex := 0, 0
	for i := 1; i < len(daftarPengeluaran); i++ {
		if daftarPengeluaran[i].Jumlah > daftarPengeluaran[maxIndex].Jumlah {
			maxIndex = i
		}
		if daftarPengeluaran[i].Jumlah < daftarPengeluaran[minIndex].Jumlah {
			minIndex = i
		}
	}

	fmt.Println("\n=== Pengeluaran Ekstrem ===")
	fmt.Printf("Pengeluaran Terbesar: %s - Rp %.2f\n", daftarPengeluaran[maxIndex].Kategori, daftarPengeluaran[maxIndex].Jumlah)
	fmt.Printf("Pengeluaran Terkecil: %s - Rp %.2f\n", daftarPengeluaran[minIndex].Kategori, daftarPengeluaran[minIndex].Jumlah)
}

func main() {
	// Meminta pengguna untuk login
	if !login() {
		return
	}

<<<<<<< HEAD
	// Deklarasi variabel untuk menyimpan daftar pengeluaran dan anggaran awal
	var daftarPengeluaran []Pengeluaran
	var anggaran float64

	// Meminta pengguna memasukkan anggaran awal
	fmt.Print("Masukkan anggaran awal (Rp): ")
	fmt.Scan(&anggaran)
=======
	var daftar DaftarPengeluaran
	var anggaran float64
	for {
		fmt.Print("Masukkan anggaran awal (Rp): ")
		if _, err := fmt.Scan(&anggaran); err == nil && anggaran >= 0 {
			break
		}
		fmt.Println("Anggaran tidak valid. Masukkan nilai positif.")
		var dummy string
		fmt.Scan(&dummy) // Clear invalid input
	}
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480

	// Loop utama untuk menampilkan menu dan memproses pilihan
	for {
		tampilkanMenu()
		pilihan := bacaPilihan()
<<<<<<< HEAD
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
			case 8:
				cariNilaiEkstrim(daftarPengeluaran)
				return
			default:
				fmt.Println("Pilihan tidak valid, silakan pilih lagi.")
			}
=======

		switch pilihan {
		case 1:
			updatePengeluaran(&daftar)
		case 2:
			cariPengeluaran(daftar)
		case 3:
			kelolaPengurutan(&daftar)
		case 4:
			tampilkanLaporan(daftar, anggaran)
		case 5:
			cariPengeluaranEkstrem(daftar)
		case 6:
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan pilih lagi.")
>>>>>>> 8fd09563af3853d0db576b5d1612e782ab80f480
		}
	}
}