package model

import (
	proto "whatsapp-api/model/pb"
)

// MessageRequest represents the request payload for sending a message.
type MessageRequest struct {
	DeviceID string          `json:"device"`
	To       string          `json:"to"`
	Type     string          `json:"type"`
	Text     string          `json:"text,omitempty"`
	Image    *proto.Media    `json:"image,omitempty"`
	Video    *proto.Media    `json:"video,omitempty"`
	Document *proto.Document `json:"document,omitempty"`
	Audio    *proto.Audio    `json:"audio,omitempty"`
	Location *proto.Location `json:"location,omitempty"`
}
