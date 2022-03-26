package socket

import (
	"context"
)

type CloseStatus int

const (
	StatusOk CloseStatus = iota + 1
	StatusProtocolError
	StatusInternalError
)

type Socket interface {
	Write(ctx context.Context, bytes []byte) error
	Close(status CloseStatus) error
}
