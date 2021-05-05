package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//Seleksi Kondisi Menggunakan SWITCH CASE

	var color = "merah"
	fmt.Print("Apel ini berwarna")

	switch color {
	case "biru":
		fmt.Println(" Biru")
	case "hijau":
		fmt.Println(" Hijau")
	case "kuning":
		fmt.Println(" Kuning")
	case "merah":
		fmt.Println(" Merah")
	default:
		fmt.Println(" Tidak Berwarna")
	}

	// Switch Case dengan banyak kondisi & menggunakan kurawal pada default

	rand.Seed(time.Now().UnixNano())
	var nilai = rand.Intn(10)
	var kata = [...]string{"Perfect", "Good", "Not Bad", "Poor"}

	switch nilai {
	case 10:
		fmt.Println("Nilai kamu adalah = ", nilai, kata[0])
	case 9, 8, 7:
		fmt.Println("Nilai kamu adalah = ", nilai, kata[1])
	case 6, 5:
		fmt.Println("Nilai kamu adalah = ", nilai, kata[2])
	default:
		{
			fmt.Println("Nilai kamu adalah = ", nilai, kata[3])
			fmt.Println("Terus Belajar")
		}
	}


	// Penggunaan Fallthrough dalam switch case
	switch {
	case nilai == 10:
		fmt.Println(kata[0])
	case (nilai < 10) && (nilai > 7):
		fmt.Println(kata[1])
		
		//Fungsinya adalah memaksa proses pengecekkan ke case selanjutnya.
		fallthrough

	case (nilai < 7) && (nilai > 5):
		fmt.Println(kata[2])
	default:
		{
			fmt.Println(kata[3])
			fmt.Println("Hello")
		}
	}
}
