package client

import (
	"gochat/internal/models"
	"github.com/gorilla/websocket"
)

func WritePump(c *models.Client) {
	defer func() {
		c.Conn.Close()
	}()
	for msg := range c.Send {
		err := c.Conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			log.Printf("Error writing message: %v", err)
			break
		}
	}
}