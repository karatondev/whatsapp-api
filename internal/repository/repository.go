package repository

import (
	"context"
	"whatsapp-api/internal/provider"
	entity "whatsapp-api/model/entity"

	"github.com/jackc/pgx/v5"
)

type PostgresRepository interface {
	FindDeviceByID(ctx context.Context, id string) (*entity.Device, error)
}

type repo struct {
	logger provider.ILogger
	conn   *pgx.Conn
}

func NewPostgresRepository(logger provider.ILogger, conn *pgx.Conn) PostgresRepository {
	return &repo{
		logger: logger,
		conn:   conn,
	}
}
