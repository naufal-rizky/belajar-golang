package main

import "fmt"

func main() {
	//Inisialisasi Nilai Awal Array
	var buah = [4]string{"Apel", "Anggur", "Jeruk", "Pepaya"}
	fmt.Println("Isi Array \t \t=", buah)
	fmt.Println("Jumlah Array \t=", len(buah))

	//Inisialisasi Nilai Awal Array Tanpa Jumlah Elemen
	var nama = [...]string{
		"Opaw",
		"Adnan",
		"Fadhil",
		"Rafli",
		"Dapel",
	}
	fmt.Println("Isi Array \t \t=", nama)
	fmt.Println("Jumlah Array \t=", len(nama))

	//Array Multidimensi
	var numbers1 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//Alokasi Elemen Array Menggunakan Keyword make
	var contoh = make([]string, 4)
	contoh[0] = "1"
	contoh[1] = "2"
	contoh[2] = "3"
	contoh[3] = "4"

	fmt.Print(contoh, " = ")
	fmt.Println(len(contoh))
}
