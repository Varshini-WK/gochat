package models

type Room struct {
	RoomID int
	Name   string
	Clients map[*Client]bool
	Broadcast chan []byte
	Join chan *Client
	Leave chan *Client
}
