package p2p

type HandshakeFunc func(any) error

func NoHandShakeFunc(any) error { return nil }
