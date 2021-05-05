package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// berfungsi untuk mengclear terminal
	var output, _ = exec.Command("clear").Output()
	fmt.Print(string(output))
main:
	var option int
	var akses string

	fmt.Println("===========================")
	fmt.Println("==   Aplikasi Sederhana  ==")
	fmt.Println("===========================")
	fmt.Println("== 1. Cetak Bangun Datar ==")
	fmt.Println("== 2. Kalkulator         ==")
	fmt.Println("== 3. Info               ==")
	fmt.Println("== 4. Rumus              ==")
	fmt.Println("== 5. Keluar             ==")
	fmt.Println("===========================")

	fmt.Print("Silahkan Pilih Opsi yang Tersedia : ")
	fmt.Scanln(&option)

	switch option {
	case 1:
		if option == 1 {
			fmt.Println()
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			bangunDatar()
		}
	case 2:
		if option == 2 {
			fmt.Println()
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			kalkulator()
		}
	case 3:
		if option == 3 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			info()
			fmt.Print("Kembali ke menu atau keluar (y/n) : ")
			fmt.Scanln(&akses)
			if akses == "y" || akses == "Y" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				main()
			} else if akses == "n" || akses == "N" {
				var output1, _ = exec.Command("clear").Output()
				fmt.Print(string(output1))
				var output, _ = exec.Command("exit").Output()
				fmt.Print(string(output))
			}
		}
	case 4:
		if option == 4 {
			fmt.Println()
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			rumus()
		}
	case 5:
		if option == 5 {
			fmt.Println()
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			var output2, _ = exec.Command("exit").Output()
			fmt.Print(string(output2))
			fmt.Println("Berhasil Keluar dari Aplikasi")
		}
	default:
		if option > 5 || option <= 0 {
			fmt.Println()
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			fmt.Println("Opsi tidak tersedia, silahkan pilih opsi yang ada")
			goto main
		}
	}
}

func bangunDatar() {
bangunDatar:
	var option int
	var akses string

	fmt.Println("===========================")
	fmt.Println("=== Cetak Bangun Datar  ===")
	fmt.Println("===========================")
	fmt.Println("=== 1. Persegi          ===")
	fmt.Println("=== 2. Persegi Panjang  ===")
	fmt.Println("=== 3. Segitiga         ===")
	fmt.Println("=== 4. Kotak            ===")
	fmt.Println("=== 5. Kembali ke Menu  ===")
	fmt.Println("===========================")

	fmt.Print("Silahkan Pilih Opsi yang Tersedia : ")
	fmt.Scanln(&option)

	switch option {
	case 1:
	persegi:
		if option == 1 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			persegi()
			fmt.Print("Keluar dari [Persegi]? (y/n) : ")
			fmt.Scanln(&akses)
			if akses == "y" || akses == "Y" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				bangunDatar()
			} else if akses == "n" || akses == "N" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				goto persegi
			} else {
				os.Exit(1)
			}
		}
	case 2:
	persegiPanjang:
		if option == 2 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			persegiPanjang()
			fmt.Print("Keluar dari [Persegi Panjang]? (y/n) : ")
			fmt.Scanln(&akses)
			if akses == "y" || akses == "Y" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				bangunDatar()
			} else if akses == "n" || akses == "N" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				goto persegiPanjang
			} else {
				os.Exit(1)
			}
		}
	case 3:
	segitiga:
		if option == 3 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			segitiga()
			fmt.Print("Keluar dari [Segitiga]? (y/n) : ")
			fmt.Scanln(&akses)
			if akses == "y" || akses == "Y" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				bangunDatar()
			} else if akses == "n" || akses == "N" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				goto segitiga
			} else {
				os.Exit(1)
			}
		}
	case 4:
	kotak:
		if option == 4 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			kotak()
			fmt.Print("Keluar dari [Kotak]? (y/n) : ")
			fmt.Scanln(&akses)
			if akses == "y" || akses == "Y" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				bangunDatar()
			} else if akses == "n" || akses == "N" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				goto kotak
			} else {
				os.Exit(1)
			}
		}
	case 5:
		main()
	default:
		fmt.Println("Opsi tidak tersedia, silahkan pilih opsi yang ada !!")
		var output, _ = exec.Command("clear").Output()
		fmt.Print(string(output))
		goto bangunDatar
	}
}

func kalkulator() {
kalkulator:
	var option int
	var akses string

	fmt.Println("===========================")
	fmt.Println("===      Kalkulator     ===")
	fmt.Println("===========================")
	fmt.Println("=== 1. Perjumlahan      ===")
	fmt.Println("=== 2. Pengurangan      ===")
	fmt.Println("=== 3. Perkalian        ===")
	fmt.Println("=== 4. Pembagian        ===")
	fmt.Println("=== 5. Modulus          ===")
	fmt.Println("=== 6. Kembali ke Menu  ===")
	fmt.Println("===========================")

	fmt.Print("Silahkan Pilih Opsi yang Tersedia : ")
	fmt.Scanln(&option)

	switch option {
	case 1:
	tambah:
		if option == 1 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			tambah()
			fmt.Print("Keluar dari [Perjumlahan]? (y/n) : ")
			fmt.Scanln(&akses)
			if akses == "y" || akses == "Y" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				kalkulator()
			} else if akses == "n" || akses == "N" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				goto tambah
			} else {
				os.Exit(1)
			}
		}
	case 2:
	kurang:
		if option == 2 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			kurang()
			fmt.Print("Keluar dari [Pengurangan]? (y/n) : ")
			fmt.Scanln(&akses)
			if akses == "y" || akses == "Y" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				kalkulator()
			} else if akses == "n" || akses == "N" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				goto kurang
			} else {
				os.Exit(1)
			}
		}
	case 3:
	kali:
		if option == 3 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			kali()
			fmt.Print("Keluar dari [Perkalian]]? (y/n) : ")
			fmt.Scanln(&akses)
			if akses == "y" || akses == "Y" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				kalkulator()
			} else if akses == "n" || akses == "N" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				goto kali
			} else {
				os.Exit(1)
			}
		}
	case 4:
	bagi:
		if option == 4 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			bagi()
			fmt.Print("Keluar dari [Pembagian]? (y/n) : ")
			fmt.Scanln(&akses)
			if akses == "y" || akses == "Y" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				kalkulator()
			} else if akses == "n" || akses == "N" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				goto bagi
			} else {
				os.Exit(1)
			}
		}
	case 5:
	modulus:
		if option == 5 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			modulus()
			fmt.Print("Keluar dari [Modulus]? (y/n) : ")
			fmt.Scanln(&akses)
			if akses == "y" || akses == "Y" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				kalkulator()
			} else if akses == "n" || akses == "N" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				goto modulus
			} else {
				os.Exit(1)
			}
		}
	case 6:
		main()
	default:
		fmt.Println("Opsi tidak tersedia, silahkan pilih opsi yang ada")
		var output, _ = exec.Command("clear").Output()
		fmt.Print(string(output))
		goto kalkulator
	}
}

func persegi() {
	var i, j, sisi int

	fmt.Print("Masukan Nilai dari [Persegi] = ")
	fmt.Scanln(&sisi)

	fmt.Println("[ Hasil Bangun Datar ]")

	for i = 0; i < sisi; i++ {
		for j = 0; j < sisi; j++ {
			fmt.Print("= ")
		}
		fmt.Println()
	}
}

func persegiPanjang() {
	var i, j, panjang, lebar int

	fmt.Print("Masukan Nilai dari Panjang [Persegi Panjang] = ")
	fmt.Scanln(&panjang)
	fmt.Print("Masukan Nilai dari Lebar [Persegi Panjang] = ")
	fmt.Scanln(&lebar)

	fmt.Println("[ Hasil Bangun Datar ]")

	for i = 0; i < lebar; i++ {
		for j = 0; j < panjang; j++ {
			fmt.Print("= ")
		}
		fmt.Println()
	}
}

func segitiga() {
	var i, j, k, nilai int

	fmt.Print("Masukan Nilai dari [Segitiga] = ")
	fmt.Scanln(&nilai)

	fmt.Println("[ Hasil Bangun Datar ]")

	for i = 0; i < nilai; i++ {
		for j = nilai; j > i; j-- {
			fmt.Print(" ")
		}
		for k = 0; k < i; k++ {
			fmt.Print("= ")
		}
		fmt.Println()
	}
	fmt.Println()
}

func kotak() {
	var i, j, nilai int

	fmt.Print("Masukan Nilai dari [Kotak] = ")
	fmt.Scanln(&nilai)

	fmt.Println("[ Hasil Bangun Datar ]")

	for i = 0; i < nilai; i++ {
		for j = 0; j < nilai; j++ {
			if i == 0 || i == nilai-1 || j == 0 || j == nilai-1 {
				fmt.Print("= ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}

func tambah() {
	var nilai1, nilai2 int

	fmt.Print("Masukan Nilai Pertama untuk [Perjumlahan] = ")
	fmt.Scanln(&nilai1)
	fmt.Print("Masukan Nilai Kedua untuk [Perjumlahan] = ")
	fmt.Scanln(&nilai2)

	tambah := nilai1 + nilai2
	fmt.Println("Hasil [Perjumlahan] = ", tambah)
}

func kurang() {
	var nilai1, nilai2 int

	fmt.Print("Masukan Nilai Pertama untuk [Pengurangan] = ")
	fmt.Scanln(&nilai1)
	fmt.Print("Masukan Nilai Kedua untuk [Pengurangan] = ")
	fmt.Scanln(&nilai2)

	kurang := nilai1 - nilai2
	fmt.Println("Hasil [Pengurangan] = ", kurang)
}

func kali() {
	var nilai1, nilai2 int

	fmt.Print("Masukan Nilai Pertama untuk [Perkalian] = ")
	fmt.Scanln(&nilai1)
	fmt.Print("Masukan Nilai Kedua untuk [Perkalian] = ")
	fmt.Scanln(&nilai2)

	kali := nilai1 * nilai2
	fmt.Println("Hasil [Perkalian] = ", kali)
}

func bagi() {
	var nilai1, nilai2 float64

	fmt.Print("Masukan Nilai Pertama untuk [Pembagian] = ")
	fmt.Scanln(&nilai1)
	fmt.Print("Masukan Nilai Kedua untuk [Pembagian] = ")
	fmt.Scanln(&nilai2)

	bagi := nilai1 / nilai2
	fmt.Println("Hasil [Pembagian] = ", bagi)
}

func modulus() {
	var nilai1, nilai2 int

	fmt.Print("Masukan Nilai Pertama untuk [Modulus] = ")
	fmt.Scanln(&nilai1)
	fmt.Print("Masukan Nilai Kedua untuk [Modulus] = ")
	fmt.Scanln(&nilai2)

	modulus := nilai1 % nilai2
	fmt.Println("Hasil [Modulus] = ", modulus)
}

func info() {
	fmt.Println("==========================================")
	fmt.Println("=== ================================== ===")
	fmt.Println("=== Aplikasi Sederhana by Naufal Rizky ===")
	fmt.Println("=== ================================== ===")
	fmt.Println("==========================================")
}

func rumus() {
	var opsi int
	// var akses string

	fmt.Println("===========================")
	fmt.Println("===        Rumus        ===")
	fmt.Println("===========================")
	fmt.Println("=== 1. Persegi          ===")
	fmt.Println("=== 2. Keluar ke menu   ===")
	fmt.Println("===========================")

	fmt.Print("Masukan Pilihan Anda = ")
	fmt.Scanln(&opsi)

	switch opsi {
	case 1:
		if opsi == 1 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			rumusPersegi()
		}
	case 2:

	}
}

func rumusPersegi() {
	var opsi int
	var akses string

	fmt.Println("===========================")
	fmt.Println("===    Rumus Persegi    ===")
	fmt.Println("===========================")
	fmt.Println("===========================")
	fmt.Println("=== 1. Luas Persegi     ===")
	fmt.Println("=== 2. Keliling Persegi ===")
	fmt.Println("=== 3. Keluar ke menu   ===")
	fmt.Println("===========================")

	fmt.Print("Pilih Opsi Rumus [Persegi] = ")
	fmt.Scanln(&opsi)

	switch opsi {
	case 1:
	luaspsg:
		if opsi == 1 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			luasPsg()
			fmt.Print("Apakah ingin keluar dari [Luas Persegi] (y/n) : ")
			fmt.Scanln(&akses)
			if akses == "y" || akses == "Y" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				rumusPersegi()
			} else if akses == "n" || akses == "N" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				goto luaspsg
			} else {
				os.Exit(1)
			}
		}
	case 2:
	klngpsg:
		if opsi == 2 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			klngPsg()
			fmt.Print("Apakah ingin keluar dari [Keliling Persegi] (y/n) : ")
			fmt.Scanln(&akses)
			if akses == "y" || akses == "Y" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				rumusPersegi()
			} else if akses == "n" || akses == "N" {
				var output, _ = exec.Command("clear").Output()
				fmt.Print(string(output))
				goto klngpsg
			} else {
				os.Exit(1)
			}
		}
	case 3:
		if opsi == 3 {
			var output, _ = exec.Command("clear").Output()
			fmt.Print(string(output))
			rumus()
		}
	}

}

func luasPsg() {
	var nilai int

	fmt.Print("Masukan Nilai untuk [Sisi Persegi] = ")
	fmt.Scanln(&nilai)

	luas := nilai * nilai
	fmt.Print("Hasil Luas [Persegi] = ", luas)
}

func klngPsg() {
	var nilai int

	fmt.Print("Masukan Nilai untuk [Sisi Persegi] = ")
	fmt.Scanln(&nilai)

	keliling := nilai * 4
	fmt.Print("Hasil keliling [Persegi] = ", keliling)
}
