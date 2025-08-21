package model

type Group struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Short string `json:"short,omitempty"`
}

type Groups struct {
	DeviceId    string  `json:"device_id"`
	DeviceName  string  `json:"device_name"`
	DeviceAlias string  `json:"device_alias"`
	Groups      []Group `json:"groups"`
}
