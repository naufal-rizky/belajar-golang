package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Seleksi Kondisi IF and ELSE dan ELSE IF

	rand.Seed(time.Now().UnixNano())
	var nilai = rand.Intn(100)

	fmt.Println("Nilai Ujian Anda adalah: ", nilai)

	if nilai == 100 {
		fmt.Println("Selamat nilai anda sempurna!")
	} else if nilai >= 70 {
		fmt.Println("Selamat anda lolos!")
	} else {
		fmt.Println("Maaf anda harus remed!")
	}

	//Variable Temporary Pada IF - ELSE

	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")
	}

}
