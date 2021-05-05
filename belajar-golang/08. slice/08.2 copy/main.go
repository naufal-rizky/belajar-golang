package main

import "fmt"

func main() {
	//Copy digunakan untuk mencopy element slice pada parameter2 ke parameter1
	//Contoh pertama
	para1 := make([]string, 3)
	para2 := []string{"Oke", "Good", "Great", "Excellent"}
	duplikat := copy(para1, para2)

	fmt.Println(para1)
	fmt.Println(para2)
	fmt.Println(duplikat)

	fmt.Println("///////////////////////////////////////")
	fmt.Println("///////////////////////////////////////")

	//Contoh kedua
	angka1 := []int{1,2,3,4}
	angka2 := []int{5,6,7}
	duplikat2 := copy(angka1, angka2)
	

	fmt.Println(angka1)
	fmt.Println(angka2)
	fmt.Println(duplikat2)

	//Penerapan akses elemen dengan 3 indeks
	angka3 := angka1[0:2:3]
	fmt.Println(cap(angka3))
	fmt.Println(len(angka3))
}

