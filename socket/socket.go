package socket

import (
	"context"
)

type Socket interface {
	Connected() bool
	Closed() bool
	Read(ctx context.Context) ([]byte, error)
	Write(ctx context.Context, bytes []byte) error
	Close() error
}
