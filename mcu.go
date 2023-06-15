package main

import "fmt"

func searchPasienId(pasien arrPasien, n, id int) int {
	var i = 0
	for i < n && pasien[i].id != id {
		i++
	}
	if i != n {
		return i
	} else {
		return -1
	}
}

func addMcu(P *arrPasien, paket arrPaketMcu, history *arrHistoryMcu, n_pasien, n_history *int, n_paket int) {
	var inPasien pasien
	var jenisMcu int = n_paket + 1
	var isRegis string
	var idPasien, historyId int
	var found bool
	var temp arrHistoryMcu = *history
	clearScreen()
	for isRegis != "y" && isRegis != "n" {
		fmt.Print("Apakah pasien sudah terdaftar? (y/n): ")
		fmt.Scanln(&isRegis)
		switch isRegis {
		case "y":
			for !found {
				fmt.Print("Silahkan masukkan ID pasien: ")
				fmt.Scanln(&idPasien)
				idx := searchPasienId(*P, *n_pasien, idPasien)
				if idx != -1 {
					found = true
					inPasien = P[idx]
					fmt.Println("")
					printSelectedPasien(*P, idx)
					toContinue()
				} else {
					fmt.Println("ID pasien tidak ditemukan")
					fmt.Println("")
				}
			}
		case "n":
			fmt.Println("Masukkan data pasien...")
			fmt.Print("Nama: ")
			fmt.Scanln(&inPasien.nama)
			fmt.Print("Umur: ")
			fmt.Scanln(&inPasien.umur)
			fmt.Print("Alamat: ")
			fmt.Scanln(&inPasien.alamat)
			fmt.Println("")
			fmt.Println("Data anda sudah tersimpan, silahkan melakukan medical checkup!")
			fmt.Println("Klik Enter to continue...")
			fmt.Scanln()
			insertPasien(P, n_pasien, inPasien)
			inPasien = P[*n_pasien-1]
		default:
			fmt.Println("Inputan tidak valid, silahkan ulangi")
			continue
		}
	}
	for jenisMcu > n_paket {
		showPaketMcu(paket, n_paket)
		fmt.Println("---------------------------------")
		fmt.Print("Silahkan masukkan pilihan: ")
		fmt.Scanln(&jenisMcu)
		if jenisMcu > n_paket {
			fmt.Println("Inputan tidak valid, silahkan ulangi")
			toContinue()
		}
	}
	fmt.Println("")
	fmt.Println("Medical Chekup anda telah terekam, Terima kasih!")
	toContinue()
	sortHistoryId(&temp, *n_history, "asc")
	if *n_history > 0 {
		historyId = temp[*n_history-1].id + 1
	} else {
		historyId = temp[*n_history].id + 1
	}
	history[*n_history].id = historyId
	history[*n_history].pasien = inPasien
	history[*n_history].paketMcu = paket[jenisMcu-1].nama
	history[*n_history].totalHarga = paket[jenisMcu-1].harga
	history[*n_history].waktu.tanggal = getDate()
	history[*n_history].waktu.bulan = getMonth()
	history[*n_history].waktu.tahun = getYear()
	*n_history++
	storeDataHistory(*history, *n_history)
	loadDataHistory(history, n_history)
}

func addPaket(P *arrPaketMcu, n *int) {
	var inPaket paketMcu
	var temp arrPaketMcu = *P

	for inPaket.nama != "stop" {
		fmt.Scanln(&inPaket.nama, &inPaket.harga)
		if inPaket.nama != "stop" {
			if !isPaketExist(*P, *n, inPaket.harga, inPaket.nama) {
				sortIdPaket(&temp, *n, "asc")
				if *n > 0 {
					inPaket.id = temp[*n-1].id + 1
				} else {
					inPaket.id = temp[*n].id + 1
				}
				P[*n] = inPaket
				temp[*n] = inPaket
				*n++
				storeDataPaket(*P, *n)
				loadDataPaket(P, n)
			}
		}
	}
}

func isPaketExist(pasien arrPaketMcu, n, harga int, nama string) bool {
	var i int
	for i < n && (pasien[i].nama != nama || pasien[i].harga != harga) {
		i++
	}
	if i != n {
		return true
	} else {
		return false
	}
}

func sortIdPaket(paket *arrPaketMcu, n int, sortType string) {
	var i, j, maxIdx int
	if n > 0 {
		for i = 0; i < n-1; i++ {
			maxIdx = i
			for j = i + 1; j < n; j++ {
				if sortType == "desc" {
					if paket[j].id > paket[maxIdx].id {
						maxIdx = j
					}
				} else {
					if paket[j].id < paket[maxIdx].id {
						maxIdx = j
					}
				}
			}
			paket[i], paket[maxIdx] = paket[maxIdx], paket[i]
		}
	}
}

func sortNamaPaket(paket *arrPaketMcu, n int, sortType string) {
	var i, j, maxIdx int
	for i = 0; i < n-1; i++ {
		maxIdx = i
		for j = i + 1; j < n; j++ {
			if sortType == "desc" {
				if paket[j].nama > paket[maxIdx].nama {
					maxIdx = j
				}
			} else {
				if paket[j].nama < paket[maxIdx].nama {
					maxIdx = j
				}
			}
		}
		paket[i], paket[maxIdx] = paket[maxIdx], paket[i]
	}
}

func sortPaket(P arrPaketMcu, n int) {
	clearScreen()
	var sortBy string
	var valid bool

	for !valid {
		fmt.Print("Urutkan data berdasarkan [id] [nama] [harga]: ")
		fmt.Scanln(&sortBy)
		switch sortBy {
		case "id":
			printSortIdPaket(P, n)
			valid = true
		case "nama":
			printPaket(P, n)
			valid = true
		case "harga":
			printSortHargaPaket(P, n)
			valid = true
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func sortHargaPaket(paket *arrPaketMcu, n int, sortType string) {
	var i, j, maxIdx int
	for i = 0; i < n-1; i++ {
		maxIdx = i
		for j = i + 1; j < n; j++ {
			if sortType == "desc" {
				if paket[j].harga > paket[maxIdx].harga {
					maxIdx = j
				}
			} else {
				if paket[j].harga < paket[maxIdx].harga {
					maxIdx = j
				}
			}
		}
		paket[i], paket[maxIdx] = paket[maxIdx], paket[i]
	}
}

func printPaket(paket arrPaketMcu, n int) {
	clearScreen()
	var sort int = 1
	temp := paket
	for sort == 1 || sort == 2 {
		clearScreen()
		fmt.Println("================== PAKET ===================")
		fmt.Printf("| %-2s | %-16s | %-5s \n", "ID", "NAMA", "HARGA")
		for i := 0; i < n; i++ {
			fmt.Printf("| %-2d | %-16s | %-5d \n", temp[i].id, temp[i].nama, temp[i].harga)
		}
		fmt.Println("---------------------------------------------")
		fmt.Println("1. ASC |  2. DESC |  3. EXIT ")
		fmt.Scanln(&sort)
		if sort == 1 {
			sortNamaPaket(&temp, n, "asc")
		} else if sort == 2 {
			sortNamaPaket(&temp, n, "desc")
		} else {
			sort = 0
		}
	}
}

func printSortHargaPaket(paket arrPaketMcu, n int) {
	clearScreen()
	var sort int = 1
	temp := paket
	for sort == 1 || sort == 2 {
		clearScreen()
		fmt.Println("================== PAKET ===================")
		fmt.Printf("| %-2s | %-16s | %-5s \n", "ID", "NAMA", "HARGA")
		for i := 0; i < n; i++ {
			fmt.Printf("| %-2d | %-16s | %-5d \n", temp[i].id, temp[i].nama, temp[i].harga)
		}
		fmt.Println("---------------------------------------------")
		fmt.Println("1. ASC |  2. DESC |  3. EXIT ")
		fmt.Scanln(&sort)
		if sort == 1 {
			sortHargaPaket(&temp, n, "asc")
		} else if sort == 2 {
			sortHargaPaket(&temp, n, "desc")
		} else {
			sort = 0
		}
	}
}

func printSortIdPaket(paket arrPaketMcu, n int) {
	clearScreen()
	var sort int = 1
	temp := paket
	for sort == 1 || sort == 2 {
		clearScreen()
		fmt.Println("================== PAKET ===================")
		fmt.Printf("| %-2s | %-16s | %-5s \n", "ID", "NAMA", "HARGA")
		for i := 0; i < n; i++ {
			fmt.Printf("| %-2d | %-16s | %-5d \n", temp[i].id, temp[i].nama, temp[i].harga)
		}
		fmt.Println("---------------------------------------------")
		fmt.Println("1. ASC |  2. DESC |  3. EXIT ")
		fmt.Scanln(&sort)
		if sort == 1 {
			sortIdPaket(&temp, n, "asc")
		} else if sort == 2 {
			sortIdPaket(&temp, n, "desc")
		} else {
			sort = 0
		}
	}
}

func editDataPaket(paket *arrPaketMcu, n *int) {
	var nama, data string
	var harga, i, id int
	var found, valid bool
	for !found {
		i = 0
		fmt.Println("")
		fmt.Print("Masukkan ID paket yang ingin diubah: ")
		fmt.Scanln(&id)
		for i < *n && paket[i].id != id {
			i++
		}
		if i != *n {
			found = true
			clearScreen()
			printSelectedPaket(*paket, i)
			for !valid {
				fmt.Print("Masukkan jenis data yang ingin diubah [nama] [umur] [alamat]: ")
				fmt.Scanln(&data)
				if data == "nama" || data == "Nama" {
					fmt.Print("Masukkan data yang baru: ")
					fmt.Scanln(&nama)
					paket[i].nama = nama
					valid = true
				} else if data == "harga" || data == "Harga" {
					fmt.Print("Masukkan data yang baru: ")
					fmt.Scanln(&harga)
					paket[i].harga = harga
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
		storeDataPaket(*paket, *n)
		loadDataPaket(paket, n)
		fmt.Println("")
		printSelectedPaket(*paket, i)
		fmt.Println("Data berhasil dirubah, Terimakasih")
		toContinue()
	}
}

func deleteDataPaket(paket *arrPaketMcu, n *int) {
	var id int
	var i int
	var found bool
	var isDelete string

	for !found {
		i = 0
		fmt.Println("")
		fmt.Print("Masukkan ID paket yang ingin dihapus: ")
		fmt.Scanln(&id)
		for i < *n && paket[i].id != id {
			i++
		}
		if i != *n {
			fmt.Println("")
			printSelectedPaket(*paket, i)
			fmt.Print("Apakah anda yakin ingin menghapus data ini? [y/n]: ")
			fmt.Scanln(&isDelete)
			if isDelete == "y" || isDelete == "Y" {
				for i <= *n-2 {
					paket[i] = paket[i+1]
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
		storeDataPaket(*paket, *n)
		loadDataPaket(paket, n)
		fmt.Println("Data berhasil dihapus, Terimakasih")
		toContinue()
	}
}

func cariDataPaket(paket arrPaketMcu, n int) {
	var i int = 0
	var nameX, nama string
	var harga, id int
	var benar bool = false

	fmt.Println("")
	fmt.Print("Cari berdasarkan [id] [nama] [harga]: ")
	fmt.Scanln(&nameX)
	for !benar {
		if nameX == "id" || nameX == "ID" {
			fmt.Print("Silahkan masukkan id yang ingin dicari: ")
			fmt.Scanln(&id)
			for i < n && paket[i].id != id {
				i++
			}
			benar = true
		} else if nameX == "Nama" || nameX == "nama" {
			fmt.Print("Silahkan masukkan nama yang ingin dicari: ")
			fmt.Scanln(&nama)
			for i < n && paket[i].nama != nama {
				i++
			}
			benar = true

		} else if nameX == "Harga" || nameX == "harga" {
			fmt.Print("Silahkan masukkan harga yang ingin dicari: ")
			fmt.Scanln(&harga)
			for i < n && paket[i].harga != harga {
				i++
			}
			benar = true

		} else {
			fmt.Println("")
			fmt.Println("Maaf kolom", nameX, "tidak terdapat dalam data, silahkan cari berdasarkan kolom data yang ada")
			fmt.Print("Cari berdasarkan [id] [nama] [harga]: ")
			fmt.Scanln(&nameX)
		}
	}
	if i != n {
		clearScreen()
		printSelectedPaket(paket, i)
		toContinue()
	} else {
		clearScreen()
		fmt.Println("Data yang dicari tidak ditemukan")
		toContinue()
	}
}
