package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"whatsapp-api/internal/handler"
	"whatsapp-api/internal/handler/rest"
	"whatsapp-api/internal/provider"
	"whatsapp-api/internal/repository"
	"whatsapp-api/internal/service"
	"whatsapp-api/model/constant"
	"whatsapp-api/util"

	"google.golang.org/grpc/connectivity"
)

func Run(cfg *util.Config) {
	ctx := context.WithValue(context.Background(), constant.CtxReqIDKey, "MAIN")

	logger := provider.NewLogger()

	postgres, err := provider.NewPostgresConnection(ctx)
	if err != nil {
		logger.Errorfctx(provider.AppLog, ctx, false, "Failed connect to Postgres: %v", err)
		return
	}

	redis, err := provider.NewRedisConnection(ctx)
	if err != nil {
		logger.Errorfctx(provider.AppLog, ctx, false, "Failed connect to Redis: %v", err)
		return
	}

	logger.Infofctx(provider.AppLog, ctx, "Application started")

	app := handler.NewApp(logger)
	repo := repository.NewPostgresRepository(logger, postgres)
	svc := service.NewService(repo, logger, app, redis)

	go func() {
		// Setup gRPC client connection
		grpcServerAddr := cfg.Grpc.DSN
		logger.Infofctx(provider.AppLog, ctx, "Starting gRPC client for server at %s", grpcServerAddr)
		conn, err := app.GRPCClient(grpcServerAddr)
		if err != nil {
			logger.Errorfctx(provider.AppLog, ctx, false, "Failed to create gRPC connection: %v", err)
			return
		}
		defer app.CloseGRPCConnection()

		logger.Infofctx(provider.AppLog, ctx, "gRPC client connected successfully to %s", grpcServerAddr)

		// Monitor connection state
		for {
			select {
			case <-ctx.Done():
				logger.Infofctx(provider.AppLog, ctx, "gRPC client shutting down")
				return
			default:
				// Check connection state
				state := conn.GetState()

				switch state {
				case connectivity.Ready:
					logger.Debugfctx(provider.AppLog, ctx, "gRPC connection is ready")
				case connectivity.Connecting:
					logger.Infofctx(provider.AppLog, ctx, "gRPC connection is connecting...")
				case connectivity.TransientFailure:
					logger.Errorfctx(provider.AppLog, ctx, false, "gRPC connection in transient failure state")
					// Wait for state change or timeout
					if !conn.WaitForStateChange(ctx, state) {
						logger.Errorfctx(provider.AppLog, ctx, false, "Context cancelled while waiting for state change")
						return
					}
				case connectivity.Idle:
					logger.Infofctx(provider.AppLog, ctx, "gRPC connection is idle")
				case connectivity.Shutdown:
					logger.Errorfctx(provider.AppLog, ctx, false, "gRPC connection is shutdown")
					return
				}

				// Wait before next state check
				time.Sleep(10 * time.Second)
			}
		}
	}()

	server := &http.Server{}
	go func(logger provider.ILogger, svc service.MessagesApi) {
		app := rest.NewRest(logger, svc)
		addr := fmt.Sprintf(":%v", util.Configuration.Server.Port)
		server, err = app.CreateServer(addr)
		if err != nil {
			logger.Errorfctx(provider.AppLog, ctx, false, "Failed to create server: %v", err)
		}

		logger.Infofctx(provider.AppLog, ctx, "Server running at: %s", addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Errorfctx(provider.AppLog, ctx, false, "Server error: %v", err)
		}

	}(logger, svc)

	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	sig := <-shutdownCh
	logger.Infofctx(provider.AppLog, ctx, "Receiving signal: %s", sig)

	func(logger provider.ILogger) {
		shutdownCtx, cancel := context.WithTimeout(ctx, util.Configuration.Server.ShutdownTimeout)
		defer cancel()
		server.Shutdown(shutdownCtx)

		logger.Infofctx(provider.AppLog, ctx, "Successfully stop Application.")
	}(logger)

}
