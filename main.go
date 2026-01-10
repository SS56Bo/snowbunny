package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
)

type Config struct {
	ListenAddress string
}

type Server struct {
	Config
	peer           map[*Peer]bool
	listen         net.Listener
	addPeerChannel chan *Peer
}

func CreateNewServer(cfg Config) *Server {
	if len(cfg.ListenAddress) == 0 {
		cfg.ListenAddress = ":7777"
	}
	return &Server{
		Config:         cfg,
		peer:           make(map[*Peer]bool),
		addPeerChannel: make(chan *Peer),
	}
}

func (s *Server) StartServer() error {
	listen, err := net.Listen("tcp", s.ListenAddress)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	s.listen = listen

	s.loop()
	s.acceptLoop()
	return nil
}

func (s *Server) loop() {
	for {
		select {
		case peer := <-s.addPeerChannel:
			s.peer[peer] = true
		default:
			fmt.Printf("Default")
		}
	}
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listen.Accept()
		if err != nil {
			slog.Error("Accept Error", "Err", err)
			continue
		}

		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	peer := NewPeer(conn)
	s.addPeerChannel <- peer

	go peer.readLoop()
}

func main() {
	server := CreateNewServer(Config{})

	log.Fatal(server.StartServer())
}
