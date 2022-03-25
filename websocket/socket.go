package websocket

import (
	"bytes"
	"context"
	"fmt"
	"sync"

	"nhooyr.io/websocket"
)

type socket struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func newSocket(c *websocket.Conn) *socket {
	return &socket{conn: c}
}

func (s *socket) Read(ctx context.Context) ([]byte, error) {
	// Only one Reader may be open at a time.
	s.mu.Lock()
	defer s.mu.Unlock()

	t, r, err := s.conn.Reader(ctx)
	if err != nil {
		return nil, err
	}

	if t != websocket.MessageBinary {
		return nil, fmt.Errorf("expected message type '%d' but received '%d'", websocket.MessageBinary, t)
	}

	// TODO: reuse buffers
	b := &bytes.Buffer{}
	_, err = b.ReadFrom(r)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func (s *socket) Write(ctx context.Context, bytes []byte) error {
	return s.conn.Write(ctx, websocket.MessageBinary, bytes)
}

func (s *socket) Close() error {
	return s.conn.Close(websocket.StatusNormalClosure, "close")
}
