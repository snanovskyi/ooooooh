package socket

import (
	"context"
)

type Socket interface {
	Read(ctx context.Context) ([]byte, error)
	Write(ctx context.Context, bytes []byte) error
	Close() error
}
