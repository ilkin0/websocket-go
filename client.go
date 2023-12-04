package main

import (
	"golang.org/x/net/websocket"
	"io"
	"log"
)

type Client struct {
	conenction *websocket.Conn
	server     *Server
}

type ClientList map[*Client]bool

func NewClient(conn *websocket.Conn, server *Server) *Client {
	return &Client{
		conenction: conn,
		server:     server,
	}
}

func (c *Client) readMessages() {
	buf := make([]byte, 1024)

	defer func() {
		c.server.removeClient(c)
	}()

	for {
		n, err := c.conenction.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Panic("Read error: ", err)
			continue
		}

		msg := buf[:n]
		// s.broadcast(msg)
		log.Println(string(msg))
	}
}
