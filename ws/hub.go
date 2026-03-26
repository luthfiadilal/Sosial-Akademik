package ws

import (
	"log"
)

// Hub maintains the set of active clients and broadcasts messages to the clients.
type Hub struct {
	// Registered clients mapped by UserID. allows finding all active sessions of a specific user.
	Clients map[uint]map[*Client]bool

	// Inbound messages from the clients.
	Broadcast chan []byte

	// Register requests from the clients.
	Register chan *Client

	// Unregister requests from clients.
	Unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[uint]map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			// Initialize map for the user if it doesn't exist yet
			if _, ok := h.Clients[client.UserID]; !ok {
				h.Clients[client.UserID] = make(map[*Client]bool)
			}
			h.Clients[client.UserID][client] = true
			log.Printf("Client registered for User ID: %d", client.UserID)

		case client := <-h.Unregister:
			if connections, ok := h.Clients[client.UserID]; ok {
				if _, exists := connections[client]; exists {
					delete(connections, client)
					close(client.Send)
					
					// Cleanup if no more connections for this user
					if len(connections) == 0 {
						delete(h.Clients, client.UserID)
					}
					log.Printf("Client unregistered for User ID: %d", client.UserID)
				}
			}

		case message := <-h.Broadcast:
			// Broadcast applies to all connected clients temporarily.
			// Ideally, we'd parse the message to see the target UserID,
			// but for generic broadcasts we can iterate over all.
			for _, userConns := range h.Clients {
				for client := range userConns {
					select {
					case client.Send <- message:
					default:
						close(client.Send)
						delete(userConns, client)
					}
				}
			}
		}
	}
}

// SendToUser sends a message to a specific user's active connections
func (h *Hub) SendToUser(userID uint, message []byte) {
	if userConns, ok := h.Clients[userID]; ok {
		for client := range userConns {
			select {
			case client.Send <- message:
			default:
				close(client.Send)
				delete(userConns, client)
			}
		}
	}
}
