package main

import "fmt"

/*
const nilainya tetap / tidak bisa berubah
*/

func main()  {
	const firstname = "Naufal"
	const lastname = "Rizky"

	//Deklarasi Multiple Constant
	const (
		nama = "Opaw"
		Kelas = "XI-RPL"
	)

	//Error
	// firstname = "Kelas"
	fmt.Println(firstname, lastname)
	fmt.Println(nama, Kelas)
}