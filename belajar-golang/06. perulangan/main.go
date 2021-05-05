package main

import "fmt"

func main() {
	o := ""
	// Perulangan menggunakan for
	for i := 0; i < 5; i++ {
		fmt.Printf("Angka "+"%d \n", i)
	}

	fmt.Println(o)
	fmt.Println("For tanpa argumen")

	//penggunaan for tanpa argumen
	var j = 1

	for j < 5 {
		fmt.Println("angka", j)
		j++
	}

	fmt.Println(o)
	//penggunaan keyword break dan continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
	fmt.Println(o)

	//PERULANGAN BERSARANG
	fmt.Println("Perulangan Bersarang")
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}

	//PEMANFAATAN LABEL DALAM PERULANGAN
	fmt.Println(o)

	fmt.Println("Pemanfaatan label dalam perulangan")
outerLoop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
