package main

import "fmt"

func main() {
	//Perulangan Array
	var buah = [...]string{"Apel", "Mangga", "Jeruk", "Pisang"}

	for i := 0; i < len(buah); i++ {
		fmt.Println(i, buah[i])
	}

	//Perulangan Elemen Array Menggunakan Keyword for - range
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}
	//Penggunaan Variabel Underscore _ Dalam for - range
	var namas = [4]string{"apple", "grape", "banana", "melon"}

	for j := range namas {
		fmt.Printf("nama buah : %d\n", j)
	}
}
