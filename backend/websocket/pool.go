package websocket

import "fmt"

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("New clinet joined : ", client.ID, " & Size of Connection Pool: ", len(pool.Clients))
			connetMessage := Message{Type: 1, Body: "New User Joined...", Sender: client.ID}
			for client := range pool.Clients {
				client.Conn.WriteJSON(connetMessage)
			}
			break
		case client := <-pool.Unregister:
			fmt.Println("Clinet disconnected :", client.ID, " & Size of Connection Pool: ", len(pool.Clients))
			leaveMessage := Message{Type: 1, Body: "User Disconnected...", Sender: client.ID}
			for client := range pool.Clients {
				client.Conn.WriteJSON(leaveMessage)
			}
			delete(pool.Clients, client)
			break
		case message := <-pool.Broadcast:
			fmt.Println("Sending message to all clients in Pool")
			for client := range pool.Clients {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
