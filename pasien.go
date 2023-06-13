package main

import (
	"fmt"
)

func addPasien(P *arrPasien, n *int) {
	var inPasien pasien
	var temp arrPasien = *P

	for inPasien.nama != "stop" {
		fmt.Scanln(&inPasien.nama, &inPasien.umur, &inPasien.alamat)
		sortId(&temp, *n, "asc")
		if *n > 0 {
			inPasien.id = temp[*n-1].id + 1
		} else {
			inPasien.id = temp[*n].id + 1
		}
		if inPasien.nama != "stop" {
			P[*n] = inPasien
			*n++
			storeDataPasien(*P, *n)
			loadDataPasien(P, n)
		}
	}
}

func sortPasien(P arrPasien, n int) {
	clearScreen()
	var sortBy string
	var valid bool

	for !valid {
		fmt.Print("Urutkan data berdasarkan [id] [nama] [umur] [alamat]: ")
		fmt.Scanln(&sortBy)
		switch sortBy {
		case "id":
			printSortId(P, n)
			valid = true
		case "nama":
			printPasien(P, n)
			valid = true
		case "umur":
			printSortUmur(P, n)
			valid = true
		case "alamat":
			printSortAlamat(P, n)
			valid = true
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func sortId(pasien *arrPasien, n int, sortType string) {
	var i, j, maxIdx int
	if n > 0 {
		for i = 0; i < n-1; i++ {
			maxIdx = i
			for j = i + 1; j < n; j++ {
				if sortType == "desc" {
					if pasien[j].id > pasien[maxIdx].id {
						maxIdx = j
					}
				} else {
					if pasien[j].id < pasien[maxIdx].id {
						maxIdx = j
					}
				}
			}
			pasien[i], pasien[maxIdx] = pasien[maxIdx], pasien[i]
		}
	}
}

func sortUmur(pasien *arrPasien, n int, sortType string) {
	var i, j, maxIdx int
	for i = 0; i < n-1; i++ {
		maxIdx = i
		for j = i + 1; j < n; j++ {
			if sortType == "desc" {
				if pasien[j].umur > pasien[maxIdx].umur {
					maxIdx = j
				}
			} else {
				if pasien[j].umur < pasien[maxIdx].umur {
					maxIdx = j
				}
			}
		}
		pasien[i], pasien[maxIdx] = pasien[maxIdx], pasien[i]
	}
}

func sortAlamat(pasien *arrPasien, n int, sortType string) {
	var i, j, maxIdx int
	for i = 0; i < n-1; i++ {
		maxIdx = i
		for j = i + 1; j < n; j++ {
			if sortType == "desc" {
				if pasien[j].alamat > pasien[maxIdx].alamat {
					maxIdx = j
				}
			} else {
				if pasien[j].alamat < pasien[maxIdx].alamat {
					maxIdx = j
				}
			}
		}
		pasien[i], pasien[maxIdx] = pasien[maxIdx], pasien[i]
	}
}

func sortNama(pasien *arrPasien, n int, sortType string) {
	var i, j, maxIdx int
	for i = 0; i < n-1; i++ {
		maxIdx = i
		for j = i + 1; j < n; j++ {
			if sortType == "desc" {
				if pasien[j].nama > pasien[maxIdx].nama {
					maxIdx = j
				}
			} else {
				if pasien[j].nama < pasien[maxIdx].nama {
					maxIdx = j
				}
			}
		}
		pasien[i], pasien[maxIdx] = pasien[maxIdx], pasien[i]
	}
}

func cariDataPasien(pasien arrPasien, n int) {
	var i int = 0
	var nameX, nama, address string
	var age, id int
	var benar bool = false

	fmt.Println("")
	fmt.Print("Cari berdasarkan [id] [nama] [umur] [alamat]: ")
	fmt.Scanln(&nameX)
	for !benar {
		if nameX == "id" || nameX == "ID" {
			fmt.Print("Silahkan masukkan id yang ingin dicari: ")
			fmt.Scanln(&id)
			for i < n && pasien[i].id != id {
				i++
			}
			benar = true
		} else if nameX == "Nama" || nameX == "nama" {
			fmt.Print("Silahkan masukkan nama yang ingin dicari: ")
			fmt.Scanln(&nama)
			for i < n && pasien[i].nama != nama {
				i++
			}
			benar = true

		} else if nameX == "Alamat" || nameX == "alamat" {
			fmt.Print("Silahkan masukkan data alamat yang ingin dicari: ")
			fmt.Scanln(&address)
			for i < n && pasien[i].alamat != address {
				i++
			}
			benar = true

		} else if nameX == "Umur" || nameX == "umur" {
			fmt.Print("Silahkan masukkan data umur yang ingin dicari: ")
			fmt.Scanln(&age)
			for i < n && pasien[i].umur != age {
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
	if i != n {
		clearScreen()
		printSelectedPasien(pasien, i)
		toContinue()
	} else {
		clearScreen()
		fmt.Println("Data yang dicari tidak ditemukan")
		toContinue()
	}
}

func printPasien(pasien arrPasien, n int) {
	clearScreen()
	var sort int = 1
	temp := pasien
	for sort == 1 || sort == 2 {
		clearScreen()
		fmt.Println("================== PASIEN ===================")
		fmt.Printf("| %-2s | %-16s | %-5s | %-16s \n", "ID", "NAMA", "UMUR", "ALAMAT")
		for i := 0; i < n; i++ {
			fmt.Printf("| %-2d | %-16s | %-5d | %-16s \n", temp[i].id, temp[i].nama, temp[i].umur, temp[i].alamat)
		}
		fmt.Println("---------------------------------------------")
		fmt.Println("1. ASC |  2. DESC |  3. EXIT ")
		fmt.Scanln(&sort)
		if sort == 1 {
			sortNama(&temp, n, "asc")
		} else if sort == 2 {
			sortNama(&temp, n, "desc")
		} else {
			sort = 0
		}
	}
}

func printSortId(pasien arrPasien, n int) {
	clearScreen()
	var sort int = 1
	temp := pasien
	for sort == 1 || sort == 2 {
		clearScreen()
		fmt.Println("================== PASIEN ===================")
		fmt.Printf("| %-2s | %-16s | %-5s | %-16s \n", "ID", "NAMA", "UMUR", "ALAMAT")
		for i := 0; i < n; i++ {
			fmt.Printf("| %-2d | %-16s | %-5d | %-16s \n", temp[i].id, temp[i].nama, temp[i].umur, temp[i].alamat)
		}
		fmt.Println("---------------------------------------------")
		fmt.Println("1. ASC |  2. DESC |  3. EXIT ")
		fmt.Scanln(&sort)
		if sort == 1 {
			sortId(&temp, n, "asc")
		} else if sort == 2 {
			sortId(&temp, n, "desc")
		} else {
			sort = 0
		}
	}
}

func printSortUmur(pasien arrPasien, n int) {
	clearScreen()
	var sort int = 1
	temp := pasien
	for sort == 1 || sort == 2 {
		clearScreen()
		fmt.Println("================== PASIEN ===================")
		fmt.Printf("| %-2s | %-16s | %-5s | %-16s \n", "ID", "NAMA", "UMUR", "ALAMAT")
		for i := 0; i < n; i++ {
			fmt.Printf("| %-2d | %-16s | %-5d | %-16s \n", temp[i].id, temp[i].nama, temp[i].umur, temp[i].alamat)
		}
		fmt.Println("---------------------------------------------")
		fmt.Println("1. ASC |  2. DESC |  3. EXIT ")
		fmt.Scanln(&sort)
		if sort == 1 {
			sortUmur(&temp, n, "asc")
		} else if sort == 2 {
			sortUmur(&temp, n, "desc")
		} else {
			sort = 0
		}
	}
}

func printSortAlamat(pasien arrPasien, n int) {
	clearScreen()
	var sort int = 1
	temp := pasien
	for sort == 1 || sort == 2 {
		clearScreen()
		fmt.Println("================== PASIEN ===================")
		fmt.Printf("| %-2s | %-16s | %-5s | %-16s \n", "ID", "NAMA", "UMUR", "ALAMAT")
		for i := 0; i < n; i++ {
			fmt.Printf("| %-2d | %-16s | %-5d | %-16s \n", temp[i].id, temp[i].nama, temp[i].umur, temp[i].alamat)
		}
		fmt.Println("---------------------------------------------")
		fmt.Println("1. ASC |  2. DESC |  3. EXIT ")
		fmt.Scanln(&sort)
		if sort == 1 {
			sortAlamat(&temp, n, "asc")
		} else if sort == 2 {
			sortAlamat(&temp, n, "desc")
		} else {
			sort = 0
		}
	}
}

func editDataPasien(pasien *arrPasien, n *int) {
	var nama, alamat, data string
	var umur, i, id int
	var found, valid bool
	for !found {
		i = 0
		fmt.Println("")
		fmt.Print("Masukkan ID pasien yang ingin diubah: ")
		fmt.Scanln(&id)
		for i < *n && pasien[i].id != id {
			i++
		}
		if i != *n {
			found = true
			clearScreen()
			printSelectedPasien(*pasien, i)
			for !valid {
				fmt.Print("Masukkan jenis data yang ingin diubah [nama] [umur] [alamat]: ")
				fmt.Scanln(&data)
				if data == "nama" || data == "Nama" {
					fmt.Print("Masukkan data yang baru: ")
					fmt.Scanln(&nama)
					pasien[i].nama = nama
					valid = true
				} else if data == "umur" || data == "Umur" {
					fmt.Print("Masukkan data yang baru: ")
					fmt.Scanln(&umur)
					pasien[i].umur = umur
					valid = true
				} else if data == "alamat" || data == "Alamat" {
					fmt.Print("Masukkan data yang baru: ")
					fmt.Scanln(&alamat)
					pasien[i].alamat = alamat
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
		storeDataPasien(*pasien, *n)
		loadDataPasien(pasien, n)
		fmt.Println("")
		printSelectedPasien(*pasien, i)
		fmt.Println("Data berhasil dirubah, Terimakasih")
		toContinue()
	}
}

func deleteData(pasien *arrPasien, n *int) {
	var id int
	var i int
	var found bool
	var isDelete string

	for !found {
		i = 0
		fmt.Println("")
		fmt.Print("Masukkan ID pasien yang ingin dihapus: ")
		fmt.Scanln(&id)
		for i < *n && pasien[i].id != id {
			i++
		}
		if i != *n {
			fmt.Println("")
			printSelectedPasien(*pasien, i)
			fmt.Print("Apakah anda yakin ingin menghapus data ini? [y/n]: ")
			fmt.Scanln(&isDelete)
			if isDelete == "y" || isDelete == "Y" {
				for i <= *n-2 {
					pasien[i] = pasien[i+1]
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
		storeDataPasien(*pasien, *n)
		loadDataPasien(pasien, n)
		fmt.Println("Data berhasil dihapus, Terimakasih")
		toContinue()
	}
}
