package model

type Message struct {
	ID string `json:"id"`
}

type MessageResponse struct {
	Message Message `json:"message"`
}
