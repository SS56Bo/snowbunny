package main

import (
	"fmt"
	"log"
	"net"
)

type Config struct {
	ListenAddress string // Port or something else
}
type Server struct {
	Config
	listen net.Listener
}

func errorLogger(err error) {
	if err != nil {
		log.Printf("Error: %v", err)
	}
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
	errorLogger(err)

	s.listen = listen
	s.acceptLoop()
	return nil
}

func (s *Server) acceptLoop() {
	conn, err := s.listen.Accept()

	s.handleConn(conn)
	errorLogger(err)
}

func (s *Server) handleConn(conn net.Conn) {

}

func main() {
	fmt.Println("Snowbunny - In-memory Database")
}
