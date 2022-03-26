package websocket

import (
	"bytes"
	"context"
	"fmt"
	"sync"

	sock "github.com/snanovskyi/ooooooh/socket"
	"nhooyr.io/websocket"
)

var closeStatusToStatusCode = map[sock.CloseStatus]websocket.StatusCode{
	sock.StatusOk:            websocket.StatusNormalClosure,
	sock.StatusProtocolError: websocket.StatusUnsupportedData,
	sock.StatusInternalError: websocket.StatusInternalError,
}

type socket struct {
	conn   *websocket.Conn
	mu     sync.Mutex
	closed bool
}

func newSocket(c *websocket.Conn) *socket {
	return &socket{conn: c}
}

func (s *socket) Connected() bool {
	return !s.closed
}

func (s *socket) Closed() bool {
	return s.closed
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

func (s *socket) Close(status sock.CloseStatus) error {
	s.closed = true
	return s.conn.Close(closeStatusToStatusCode[status], fmt.Sprintf("close status '%d'", status))
}
