package handler

import (
	"context"
	"whatsapp-api/internal/provider"
	proto "whatsapp-api/model/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type App struct {
	log        provider.ILogger
	grpcClient proto.WaCoreGatewayClient
	grpcConn   *grpc.ClientConn
}

type server struct {
	proto.UnimplementedWaCoreGatewayServer
}

func NewApp(log provider.ILogger) *App {
	return &App{log: log}
}

func (a *App) GRPCClient(addr string) (*grpc.ClientConn, error) {
	// Create gRPC connection with insecure credentials
	grpcConn, err := grpc.Dial(addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, err
	}

	// Store the gRPC client in the App struct
	a.grpcClient = proto.NewWaCoreGatewayClient(grpcConn)
	a.grpcConn = grpcConn

	return grpcConn, err
}

// GetGRPCClient returns the stored gRPC client
func (a *App) GetGRPCClient() proto.WaCoreGatewayClient {
	return a.grpcClient
}

// GetGRPCConnection returns the stored gRPC connection
func (a *App) GetGRPCConnection() *grpc.ClientConn {
	return a.grpcConn
}

// CloseGRPCConnection closes the gRPC connection
func (a *App) CloseGRPCConnection() error {
	if a.grpcConn != nil {
		return a.grpcConn.Close()
	}
	return nil
}

// IsGRPCConnected checks if gRPC client is available
func (a *App) IsGRPCConnected() bool {
	return a.grpcClient != nil && a.grpcConn != nil
}

// Example method to use the gRPC client - Get All Devices
func (a *App) GetAllDevices(ctx context.Context) (*proto.DeviceListResponse, error) {
	if !a.IsGRPCConnected() {
		return nil, ErrGRPCClientNotConnected
	}

	response, err := a.grpcClient.GetAllDevice(ctx, &emptypb.Empty{})
	return response, err
}

// Example method to get client contacts
func (a *App) GetClientContact(ctx context.Context, req *proto.ClientdataRequest) (*proto.ContactListResponse, error) {
	if !a.IsGRPCConnected() {
		return nil, ErrGRPCClientNotConnected
	}

	response, err := a.grpcClient.GetClientContact(ctx, req)
	return response, err
}

func (a *App) GetClientGroup(ctx context.Context, req *proto.ClientdataRequest) (*proto.GroupListResponse, error) {
	if !a.IsGRPCConnected() {
		return nil, ErrGRPCClientNotConnected
	}

	response, err := a.grpcClient.GetClientGroup(ctx, req)
	return response, err
}

// Example method to send message
func (a *App) SendMessage(ctx context.Context, messagePayload *proto.MessagePayload) (*proto.MessageResponse, error) {
	if !a.IsGRPCConnected() {
		return nil, ErrGRPCClientNotConnected
	}

	response, err := a.grpcClient.SendMessage(ctx, messagePayload)
	return response, err
}
