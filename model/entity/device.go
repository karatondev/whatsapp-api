package model

import "time"

// Device represents a device entity mapped to the devices table.
type Device struct {
	DeviceID      string     `json:"device_id" db:"device_id"`
	DeviceName    string     `json:"device_name" db:"device_name"`
	SenderJID     string     `json:"sender_jid" db:"sender_jid"`
	InitiatedAt   *time.Time `json:"initiated_at,omitempty" db:"initiated_at"`
	ConnectedAt   *time.Time `json:"connected_at,omitempty" db:"connected_at"`
	DeviceAlias   string     `json:"device_alias" db:"device_alias"`
	ConnectStatus string     `json:"connect_status" db:"connect_status"`
	CreatedAt     time.Time  `json:"created_at" db:"created_at"`
	CreatedBy     string     `json:"created_by" db:"created_by"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}
