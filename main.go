package main

import (
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func main() {
	server := NewServer()
	http.Handle("/chat", websocket.Handler(server.serveWS))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
