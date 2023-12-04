package main

// import (
// 	"fmt"
// 	"io"
// 	"net/http"

// 	"golang.org/x/net/websocket"
// )

// type Server struct {
// 	conns map[*websocket.Conn]bool
// }

// func NewServer() *Server {
// 	fmt.Println("Starting new WS Server...")
// 	return &Server{
// 		conns: make(map[*websocket.Conn]bool),
// 	}
// }

// func (s *Server) handleWS(ws *websocket.Conn) {
// 	fmt.Println("New incoming connection from client:", ws.RemoteAddr())

// 	s.conns[ws] = true // TODO in golang map is not thread-safe/concurrent. Have to add Mutex implementation

// 	s.readLoop(ws)
// }

// func (s *Server) readLoop(ws *websocket.Conn) {
// 	buf := make([]byte, 1024)

// 	for {
// 		n, err := ws.Read(buf)
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			fmt.Println("Read error: ", err)
// 			continue
// 		}

// 		msg := buf[:n]
// 		s.broadcast(msg)
// 	}
// }

// func (s *Server) broadcast(data []byte) {
// 	for ws := range s.conns {
// 		go func(c *websocket.Conn) {
// 			if _, err := ws.Write(data); err != nil {
// 				fmt.Println("Write error on data broadcast", err)
// 			}
// 		}(ws)
// 	}
// }

// func main() {
// 	server := NewServer()
// 	http.Handle("/ws", websocket.Handler(server.handleWS))

// 	http.ListenAndServe(":3000", nil)
// }
