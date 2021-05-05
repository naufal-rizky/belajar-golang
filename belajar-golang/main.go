package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Memanggil Fungsi Bangun Datar
	// bangunDatar()
	Rndm()
}

// func bangunDatar() {
// 	var i, j, sisi float64

// 	fmt.Print("Masukan Nilai : ")
// 	fmt.Scanln(&sisi)

// 	for i = 0; i < sisi; i++ {
// 		for j = 0; j < sisi; j++ {
// 			fmt.Print("* ")
// 		}
// 		fmt.Println()
// 	}
// }

func Rndm() {
	rand.Seed(time.Now().UnixNano())

	randomValue := randomWithRange(1, 10)
	fmt.Println("Random Number : ", randomValue)
}

func randomWithRange(min, max int) int {
	var value = rand.Int() % (max - min)
	return value
}