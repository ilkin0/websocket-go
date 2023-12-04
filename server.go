package main

import (
	"fmt"
	"sync"

	"golang.org/x/net/websocket"
)

type Server struct {
	connections map[*websocket.Conn]bool
	clients     ClientList
	sync.RWMutex
	bufferSize int
}

func NewServer() *Server {
	return &Server{
		connections: make(map[*websocket.Conn]bool),
		clients:     make(ClientList),
	}
}

func (s *Server) serveWS(ws *websocket.Conn) {
	fmt.Println("New incoming connection from client:", ws.RemoteAddr())

	s.connections[ws] = true

	client := NewClient(ws, s)
	s.addClient(client)

	go client.readMessages()
}

func (s *Server) addClient(c *Client) {
	s.Lock()
	defer s.Unlock()

	s.clients[c] = true
}

func (s *Server) removeClient(c *Client) {
	s.Lock()
	defer s.Unlock()

	if _, ok := s.clients[c]; ok {
		c.conenction.Close()
		delete(s.clients, c)
	}
}
