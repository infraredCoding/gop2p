package p2p

// HandshakeFunc is the function that does, well, handshaking if necessary
type HandshakeFunc func(Peer) error

func NopHandshakeFunc(Peer) error {
	return nil
}
