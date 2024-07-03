package p2p

// Peer represents a node
type Peer interface{}

// Transport is something that handles communication
type Transport interface {
	ListenAndAccept() error
}
