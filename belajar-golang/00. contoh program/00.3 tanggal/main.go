package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Kapan adnan akan menikah dengan isa?")
	today := time.Now().Weekday()

	switch time.Wednesday {
	case today + 0:
		fmt.Println("sekarang")
	case today + 1:
		fmt.Println("besok.")
	case today + 2:
		fmt.Println("dalam waktu dekat")
	default:
		fmt.Println("masih lama")
	}
}
