package main

import "fmt"

/*note

1. pembuatan fungsi
2. mendeklarasikan variable menggunakan manifest typing
3. fungsi len
4. mendeklarasikan variable tanpa tipe data
5. mendeklarasikan variable menggunakan type inference

*/

func main() {
	/*note error = jika sebuah variable yang sudah dideklarasikan
	tidak dipakai maka program golang akan error*/

	var text string //1. deklarasi variable manifest typing

	text = "hello world"
	fmt.Println(text)

	//len = berfungsi untuk menghitung jumlah karakter pada string
	text = "Halo saya Naufal"
	fmt.Println(len(text))

	//variable dapat dideklarasikan tanpa tipe data
	var angka = 7
	fmt.Println(angka)

	//deklarasi variable type inference
	country := "Indonesia"
	fmt.Println(country)

	//variable underscore = keranjang sampah
	_ = "golang itu ...."

	//Deklarasi Multiple Variable
	var (
		name = "naruto"
		nickname = "shippuden"
	)
	fmt.Println(name, nickname)
}
