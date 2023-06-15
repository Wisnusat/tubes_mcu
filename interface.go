package main

import "fmt"

func menuInterface(menu *int) {
	/* I.S Diberikan parameter in/out menu untuk menampung pilihan menu dari user
	   F.S Mengubah nilai dari parameter menu sesuai inputan user */
	clearScreen()
	fmt.Println("==============================")
	fmt.Println("   Layanan Medical Checkup    ")
	fmt.Println("==============================")
	fmt.Println("Pilihan menu yang tersedia: ")
	fmt.Println("1. Medical Checkup")
	fmt.Println("2. Data Paket MCU")
	fmt.Println("3. Data Pasien")
	fmt.Println("4. Rekap Hasil MCU")
	fmt.Println("5. Exit")
	fmt.Println("------------------------------")
	fmt.Print("Silahkan masukkan pilihan: ")
	fmt.Scanln(menu)
}

func subMenuInterface(menu int, subMenu *int) {
	/* I.S Diberikan Parameter menu untuk menentukan data apa yang akan diolah
	   F.S Menampilkan Pilihan pengolahan data */
	var title string
	if menu == 2 {
		title = "Data Paket MCU"
	} else if menu == 3 {
		title = "Data Pasien"
	} else if menu == 4 {
		title = "Data History MCU"
	}
	clearScreen()
	fmt.Println("==============================")
	fmt.Println("        ", title, "           ")
	fmt.Println("==============================")
	fmt.Println("Pilihan menu yang tersedia: ")
	fmt.Println("1. Lihat Data")
	fmt.Println("2. Edit Data")
	fmt.Println("3. Tambah Data")
	fmt.Println("4. Delete Data")
	fmt.Println("5. Cari Data")
	fmt.Println("6. Sort Data")
	if menu == 4 {
		fmt.Println("7. Laporan Pemasukan")
		fmt.Println("8. Back")
	} else {
		fmt.Println("7. Back")
	}
	fmt.Println(" ------------------------------")
	fmt.Print("Silahkan masukkan pilihan: ")
	fmt.Scanln(subMenu)
}

func showPaketMcu(P arrPaketMcu, n int) {
	clearScreen()
	fmt.Println("==============================")
	fmt.Println("Daftar Layanan Medical Checkup")
	fmt.Println("==============================")
	for i := 0; i < n; i++ {
		fmt.Printf("%-1d | %-16s | %-5d \n", i+1, P[i].nama, P[i].harga)
	}
}

func printSelectedPasien(pasien arrPasien, i int) {
	fmt.Println("================== PASIEN ===================")
	fmt.Printf("| %-2s | %-16s | %-5s | %-16s \n", "ID", "NAMA", "UMUR", "ALAMAT")
	fmt.Printf("| %-2d | %-16s | %-5d | %-16s \n", pasien[i].id, pasien[i].nama, pasien[i].umur, pasien[i].alamat)
	fmt.Println("---------------------------------------------")
}

func printSelectedPaket(paket arrPaketMcu, i int) {
	fmt.Println("================== PASIEN ===================")
	fmt.Printf("| %-2s | %-16s | %-5s \n", "ID", "NAMA", "HARGA")
	fmt.Printf("| %-2d | %-16s | %-5d \n", paket[i].id, paket[i].nama, paket[i].harga)
	fmt.Println("---------------------------------------------")
}
