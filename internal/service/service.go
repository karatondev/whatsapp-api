package service

import (
	"context"
	"whatsapp-api/internal/handler"
	"whatsapp-api/internal/provider"
	"whatsapp-api/internal/repository"
	"whatsapp-api/model"

	"github.com/redis/go-redis/v9"
)

type MessagesApi interface {
	PushMessage(ctx context.Context, msg *model.MessageRequest) (*model.MessageResponse, error)
	GetContacts(ctx context.Context, senderJid string) (*model.Contacts, error)
	GetGroups(ctx context.Context, senderJid string) (*model.Groups, error)
}

type service struct {
	repo   repository.PostgresRepository
	logger provider.ILogger
	grpc   *handler.App
	redis  *redis.Client
}

func NewService(repo repository.PostgresRepository, logger provider.ILogger, grpc *handler.App, redis *redis.Client) MessagesApi {
	return &service{
		repo:   repo,
		logger: logger,
		grpc:   grpc,
		redis:  redis,
	}
}
