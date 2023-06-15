package main

import "fmt"

func sortHistory(P arrHistoryMcu, n int) {
	clearScreen()
	var sortBy string
	var valid bool

	for !valid {
		fmt.Print("Urutkan data berdasarkan [paket] [waktu]: ")
		fmt.Scanln(&sortBy)
		switch sortBy {
		case "paket":
			printSortHistoryPaket(P, n)
			valid = true
		case "waktu":
			printSortHistoryWaktu(P, n)
			valid = true
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func sortHistoryId(H *arrHistoryMcu, n int, sortType string) {
	var i, j, maxIdx int
	if n > 0 {
		for i = 0; i < n-1; i++ {
			maxIdx = i
			for j = i + 1; j < n; j++ {
				if sortType == "desc" {
					if H[j].id > H[maxIdx].id {
						maxIdx = j
					}
				} else {
					if H[j].id < H[maxIdx].id {
						maxIdx = j
					}
				}
			}
			H[i], H[maxIdx] = H[maxIdx], H[i]
		}
	}
}

func sortHistoryPaket(H *arrHistoryMcu, n int, sortType string) {
	var i, j, maxIdx int
	if n > 0 {
		for i = 0; i < n-1; i++ {
			maxIdx = i
			for j = i + 1; j < n; j++ {
				if sortType == "desc" {
					if H[j].paketMcu > H[maxIdx].paketMcu {
						maxIdx = j
					}
				} else {
					if H[j].paketMcu < H[maxIdx].paketMcu {
						maxIdx = j
					}
				}
			}
			H[i], H[maxIdx] = H[maxIdx], H[i]
		}
	}
}

func sortHistoryWaktu(H *arrHistoryMcu, n int, sortType string) {
	var i, j, maxIdx int
	if n > 0 {
		for i = 0; i < n-1; i++ {
			maxIdx = i
			for j = i + 1; j < n; j++ {
				if sortType == "desc" {
					if H[j].waktu.tanggal >= H[maxIdx].waktu.tanggal && H[j].waktu.bulan >= H[maxIdx].waktu.bulan && H[j].waktu.tahun >= H[maxIdx].waktu.tahun {
						maxIdx = j
					}
				} else {
					if H[j].waktu.tanggal <= H[maxIdx].waktu.tanggal && H[j].waktu.bulan <= H[maxIdx].waktu.bulan && H[j].waktu.tahun <= H[maxIdx].waktu.tahun {
						maxIdx = j
					}
				}
			}
			H[i], H[maxIdx] = H[maxIdx], H[i]
		}
	}
}

func printSortHistoryPaket(H arrHistoryMcu, n int) {
	clearScreen()
	var sort int = 1
	temp := H
	for sort == 1 || sort == 2 {
		clearScreen()
		fmt.Println("================================================== HISTORY MCU ====================================================")
		fmt.Printf("| %-2s | %-12s | %-16s | %-5s | %-16s | %-10s | %-10s | %-12s \n", "ID", "ID PASIEN", "NAMA", "UMUR", "ALAMAT", "PAKET", "TOTAL", "WAKTU")
		for i := 0; i < n; i++ {
			fmt.Printf("| %-2d | %-12d | %-16s | %-5d | %-16s | %-10s | %-10d | %d %s %d %s %d \n", temp[i].id, temp[i].pasien.id, temp[i].pasien.nama, temp[i].pasien.umur, temp[i].pasien.alamat, temp[i].paketMcu, temp[i].totalHarga, temp[i].waktu.tanggal, "/", temp[i].waktu.bulan, "/", temp[i].waktu.tahun)
		}
		fmt.Println("-------------------------------------------------------------------------------------------------------------------")
		fmt.Println("1. ASC |  2. DESC |  3. EXIT ")
		fmt.Scanln(&sort)
		if sort == 1 {
			sortHistoryPaket(&temp, n, "asc")
		} else if sort == 2 {
			sortHistoryPaket(&temp, n, "desc")
		} else {
			sort = 0
		}
	}
}

func printSortHistoryWaktu(H arrHistoryMcu, n int) {
	clearScreen()
	var sort int = 1
	temp := H
	for sort == 1 || sort == 2 {
		clearScreen()
		fmt.Println("================================================== HISTORY MCU ====================================================")
		fmt.Printf("| %-2s | %-12s | %-16s | %-5s | %-16s | %-10s | %-10s | %-12s \n", "ID", "ID PASIEN", "NAMA", "UMUR", "ALAMAT", "PAKET", "TOTAL", "WAKTU")
		for i := 0; i < n; i++ {
			fmt.Printf("| %-2d | %-12d | %-16s | %-5d | %-16s | %-10s | %-10d | %d %s %d %s %d \n", temp[i].id, temp[i].pasien.id, temp[i].pasien.nama, temp[i].pasien.umur, temp[i].pasien.alamat, temp[i].paketMcu, temp[i].totalHarga, temp[i].waktu.tanggal, "/", temp[i].waktu.bulan, "/", temp[i].waktu.tahun)
		}
		fmt.Println("-------------------------------------------------------------------------------------------------------------------")
		fmt.Println("1. ASC |  2. DESC |  3. EXIT ")
		fmt.Scanln(&sort)
		if sort == 1 {
			sortHistoryWaktu(&temp, n, "asc")
		} else if sort == 2 {
			sortHistoryWaktu(&temp, n, "desc")
		} else {
			sort = 0
		}
	}
}

func editDataHistory(H *arrHistoryMcu, n *int, P arrPaketMcu, n_paket int) {
	var nama, alamat, data, harga string
	var umur, i, id, paket int
	var found, valid bool
	var waktu mcuDate
	for !found {
		i = 0
		fmt.Println("")
		fmt.Print("Masukkan ID History yang ingin diubah: ")
		fmt.Scanln(&id)
		for i < *n && H[i].id != id {
			i++
		}
		if i != *n {
			found = true
			clearScreen()
			printSelectedHistory(*H, i)
			for !valid {
				fmt.Print("Masukkan jenis data yang ingin diubah [nama] [umur] [alamat] [paket] [harga] [waktu]: ")
				fmt.Scanln(&data)
				if data == "nama" || data == "Nama" {
					fmt.Print("Masukkan data yang baru: ")
					fmt.Scanln(&nama)
					H[i].pasien.nama = nama
					valid = true
				} else if data == "umur" || data == "Umur" {
					fmt.Print("Masukkan data yang baru: ")
					fmt.Scanln(&umur)
					H[i].pasien.umur = umur
					valid = true
				} else if data == "alamat" || data == "Alamat" {
					fmt.Print("Masukkan data yang baru: ")
					fmt.Scanln(&alamat)
					H[i].pasien.alamat = alamat
					valid = true
				} else if data == "paket" || data == "Paket" {
					showPaketMcu(P, n_paket)
					fmt.Print("Masukkan data yang baru: ")
					fmt.Scanln(&paket)
					H[i].paketMcu = P[paket-1].nama
					valid = true
				} else if data == "harga" || data == "Harga" {
					fmt.Print("Masukkan data yang baru: ")
					fmt.Scanln(&harga)
					H[i].paketMcu = harga
					valid = true
				} else if data == "waktu" || data == "Waktu" {
					fmt.Print("Masukkan data yang baru [tanggal] [bulan] [tahun]: ")
					fmt.Scanln(&waktu.tanggal, &waktu.bulan, &waktu.tahun)
					H[i].waktu = waktu
					valid = true
				} else {
					fmt.Println("Pilihan tidak valid")
				}
			}
		} else {
			fmt.Println("Maaf data yang anda masukkan tidak ditemukan")
		}
	}
	if valid {
		storeDataHistory(*H, *n)
		loadDataHistory(H, n)
		fmt.Println("")
		printSelectedHistory(*H, i)
		fmt.Println("Data berhasil dirubah, Terimakasih")
		toContinue()
	}
}

func printSelectedHistory(H arrHistoryMcu, i int) {
	fmt.Println("================================================== HISTORY MCU ====================================================")
	fmt.Printf("| %-2s | %-12s | %-16s | %-5s | %-16s | %-10s | %-10s | %-12s \n", "ID", "ID PASIEN", "NAMA", "UMUR", "ALAMAT", "PAKET", "TOTAL", "WAKTU")
	fmt.Printf("| %-2d | %-12d | %-16s | %-5d | %-16s | %-10s | %-10d | %d %s %d %s %d \n", H[i].id, H[i].pasien.id, H[i].pasien.nama, H[i].pasien.umur, H[i].pasien.alamat, H[i].paketMcu, H[i].totalHarga, H[i].waktu.tanggal, "/", H[i].waktu.bulan, "/", H[i].waktu.tahun)
	fmt.Println("-------------------------------------------------------------------------------------------------------------------")
}

func deleteHistory(H *arrHistoryMcu, n *int) {
	var id int
	var i int
	var found bool
	var isDelete string

	for !found {
		i = 0
		fmt.Println("")
		fmt.Print("Masukkan ID History yang ingin dihapus: ")
		fmt.Scanln(&id)
		for i < *n && H[i].id != id {
			i++
		}
		if i != *n {
			fmt.Println("")
			printSelectedHistory(*H, i)
			fmt.Print("Apakah anda yakin ingin menghapus data ini? [y/n]: ")
			fmt.Scanln(&isDelete)
			if isDelete == "y" || isDelete == "Y" {
				for i <= *n-2 {
					H[i] = H[i+1]
					i++
				}
				*n--
				found = true
			}
		} else {
			fmt.Println("Maaf data yang anda masukkan tidak ditemukan")
		}
	}
	if isDelete == "y" || isDelete == "Y" {
		storeDataHistory(*H, *n)
		loadDataHistory(H, n)
		fmt.Println("Data berhasil dihapus, Terimakasih")
		toContinue()
	}
}

func cariDataHistory(H arrHistoryMcu, n int) {
	var i, idx int
	var temp [NMAX]int
	var nameX, paket string
	var waktu mcuDate
	var benar bool = false
	idx = 0

	fmt.Println("")
	fmt.Print("Cari berdasarkan [paket] [waktu]: ")
	fmt.Scanln(&nameX)
	for !benar {
		if nameX == "paket" || nameX == "Paket" {
			fmt.Print("Silahkan masukkan Paket yang ingin dicari: ")
			fmt.Scanln(&paket)
			for i < n {
				if H[i].paketMcu == paket {
					temp[idx] = i
					idx++
				}
				i++
			}
			benar = true
		} else if nameX == "waktu" || nameX == "Waktu" {
			fmt.Print("Silahkan masukkan Waktu yang ingin dicari [tanggal] [bulan] [tahun]: ")
			fmt.Scanln(&waktu.tanggal, &waktu.bulan, &waktu.tahun)
			for i < n {
				if H[i].waktu.tanggal == waktu.tanggal && H[i].waktu.bulan == waktu.bulan && H[i].waktu.tahun == waktu.tahun {
					temp[idx] = i
					idx++
				}
				i++
			}
			benar = true
		} else {
			fmt.Println("")
			fmt.Println("Maaf kolom", nameX, "tidak terdapat dalam data, silahkan cari berdasarkan kolom data yang ada")
			fmt.Print("Cari berdasarkan: ")
			fmt.Scanln(&nameX)
		}
	}
	if idx != 0 {
		clearScreen()
		fmt.Println("================================================== HISTORY MCU ====================================================")
		fmt.Printf("| %-2s | %-12s | %-16s | %-5s | %-16s | %-10s | %-10s | %-12s \n", "ID", "ID PASIEN", "NAMA", "UMUR", "ALAMAT", "PAKET", "TOTAL", "WAKTU")
		for i := 0; i < idx; i++ {
			fmt.Printf("| %-2d | %-12d | %-16s | %-5d | %-16s | %-10s | %-10d | %d %s %d %s %d \n", H[temp[i]].id, H[temp[i]].pasien.id, H[temp[i]].pasien.nama, H[temp[i]].pasien.umur, H[temp[i]].pasien.alamat, H[temp[i]].paketMcu, H[temp[i]].totalHarga, H[temp[i]].waktu.tanggal, "/", H[temp[i]].waktu.bulan, "/", H[temp[i]].waktu.tahun)
		}
		fmt.Println("-------------------------------------------------------------------------------------------------------------------")
		toContinue()
	} else {
		clearScreen()
		fmt.Println("Data yang dicari tidak ditemukan")
		toContinue()
	}
}

func hitungPemasukan(H arrHistoryMcu, n int) {
	var pemasukan int
	var awal, akhir mcuDate

	clearScreen()
	fmt.Print("Masukkan Periode Waktu Awal [tanggal] [bulan] [tahun]: ")
	fmt.Scanln(&awal.tanggal, &awal.bulan, &awal.tahun)

	fmt.Print("Masukkan Periode Waktu Akhir [tanggal] [bulan] [tahun]: ")
	fmt.Scanln(&akhir.tanggal, &akhir.bulan, &akhir.tahun)

	for j := 0; j < n; j++ {
		if H[j].waktu.tahun > awal.tahun && H[j].waktu.tahun < akhir.tahun {
			pemasukan += H[j].totalHarga
		} else if H[j].waktu.tahun == awal.tahun {
			if H[j].waktu.bulan > awal.bulan {
				pemasukan += H[j].totalHarga
			} else if H[j].waktu.bulan == awal.bulan {
				if H[j].waktu.tanggal >= awal.tanggal {
					pemasukan += H[j].totalHarga
				}
			}
		} else if H[j].waktu.tahun == akhir.tahun {
			if H[j].waktu.bulan < akhir.bulan {
				pemasukan += H[j].totalHarga
			} else if H[j].waktu.bulan == akhir.bulan {
				if H[j].waktu.tanggal <= akhir.tanggal {
					pemasukan += H[j].totalHarga
				}
			}
		}
	}
	fmt.Println("Pemasukan Periode", awal.tanggal, "/", awal.bulan, "/", awal.tahun, "-", akhir.tanggal, "/", akhir.bulan, "/", akhir.tahun, ":", pemasukan)
	fmt.Println("")
	toContinue()
}
