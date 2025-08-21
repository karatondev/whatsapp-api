package model

type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Short string `json:"short,omitempty"`
}

type Contacts struct {
	DeviceId      string    `json:"device_id"`
	DeviceName    string    `json:"device_name"`
	DeviceAlias   string    `json:"device_alias"`
	ConnectStatus string    `json:"connect_status"`
	Contacts      []Contact `json:"contacts"`
}
