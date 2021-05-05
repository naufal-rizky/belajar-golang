package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strconv"

	"github.com/nats-io/nats.go"
)

func main() {

	flag.Parse()
	// Agar Package flag bisa diakses program
	args := flag.Args()

	// Mendeklarasikan variable args

	// Menyambungkan ke Nats
	nc, err := nats.Connect("nats://127.0.0.1:4222")
	if err != nil {
		log.Fatal(err)
	}

	// Mendeklarasikan variable subj dan i
	subj := args[0]

	// Subscriber
	nc.Subscribe(subj, func(msg *nats.Msg) {

		// Memanggil fungsi printMsg di func main
		printMsg(msg)

		// Mereply request
		nc.Publish(msg.Reply, []byte("Sukses Membuat Bangun Datar"))
	})
	// nc.Flush()

	// Mencetak kata Listening on [Nama Subjek]
	fmt.Printf("Listening on [%s]", subj)

	// Menjalankan Goexit
	runtime.Goexit()
}

func printMsg(m *nats.Msg) {
	fmt.Printf("Hasil Bangun Datar Persegi di [%s]: %s \n", m.Subject, m.Data)

	// Mengkonversi m.Data menjadi String
	toString := string(m.Data)

	// Mengkonversi String menjadi Interger
	sisi, _ := strconv.Atoi(toString)

	for i := 0; i < sisi; i++ {
		for j := 0; j < sisi; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}
