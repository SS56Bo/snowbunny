package main

import (
	"fmt"
	"log"
	"log/slog"
	"net"
)

type Config struct {
	ListenAddress string // Port or something else
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
		Config: cfg,
	}
}

func (s *Server) StartServer() error {
	listen, err := net.Listen("tcp", s.ListenAddress)
	if err != nil {
		log.Printf("Error: %v", err)
	}
	s.listen = listen
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

}

func main() {
	fmt.Println("Snowbunny - In-memory Database")
}
