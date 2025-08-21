package repository

import "errors"

var (
	ErrDeviceNotFound = errors.New("Device not found")
	ErrDataIsEmpty    = errors.New("Data is empty")
)
