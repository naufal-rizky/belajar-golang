package main

import "fmt"

func main() {
	var buah = map[string]int{
		"mangga": 10,
		"apel":   20,
		"nanas":  30,
	}

	fmt.Println(buah["mangga"])
	makes()
}

//Function Make pada Map
func makes() {
	var ayam = make(map[string]int) //method make un
	ayam["januari"] = 1
	ayam["februari"] = 2
	ayam["maret"] = 3

	fmt.Println(ayam["maret"])

	forRange()
}

//Perulangan Map
func forRange() {
	var hari = map[string]int{
		"Senin":  1,
		"Selasa": 2,
		"Rabu":   3,
		"Kamis":  4,
		"Jumat":  5,
	}
	fmt.Println(len(hari))

	for key, val := range hari {
		fmt.Println(key, "\t:", val)
	}
	fmt.Println(hari)
	delete(hari, "Selasa") //Method untuk menghapus key hari "Selasa"
	fmt.Println(len(hari))
	fmt.Println(hari)

	findKey()
}

//Deteksi Keberadaan Item Dengan Key Tertentu
func findKey() {
	var club = map[string]string{
		"Real Madrid":       "La Liga",
		"Barcelona":         "La Liga",
		"Chelsea":           "EPL",
		"Liverpool":         "EPL",
		"Manchester United": "EPL",
		"Bayern Munchen":    "Bundesliga",
	}

	var value, isExist = club["Real Madrid"]

	if isExist {
		fmt.Println("Hasil", value)
	} else {
		fmt.Println("Error")
	}

	sliceMap()
}


//KOMBINASI SLICE DAN MAP
func sliceMap() {
	var baskets = []map[string]string{
		{"nama": "Lebron James", "club": "LA Lakers"},
		{"nama": "James Harden", "club": "Brooklyn Nets"},
		{"nama": "Stephen Curry", "club": "Golden States Warrior"},
		{"nama": "Zion Williamson", "club": "New Orleans Pelicans"},
	}

	for _, basket := range baskets {
		fmt.Println(basket["nama"], basket["club"])
	}
}
