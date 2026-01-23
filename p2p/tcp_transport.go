package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents remote node/peer over TCP
type TCPPeer struct {
	// underlying conn
	conn net.Conn

	// dial and retrieve conn => outbound == true
	// accept and retrieve conn => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddress string
	listener      net.Listener
	mu            sync.RWMutex
	peers         map[net.Addr]Peer
}

func NewTCPTransport(listenAddr string) *TCPTransport {
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.listenAddress)

	if err != nil {
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
			fmt.Printf("TCP ERROR: accept => %s\n", err)
		}

		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	fmt.Printf("INFO: new incoming conn %+v\n", peer)
}
