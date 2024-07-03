package main

import (
	"log"

	"github.com/hanzalahimran7/go-distributed-fs/p2p"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal("Destruction")
	}
	select {}
}
