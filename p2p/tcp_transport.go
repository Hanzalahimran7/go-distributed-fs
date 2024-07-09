package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct {
	conn net.Conn
	// if we dial to a remote node, outbound = true
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransportOps struct {
	ListenAddr    string
	HandShakeFunc HandshakeFunc
	Decoder       Decoder
}
type TCPTransport struct {
	TCPTransportOps
	listener net.Listener

	mu    sync.Mutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOps) *TCPTransport {
	return &TCPTransport{
		TCPTransportOps: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.ListenAddr)
	if err != nil {
		fmt.Printf("Error when listening on %s\n", t.ListenAddr)
		return err
	}
	t.listener = ln

	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept Error %s\n", err)
		}
		fmt.Printf("New incoming connection %s\n", conn)
		go t.handleConn(conn)
	}
}

type Temp struct {
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)
	fmt.Printf("New incoming conn %v\n", peer.conn.RemoteAddr())

	if err := t.HandShakeFunc(conn); err != nil {
		fmt.Printf("TCP Handshake Error %s\n", err)
		conn.Close()
		return
	}
	msg := &Temp{}
	for {
		if err := t.Decoder.Decoder(conn, msg); err != nil {
			fmt.Printf("TCP Error %s\n", err)
			continue
		}
	}

}
