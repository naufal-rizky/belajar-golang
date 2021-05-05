package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Switch case dengan gaya IF - Else
	rand.Seed(time.Now().UnixNano())
	var nilai = rand.Intn(100)

	switch {
	case nilai == 100:
		fmt.Println(nilai, "Sempurna")
	case (nilai < 100) && (nilai > 75):
		fmt.Println(nilai, "Bagus")
	case (nilai < 75) && (nilai > 60):
		fmt.Println(nilai, "Lumayan")
	default:
		{
			fmt.Println(nilai, "Remedial")
			fmt.Println("Terus Belajar")
		}
	}
}
