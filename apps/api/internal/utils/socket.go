package utils

import (
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	Send chan []byte
}

type Socket struct {
	Register chan *Client
	Unregister chan *Client
	Broadcast chan []byte
	Clients map[*Client]bool
}

// var register = make(chan *Client)
// var unregister = make(chan *Client)

// var clients = make(map[*Client]bool)

func NewSocket() *Socket {
	return &Socket{
	Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Broadcast:  make(chan []byte),
		Clients:    make(map[*Client]bool),
	}
}



func (s *Socket) RunHub() {
	for {
		select{
		case client := <- s.Register:
			s.Clients[client] = true
			log.Println("user registered")
		case client:= <- s.Unregister:
			if _, ok := s.Clients[client]; ok {
				close(client.Send)
				delete(s.Clients,client)
			}
		case message := <- s.Broadcast:
			for client := range s.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(s.Clients,client)
				}
			}
		}
	
	}
}

func (c *Client) WritePump() {
	defer c.Conn.Close()
	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage,msg)
		if err != nil {
			log.Printf("Fail to Send message: %s", err)
			return
		}
	}
}