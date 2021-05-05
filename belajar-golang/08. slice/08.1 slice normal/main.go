package main

import "fmt"

func main() {
	var nama = []string{"Ally", "Ploy", "Xu", "Rose", "Lord"}
	var subName1 = nama[0:4]
	var subName2 = nama[1:5]
	fmt.Println(nama)
	fmt.Println(subName1)
	fmt.Println(subName2)

	fmt.Println("Sesudah Diubah")
	//Sesudah Diubah
	subName2[3] = "Nano"
	fmt.Println(nama)
	fmt.Println(subName1)
	fmt.Println(subName2)

	fmt.Println("")

	//Fungsi Cap untuk menghitung lebar atau kapasitas maksimum slice.
	fmt.Println("Cap")
	fmt.Println(cap(nama))
	fmt.Println(cap(subName1))
	fmt.Println(cap(subName2))

	fmt.Println("")

	fmt.Println("Append")
	//Append berfungsi menambah elemen pada slice
	var subName3 = append(subName1, "Wataboku")
	nama = append(nama, "Kagura")
	fmt.Println(subName3)
	fmt.Println(subName1)
	fmt.Println(subName2)
	fmt.Println(nama)

}
