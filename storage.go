package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const databaseFolder = "./database"

// STORE DATA TO FILE
func storeDataPasien(P arrPasien, n int) {
	// Create the database folder if it doesn't exist
	if _, err := os.Stat(databaseFolder); os.IsNotExist(err) {
		err := os.Mkdir(databaseFolder, 0755)
		if err != nil {
			fmt.Println("STORE PASIEN: Error creating database folder:", err)
			return
		}
	}

	filePath := filepath.Join(databaseFolder, "dataPasien.txt")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("STORE PASIEN: Error:", err)
		return
	}
	defer file.Close()

	for i := 0; i < n; i++ {
		fmt.Fprintf(file, "%d_%s_%d_%s\n", P[i].id, P[i].nama, P[i].umur, P[i].alamat)
	}
}

// LOAD DATA FROM FILE
func loadDataPasien(P *arrPasien, n *int) {
	file, err := os.Open("./database/dataPasien.txt")
	if err != nil {
		fmt.Println("LOAD PASIEN: Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() && i < len(P) {
		line := scanner.Text()
		fields := strings.Split(line, "_")
		if len(fields) != 4 {
			continue
		}

		id, _ := strconv.Atoi(fields[0])
		nama := fields[1]
		umur, _ := strconv.Atoi(fields[2])
		alamat := fields[3]

		p := pasien{
			id:     id,
			nama:   nama,
			umur:   umur,
			alamat: alamat,
		}

		P[i] = p
		i++
	}
	*n = i
}

// STORE DATA TO FILE
func storeDataPaket(P arrPaketMcu, n int) {
	// Create the database folder if it doesn't exist
	if _, err := os.Stat(databaseFolder); os.IsNotExist(err) {
		err := os.Mkdir(databaseFolder, 0755)
		if err != nil {
			fmt.Println("STORE PAKET: Error creating database folder:", err)
			return
		}
	}

	filePath := filepath.Join(databaseFolder, "dataPaket.txt")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("STORE PAKET: Error:", err)
		return
	}
	defer file.Close()

	for i := 0; i < n; i++ {
		fmt.Fprintf(file, "%d_%s_%d\n", P[i].id, P[i].nama, P[i].harga)
	}
}

// LOAD DATA FROM FILE
func loadDataPaket(P *arrPaketMcu, n *int) {
	file, err := os.Open("./database/dataPaket.txt")
	if err != nil {
		fmt.Println("LOAD PASIEN: Error:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() && i < len(P) {
		line := scanner.Text()
		fields := strings.Split(line, "_")
		if len(fields) != 3 {
			continue
		}

		id, _ := strconv.Atoi(fields[0])
		nama := fields[1]
		harga, _ := strconv.Atoi(fields[2])

		p := paketMcu{
			id:    id,
			nama:  nama,
			harga: harga,
		}

		P[i] = p
		i++
	}
	*n = i
}
