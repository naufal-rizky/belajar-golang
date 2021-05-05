package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	name := []string{"Dapel", "Adnan", "Fadhil", "Rafli", "Opaw"}
	var value = rand.Intn(10)
	rand.Shuffle(len(name), func(i, j int) {
		name[i], name[j] = name[j], name[i]
	})
	// value := 0
	// fmt.Print("Input Value : ")
	// fmt.Scan(&value)

	fmt.Println("Nilai ujian " + name[0] + ":", value)
	if value == 10 {
		fmt.Println("Selamat kamu lolos dengan nilai sempurna!")
	} else if value >= 7 {
		fmt.Println("Selamat " + name[0] + " kamu lolos")
	} else {
		fmt.Println("Maaf " + name[0] + " kamu harus belajar lagi!")
	}
}