package p2p

// Message: holds any data that is sent over each transport
// between two nodes

type Message struct {
	Payload []byte
}
