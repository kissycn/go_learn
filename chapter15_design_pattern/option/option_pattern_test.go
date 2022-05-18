package option

import (
	"fmt"
	"testing"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConn  int
}

type Option func(*Server)

func WithProtocol(p string) Option {
	return func(s *Server) {
		s.Protocol = p
	}
}

func WithTimeout(t time.Duration) Option {
	return func(s *Server) {
		s.Timeout = t
	}
}

func NewServer(addr string, port int, options ...func(*Server)) *Server {
	server := Server{
		Addr:     addr,
		Port:     port,
		Protocol: "tcp",
		Timeout:  30 * time.Second,
		MaxConn:  1000,
	}

	for _, opt := range options {
		opt(&server)
	}

	return &server
}

func TestOptionPattern(t *testing.T) {
	s1 := NewServer("localhost", 11)
	s2 := NewServer("localhost", 12, WithProtocol("http"), WithTimeout(10*time.Second))

	fmt.Println(s1)
	fmt.Println(s2)
}
