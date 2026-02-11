package client

import (
	"gochat/internal/models"
	"github.com/gorilla/websocket"
	"github.com/google/uuid"
)

func NewClient(conn *websocket.Conn, username string) *models.Client {
	return &models.Client{
		Conn: conn,
		Send: make(chan []byte),
		ID: uuid.NewString(),
		Username: username,
		RoomID: []int{},
	}
}