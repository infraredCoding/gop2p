package p2p

// Peer : interface that represents a remote node
type Peer interface {
}

// Transport : that handles comm between nodes in a network (protocol)
type Transport interface {
	ListenAndAccept() error
}
