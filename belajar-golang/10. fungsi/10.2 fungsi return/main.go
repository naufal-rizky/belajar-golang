package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	// Merandom angka
	rand.Seed(time.Now().Unix())
	// Mendeklarasikan variable nilaiRandom bertipe data interger
	var nilaiRandom int

	// variable yang memanggil funsi
	nilaiRandom = randomWithRange(2, 20)
	fmt.Println("angka random: ", nilaiRandom)
}

// Fungsi randowWithRange dengan parameter 1 = max dan parameter 2 = min bertipe data int
// Dengan return value bertipe data interger
func randomWithRange(max, min int) int  {
	// Mendeklarasikan variable nilai
	var nilai = rand.Int() % (max - min + 1) + min
	return nilai
}