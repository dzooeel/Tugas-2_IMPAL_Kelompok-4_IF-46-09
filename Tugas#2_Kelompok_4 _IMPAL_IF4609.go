package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func HandleUserInput(input string) string {
	switch input {
	case "*858#":
		return "Selamat datang di layanan *858#.\n1. Transfer Pulsa\n2. Minta Pulsa\n3. Kembali\n"
	case "1":
		return TransferPulsa()
	case "2":
		return MintaPulsa()
	case "3":
		return "Terima kasih telah menggunakan layanan kami. Sampai jumpa!"
	default:
		return "Pilihan tidak valid. Silakan coba lagi."
	}
}

func main() {
	var input string

	fmt.Println("Silakan masukkan kode USSD (contoh: *858#):")
	fmt.Scanln(&input)
	fmt.Println(HandleUserInput(input))

	for {
		fmt.Println("Masukkan pilihan Anda:")
		fmt.Scanln(&input)

		response := HandleUserInput(input)
		fmt.Println(response)

		if input == "3" {
			break
		}
	}
}

func TransferPulsa() string {
	var nomorTujuan, konfirmasi, jumlahTransferStr string

	fmt.Println("")
	fmt.Println("Silahklan masukkan nomor tujuan Tranfer Pulsa :")
	fmt.Println("Contoh (08xxxx atau 628xxxx)")
	fmt.Scanln(&nomorTujuan)
	fmt.Println("")

	if !(strings.HasPrefix(nomorTujuan, "08") || strings.HasPrefix(nomorTujuan, "628")) {
		fmt.Println("Nomor yang Anda masukkan salah. Kembali ke menu awal.")
		return HandleUserInput("*858#")
	}

	fmt.Println("Masukkan jumlah pulsa yang akan ditransfer :")
	fmt.Println("(min 5000 dan max 1jt)")
	fmt.Scanln(&jumlahTransferStr)
	fmt.Println("")

	jumlahTransfer, err := strconv.Atoi(jumlahTransferStr)
	if err != nil {
		fmt.Println("Jumlah pulsa tidak valid.")
		os.Exit(0)
	}

	if jumlahTransfer < 5000 || jumlahTransfer > 1000000 {
		fmt.Println("Jumlah pulsa harus antara 5000 hingga 1000000.")
		os.Exit(0)
	}

	fmt.Println("Anda akan Tranfers Pulsa ", jumlahTransfer, " ke ", nomorTujuan, " ?")
	fmt.Println("1. Ya\n9. Back")
	fmt.Scanln(&konfirmasi)
	fmt.Println("")
	if konfirmasi != "1" {
		return HandleUserInput("*858#")
	}

	fmt.Printf("Transfer pulsa sebesar %d ke nomor %s berhasil.\n", jumlahTransfer, nomorTujuan)

	os.Exit(0)
	return ""
}

func MintaPulsa() string {
	var nomorTujuan, konfirmasi, jumlahPermintaanStr string

	fmt.Println("")
	fmt.Println("Masukkan nomor yang diminta untuk mengirim pulsa:")
	fmt.Println("Contoh (08xxxx atau 628xxxx)")
	fmt.Scanln(&nomorTujuan)
	fmt.Println("")

	if !(strings.HasPrefix(nomorTujuan, "08") || strings.HasPrefix(nomorTujuan, "628")) {
		fmt.Println("Nomor yang Anda masukkan salah. Kembali ke menu awal.")
		return HandleUserInput("*858#")
	}

	fmt.Println("Masukkan jumlah pulsa yang diminta :")
	fmt.Println("(min 5000 dan max 1jt)")
	fmt.Scanln(&jumlahPermintaanStr)
	fmt.Println("")

	jumlahPermintaan, err := strconv.Atoi(jumlahPermintaanStr)
	if err != nil {
		fmt.Println("Jumlah pulsa tidak valid.")
		os.Exit(0)
	}

	if jumlahPermintaan < 5000 || jumlahPermintaan > 1000000 {
		fmt.Println("Jumlah pulsa harus antara 5000 hingga 1000000.")
		os.Exit(0)
	}

	fmt.Println("Anda akan Tranfers Pulsa ", jumlahPermintaan, " ke ", nomorTujuan, " ?")
	fmt.Println("1. Ya\n9. Back")
	fmt.Scanln(&konfirmasi)
	fmt.Println("")
	if konfirmasi != "1" {
		return HandleUserInput("*858#")
	}

	fmt.Printf("Permintaan pulsa sebesar %d ke nomor %s berhasil dikirim.\n", jumlahPermintaan, nomorTujuan)

	os.Exit(0)
	return ""
}
