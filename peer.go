package main

import "net"

type Peer struct {
	conn net.Conn
}

func NewPeer(connection net.Conn) *Peer {
	return &Peer{
		conn: connection,
	}
}

func (p *Peer) readLoop() {
	for {

	}
}
