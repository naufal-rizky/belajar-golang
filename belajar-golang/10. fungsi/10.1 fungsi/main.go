package main

import (
	"fmt"
	"strings"
)

func main() {
	// Deklarasi variable nama dengan slice tipe string
	var nama = []string{"Naufal", "Rizky"}

	// Memanggil func printPesan
	printPesan("halo", nama)
}

// Fungsi printPesan dengan 2 parameter yaitu parameter pesan bertipe data string
// dan slice bertipe data string dengan nama = name
func printPesan(pesan string, names []string) {

	
	var nameString = strings.Join(names, " ")
	fmt.Println(pesan, nameString)
}
