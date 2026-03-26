package handlers

import (
	"akademik-backend/ws"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Dalam tahap development bebaskan cross origin
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Variabel global untuk menampung instansi Hub
var AppHub *ws.Hub

// ServeWS menangani request masuk HTTP menjadi koneksi Websocket
func ServeWS(c *gin.Context) {
	// Pada skenario Production, sebaiknya kita validasi token JWT 
	// yang dikirim melalui parameter query URL (karena WS tidak punya header Authorization standar di JS Browser API awam)
	// uint(c.Get("user_id")) bisa digunakan jika kita parsing JWT di middleware
	
	userIDInterface, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized WS Access"})
		return
	}
	
	userID := uint(userIDInterface.(float64))

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	client := &ws.Client{
		Hub:    AppHub,
		UserID: userID,
		Conn:   conn,
		Send:   make(chan []byte, 256),
	}
	client.Hub.Register <- client

	// Jalankan rutin background untuk membaca dan menulis dari/ke client websocket
	go client.WritePump()
	go client.ReadPump()
}
