package socket

import (
	"context"
)

type Handler interface {
	Open(ctx context.Context, socket Socket)
	Message(ctx context.Context, socket Socket, bytes []byte)
	Close(ctx context.Context, socket Socket)
	Error(ctx context.Context, socket Socket, err error)
}
