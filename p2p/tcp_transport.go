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

type TCPTransportOptions struct {
	ListenAddr    string
	HandshakeFunc HandshakeFunc
	Decoder       Decoder
}

type TCPTransport struct {
	TCPTransportOptions
	listener net.Listener

	mu    sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(opts TCPTransportOptions) *TCPTransport {
	return &TCPTransport{
		TCPTransportOptions: opts,
	}
}

func (t *TCPTransport) ListenAndAccept() error {
	ln, err := net.Listen("tcp", t.ListenAddr)

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
		fmt.Printf("INFO: new incoming conn %+v\n", conn)

		go t.handleConn(conn)
	}
}

type Temp struct{}

func (t *TCPTransport) handleConn(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := t.HandshakeFunc(peer); err != nil {
		conn.Close()
		fmt.Printf("TCP error: invalid handshake %s\n", err)
		return
	}

	// data read loop
	buf := make([]byte, 2000)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("TCP error: %s\n", err)

		}
		//if err := t.Decoder.Decode(conn, msg); err != nil {
		//	fmt.Printf("TCP error: %s\n", err)
		//	continue
		//}

		fmt.Printf("message: %+v\n", buf[:n])
	}
}
