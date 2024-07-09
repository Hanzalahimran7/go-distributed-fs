package main

import (
	"log"

	"github.com/hanzalahimran7/go-distributed-fs/p2p"
)

func main() {
	opts := p2p.TCPTransportOps{
		ListenAddr:    ":3000",
		HandShakeFunc: p2p.NoHandShakeFunc,
		Decoder:       p2p.GOBDecoder{},
	}
	tr := p2p.NewTCPTransport(opts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal("Destruction")
	}
	select {}
}
