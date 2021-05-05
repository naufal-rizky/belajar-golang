package main

import "fmt"
	// "math/rand"
	// "time"

func main() {
	// rand.Seed(time.Now().UnixNano())
	// opaw := rand.Intn(100)

	// fmt.Println("Nama Saya Opaw umur saya: ", opaw)

	// fmt.Println("////////////////////")

	//Operator Logika
	var benar = true
	var salah = false

	benarAndSalah := benar && salah //perbandingan AND
	fmt.Println("benar && salah", benarAndSalah)

	benarOrSalah := benar || salah //perbandingan OR
	fmt.Println("benar || salah", benarOrSalah)

	benarReverse := !benar //negasi atau kebalikan
	fmt.Println("!benar", benarReverse)
}
