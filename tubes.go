package main

import "fmt"

type Pakaian struct {
	Nama     string
	Kategori string
	Warna    string
}

type Outfit struct {
	Nama    string
	Pakaian []Pakaian
	Formal  string
}

var daftarPakaian []Pakaian
var daftarOOTD []Outfit

func tambahPakaian(nama, kategori, warna string) {
	daftarPakaian = append(daftarPakaian, Pakaian{nama, kategori, warna})
	fmt.Println("Pakaian berhasil ditambahkan!")
}

func ubahPakaian(index int, nama, kategori, warna string) {
	if index >= 0 && index < len(daftarPakaian) {
		daftarPakaian[index] = Pakaian{nama, kategori, warna}
		fmt.Println("Pakaian berhasil diubah.")
	} else {
		fmt.Println("Index tidak valid.")
	}
}

func hapusPakaian(index int) {
	if index >= 0 && index < len(daftarPakaian) {
		daftarPakaian = append(daftarPakaian[:index], daftarPakaian[index+1:]...)
		fmt.Println("Pakaian berhasil dihapus.")
	} else {
		fmt.Println("Index tidak valid.")
	}
}

func tampilkanPakaian() {
	if len(daftarPakaian) == 0 {
		fmt.Println("Daftar pakaian kosong.")
		return
	}
	for i, p := range daftarPakaian {
		fmt.Printf("%d. %s (%s, warna: %s)\n", i+1, p.Nama, p.Kategori, p.Warna)
	}
}

func selectionSort() {
	n := len(daftarPakaian)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if daftarPakaian[j].Nama < daftarPakaian[minIdx].Nama {
				minIdx = j
			}
		}
		daftarPakaian[i], daftarPakaian[minIdx] = daftarPakaian[minIdx], daftarPakaian[i]
	}
	fmt.Println("Data pakaian telah diurutkan berdasarkan NAMA (Selection Sort).")
}

func insertionSort() {
	for i := 1; i < len(daftarPakaian); i++ {
		key := daftarPakaian[i]
		j := i - 1
		for j >= 0 && daftarPakaian[j].Kategori > key.Kategori {
			daftarPakaian[j+1] = daftarPakaian[j]
			j--
		}
		daftarPakaian[j+1] = key
	}
	fmt.Println("Data pakaian telah diurutkan berdasarkan KATEGORI (Insertion Sort).")
}

func cariPakaianSequential(warna string) {
	fmt.Println("Hasil pencarian pakaian dengan warna:", warna)
	ditemukan := false
	for _, p := range daftarPakaian {
		if p.Warna == warna {
			fmt.Printf("- %s (%s)\n", p.Nama, p.Kategori)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("Tidak ada pakaian dengan warna tersebut.")
	}
}

func binarySearch(nama string) {
	low, high := 0, len(daftarPakaian)-1
	for low <= high {
		mid := (low + high) / 2
		if daftarPakaian[mid].Nama == nama {
			fmt.Printf("Pakaian ditemukan: %s (%s, warna: %s)\n", daftarPakaian[mid].Nama, daftarPakaian[mid].Kategori, daftarPakaian[mid].Warna)
			return
		} else if daftarPakaian[mid].Nama < nama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Pakaian tidak ditemukan.")
}

func buatOOTD(nama string, indeksPakaian []int, formal string) {
	var kombinasi []Pakaian
	for _, idx := range indeksPakaian {
		if idx >= 0 && idx < len(daftarPakaian) {
			kombinasi = append(kombinasi, daftarPakaian[idx])
		}
	}
	daftarOOTD = append(daftarOOTD, Outfit{nama, kombinasi, formal})
	fmt.Println("OOTD berhasil dibuat:", nama)
}

func tampilkanOOTD() {
	if len(daftarOOTD) == 0 {
		fmt.Println("Belum ada OOTD yang dibuat.")
		return
	}
	for i, o := range daftarOOTD {
		fmt.Printf("%d. %s (%s)\n", i+1, o.Nama, o.Formal)
		for _, p := range o.Pakaian {
			fmt.Printf("   - %s (%s, warna: %s)\n", p.Nama, p.Kategori, p.Warna)
		}
	}
}

func rekomendasi(acara, cuaca string) string {
	switch acara {
	case "formal":
		switch cuaca {
		case "hujan":
			return "Setelan jas dengan jas hujan dan sepatu tahan air."
		case "cerah":
			return "Kemeja dan celana bahan dengan sepatu kulit."
		case "berawan":
			return "Batik dengan payung untuk berjaga-jaga."
		}
	case "casual":
		switch cuaca {
		case "hujan":
			return "Hoodie, celana jeans, dan sepatu sneakers tahan air."
		case "cerah":
			return "Kaos, celana pendek, dan sandal."
		case "berawan":
			return "Kemeja casual dengan celana chino."
		}
	case "sporty":
		switch cuaca {
		case "cerah":
			return "Kaos olahraga dan celana training dengan sepatu lari."
		case "hujan":
			return "Jaket waterproof dan celana training."
		case "berawan":
			return "Kaos lengan panjang dan celana olahraga pendek."
		}
	}
	return "OOTD tidak tersedia untuk kombinasi ini."
}

func main() {
	for {
		fmt.Println("\n========== MENU OOTD PLANNER ==========")
		fmt.Println("1. Tampilkan daftar pakaian")
		fmt.Println("2. Tambah pakaian")
		fmt.Println("3. Ubah pakaian")
		fmt.Println("4. Hapus pakaian")
		fmt.Println("5. Pengurutan Selection Sort (Nama)")
		fmt.Println("6. Pengurutan Insertion Sort (Kategori)")
		fmt.Println("7. Pencarian Sequential (Warna)")
		fmt.Println("8. Pencarian Binary (Nama)")
		fmt.Println("9. Buat OOTD")
		fmt.Println("10. Tampilkan daftar OOTD")
		fmt.Println("11. Rekomendasi OOTD berdasarkan acara/cuaca")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilkanPakaian()
		case 2:
			var nama, kategori, warna string
			fmt.Print("Nama pakaian: ")
			fmt.Scan(&nama)
			fmt.Print("Kategori (casual/formal/sporty): ")
			fmt.Scan(&kategori)
			fmt.Print("Warna: ")
			fmt.Scan(&warna)
			tambahPakaian(nama, kategori, warna)
		case 3:
			tampilkanPakaian()
			if len(daftarPakaian) == 0 {
				break
			}
			var idx int
			var namaBaru, kategoriBaru, warnaBaru string
			fmt.Print("Nomor pakaian yang diubah: ")
			fmt.Scan(&idx)
			idx--
			fmt.Print("Nama baru: ")
			fmt.Scan(&namaBaru)
			fmt.Print("Kategori baru: ")
			fmt.Scan(&kategoriBaru)
			fmt.Print("Warna baru: ")
			fmt.Scan(&warnaBaru)
			ubahPakaian(idx, namaBaru, kategoriBaru, warnaBaru)
		case 4:
			tampilkanPakaian()
			if len(daftarPakaian) == 0 {
				break
			}
			var idx int
			fmt.Print("Nomor pakaian yang dihapus: ")
			fmt.Scan(&idx)
			hapusPakaian(idx - 1)
		case 5:
			selectionSort()
		case 6:
			insertionSort()
		case 7:
			var warna string
			fmt.Print("Masukkan warna yang dicari: ")
			fmt.Scan(&warna)
			cariPakaianSequential(warna)
		case 8:
			if len(daftarPakaian) == 0 {
				fmt.Println("Daftar pakaian kosong. Tambah atau urutkan data dahulu.")
				break
			}
			var nama string
			fmt.Print("Masukkan nama pakaian yang dicari: ")
			fmt.Scan(&nama)
			selectionSort()
			binarySearch(nama)
		case 9:
			tampilkanPakaian()
			if len(daftarPakaian) == 0 {
				fmt.Println("Tambahkan pakaian terlebih dahulu sebelum membuat OOTD.")
				break
			}
			var namaOOTD, formal string
			var jumlah int
			fmt.Print("Nama OOTD: ")
			fmt.Scan(&namaOOTD)
			fmt.Print("Tingkat formalitas (casual/formal): ")
			fmt.Scan(&formal)
			fmt.Print("Berapa banyak item pakaian? ")
			fmt.Scan(&jumlah)
			indeks := make([]int, 0, jumlah)
			for i := 0; i < jumlah; i++ {
				var idx int
				fmt.Printf("Nomor pakaian ke-%d: ", i+1)
				fmt.Scan(&idx)
				indeks = append(indeks, idx-1)
			}
			buatOOTD(namaOOTD, indeks, formal)
		case 10:
			tampilkanOOTD()
		case 11:
			var acara, cuaca string
			fmt.Print("Masukkan jenis acara (formal/casual/sporty): ")
			fmt.Scan(&acara)
			fmt.Print("Masukkan kondisi cuaca (cerah/hujan/berawan): ")
			fmt.Scan(&cuaca)
			hasil := rekomendasi(acara, cuaca)
			fmt.Println("Rekomendasi OOTD:", hasil)
		case 0:
			fmt.Println("Terima kasih telah menggunakan OOTD Planner. Bye!")
			return
		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
