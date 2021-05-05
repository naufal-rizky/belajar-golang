package main

import "fmt"

func main() {
	pembagian(5, 2)  // Memanggil fungsi pembagian dengan argumen 5 dan 2
	pembagian(10, 5) // Memanggil fungsi pembagian dengan argumen 10 dan 5
	pembagian(4, 3)  // Memanggil fungsi pembagian dengan argumen 4 dan 3
	pembagian(5, 0)  // Memanggil fungsi pembagian dengan argumen 5 dan 0

}

// Fungsi pembagian dengan para1 x dan para2 y bertipe data float64
func pembagian(x, y float64) {
	// Pengkondisian jika y sama dengan 0 maka akan mencetak kata yang dibawahnya
	if y == 0 {
		fmt.Printf("Error. %v tidak dapat dibagi oleh %v", x, y)
		// Return berguna untuk menghentikan proses jika y == 0
		return
	}

	// Deklarasi variable divide yang didalamnya ada method pembagian
	var divide = x / y
	fmt.Printf("Hasil dari %v / %v 		= %v\n", x, y, divide)
}
