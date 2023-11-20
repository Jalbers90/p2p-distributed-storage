package main

import (
	"log"

	"github.com/Jalbers90/p2p-distributed-storage/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
