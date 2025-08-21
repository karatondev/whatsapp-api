package repository

import (
	"context"
	"fmt"
	entity "whatsapp-api/model/entity"

	"github.com/jackc/pgx/v5"
)

// SaveEvent saves an inbound event to the database

const (
	findDeviceByIDQuery = `
		SELECT device_id, device_name, sender_jid, initiated_at, connected_at, device_alias, connect_status, created_at, created_by, updated_at, deleted_at
		FROM devices
		WHERE device_id = @device_id`
)

// SaveInboundEvent saves an inbound event to the inbound table
func (r *repo) FindDeviceByID(ctx context.Context, id string) (*entity.Device, error) {
	params := pgx.NamedArgs{
		"device_id": id,
	}

	var device entity.Device
	if err := r.conn.QueryRow(ctx, findDeviceByIDQuery, params).Scan(
		&device.DeviceID,
		&device.DeviceName,
		&device.SenderJID,
		&device.InitiatedAt,
		&device.ConnectedAt,
		&device.DeviceAlias,
		&device.ConnectStatus,
		&device.CreatedAt,
		&device.CreatedBy,
		&device.UpdatedAt,
		&device.DeletedAt,
	); err != nil {
		if pgx.ErrNoRows == err {
			return nil, ErrDeviceNotFound
		}
		err := fmt.Errorf("failed to scan device: %w", err)
		return nil, err
	}

	return &device, nil
}
