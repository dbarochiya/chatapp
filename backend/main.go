package main

import (
	"fmt"
	"net/http"

	"github.com/dbarochiya/chatapp/pkg/websocket"
)

func serveRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am root path ")
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func main() {
	http.HandleFunc("/root", serveRoot)
	http.HandleFunc("/ws", serveWs)
	http.ListenAndServe(":8080", nil)
}
