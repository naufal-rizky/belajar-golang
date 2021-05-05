package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Seleksi kondisi bersarang mengunakkan kombinasi switch case dan if else
	rand.Seed(time.Now().UnixNano())
	var point = rand.Intn(10)

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}

}
