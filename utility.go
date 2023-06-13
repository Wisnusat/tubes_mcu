package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func getDate() int {
	/* I.S.
	   F.S. Output memberikan nilai tanggal pada hari ini */

	// Get the current date
	currentTime := time.Now()

	// Extract the year, month, and day
	day := currentTime.Day()

	return day
}

func getMonth() int {
	/* I.S.
	   F.S. Output memberikan nilai bulan pada hari ini */

	// Get the current date
	currentTime := time.Now()

	// Extract the year, month, and day
	month := int(currentTime.Month())
	return month
}

func getYear() int {
	/* I.S.
	   F.S. Output memberikan nilai tahun pada hari ini */

	// Get the current date
	currentTime := time.Now()

	// Extract the year, month, and day
	year := currentTime.Year()

	return year
}

func clearScreen() {
	/* I.S.
	   F.S. Output membersihkan tampilan terminal */

	cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func toContinue() {
	/* I.S.
	   F.S. Tampilan untuk berhenti sebelum lanjut  */
	fmt.Print("Klik Enter to continue...")
	fmt.Scanln()
}
