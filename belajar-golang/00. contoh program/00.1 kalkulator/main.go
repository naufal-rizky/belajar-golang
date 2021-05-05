package main

import (
	"fmt"
)

func main() {
kalkulator:
	number1 := 0
	number2 := 0
	number := 0

	fmt.Println("Pilih opsi \n 1. Perjumlahan \n 2. Pengurangan \n 3. Perkalian \n 4. Pembagian \n 5. Modulus")
	fmt.Print("Opsi Pilihan : ")
	fmt.Scanln(&number)
	fmt.Print("Masukan nilai pertama: ")
	fmt.Scanln(&number1)
	fmt.Print("Masukan nilai kedua: ")
	fmt.Scanln(&number2)

	tambah := number1 + number2
	kurang := number1 - number2
	kali := number1 * number2
	bagi := number1 / number2
	modulus := number1 % number2

	switch number {
	case 1:
		fmt.Println("Hasilnya =", tambah)
	case 2:
		fmt.Println("Hasilnya =", kurang)
	case 3:
		fmt.Println("Hasilnya =", kali)
	case 4:
		fmt.Println("Hasilnya =", bagi)
	case 5:
		fmt.Println("Hasilnya =", modulus)
	default:
		{
			fmt.Println("Error, Silahkan Pilih opsi")
			goto kalkulator
		}
	}

}
