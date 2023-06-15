package main

import (
	"fmt"
)

const NMAX = 100

type pasien struct {
	id     int
	nama   string
	umur   int
	alamat string
}

type paketMcu struct {
	id    int
	nama  string
	harga int
}

type historyMcu struct {
	id         int
	pasien     pasien
	paketMcu   string
	totalHarga int
	waktu      mcuDate
}

type mcuDate struct {
	tanggal int
	bulan   int
	tahun   int
}

type arrPasien [NMAX]pasien
type arrPaketMcu [NMAX]paketMcu
type arrHistoryMcu [NMAX]historyMcu

func main() {
	var menu, subMenu int
	var paketMcu arrPaketMcu
	var pasien arrPasien
	var historyMcu arrHistoryMcu
	var n_pasien, n_paket, n_history int
	var isSubMenu bool
	loadDataPaket(&paketMcu, &n_paket)
	loadDataPasien(&pasien, &n_pasien)
	loadDataHistory(&historyMcu, &n_history)
	menuInterface(&menu)
	for menu != 5 {
		if menu == 1 {
			addMcu(&pasien, paketMcu, &historyMcu, &n_pasien, &n_history, n_paket)
		} else if menu == 2 || menu == 3 || menu == 4 {
			isSubMenu = true
			for isSubMenu {
				subMenuInterface(menu, &subMenu)
				if menu == 2 {
					switch subMenu {
					case 1:
						printPaket(paketMcu, n_paket)
					case 2:
						editDataPaket(&paketMcu, &n_paket)
					case 3:
						addPaket(&paketMcu, &n_paket)
					case 4:
						deleteDataPaket(&paketMcu, &n_paket)
					case 5:
						cariDataPaket(paketMcu, n_paket)
					case 6:
						sortPaket(paketMcu, n_paket)
					case 7:
						isSubMenu = false
					default:
						fmt.Println("Pilihan tidak valid, silahkan masukan pilihan lagi")
						continue
					}
				} else if menu == 3 {
					switch subMenu {
					case 1:
						printPasien(pasien, n_pasien)
					case 2:
						editDataPasien(&pasien, &n_pasien)
					case 3:
						addPasien(&pasien, &n_pasien)
					case 4:
						deleteData(&pasien, &n_pasien)
					case 5:
						cariDataPasien(pasien, n_pasien)
					case 6:
						sortPasien(pasien, n_pasien)
					case 7:
						isSubMenu = false
					default:
						fmt.Println("Pilihan tidak valid, silahkan masukan pilihan lagi")
						continue
					}
				} else if menu == 4 {
					switch subMenu {
					case 1:
						printSortHistoryWaktu(historyMcu, n_history)
					case 2:
						editDataHistory(&historyMcu, &n_history, paketMcu, n_paket)
					case 3:
						fmt.Println("")
						fmt.Println("Silahkan lakukan Medical Checkup")
						toContinue()
						addMcu(&pasien, paketMcu, &historyMcu, &n_pasien, &n_history, n_paket)
					case 4:
						deleteHistory(&historyMcu, &n_history)
					case 5:
						cariDataHistory(historyMcu, n_history)
					case 6:
						sortHistory(historyMcu, n_history)
					case 7:
						hitungPemasukan(historyMcu, n_history)
					case 8:
						isSubMenu = false
					default:
						fmt.Println("Pilihan tidak valid, silahkan masukan pilihan lagi")
						continue
					}
				}
			}
		} else if menu > 1 || menu < 5 {
			fmt.Println("Inputan Tidak Valid, silahkan ulangi")
		}
		menuInterface(&menu)
	}
}
