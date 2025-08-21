package handler

type GrpcErrorType int

const (
	ConnectionError GrpcErrorType = iota
)

type GrpcError struct {
	Type    GrpcErrorType
	Message string
}

func (e *GrpcError) Error() string {
	return e.Message
}

var (
	ErrGRPCClientNotConnected = &GrpcError{Type: ConnectionError, Message: "gRPC client not connected"}
)
