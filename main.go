package main

import (
	"fmt"
	"net"
)

type Config struct {
	ListenAddress string
}
type Server struct {
	Config
	listen net.Listener
}

func CreateNewServer(cfg Config) *Server {
	if (len(cfg.ListenAddress)) == 0 {
		cfg.ListenAddress = ":7777"
	}
	return &Server{
		Config: cfg,
	}
}

func (s *Server) StartServer() error {
	listen, err := net.Listen("tcp", s.ListenAddress)
	if err != nil {
		return err
	}
}

func main() {
	fmt.Println("Snowbunny - SMTP built in Golang")
}
