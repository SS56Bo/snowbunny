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
	return &Server{
		Config: cfg,
	}
}

func main() {
	fmt.Println("Snowbunny - SMTP built in Golang")
}
