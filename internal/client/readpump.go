package client

import (
	"gochat/internal/models"
	"log"
)

func ReadPump(c *models.Client,room *models.Room) {
defer func() {
	room.Leave <- c
	c.Conn.Close()
}()

for{
	_, message, err := c.Conn.ReadMessage()
	if err != nil {
		log.Printf("Error reading message: %v", err)
		break
	}
	room.Broadcast <- message
}
}