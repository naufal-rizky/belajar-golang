package main

import (
	"flag"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	flag.Parse()
	// Mendeklarasikan variable args dengan flag tipe args
	args := flag.Args()

	// Untuk memparsing args

	// Menyambungkan ke Nats
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}

	subj, payload := args[0], []byte(args[1])

	msg, err := nc.Request(subj, payload, 4*time.Second)

	log.Printf("Mengirim [%s] : %s ", subj, payload)
	log.Printf("Menerima [%v] : %s ", msg.Subject, string(msg.Data))
}
