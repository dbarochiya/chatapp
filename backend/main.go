package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/dbarochiya/chatapp/websocket"
)

var ClientC = 0

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	ClientC++

	client := &websocket.Client{
		ID:   strconv.Itoa(ClientC),
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

	// http.HandleFunc("/createpool/:name", func(w http.ResponseWriter, r *http.Request) {
	// 	createa a Pool with name (if given otherwise random generator)
	// })

	// http.HandleFunc("/joinpool/:name", func(w http.ResponseWriter, r *http.Request) {
	// 	serveWs(pool, w, r)
	// })
}

func main() {
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
