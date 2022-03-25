package websocket

import (
	"net/http"

	sock "github.com/snanovskyi/ooooooh/socket"
	"nhooyr.io/websocket"
)

type handler struct {
	handler sock.Handler
}

func NewHandler(h sock.Handler) http.Handler {
	return &handler{handler: h}
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		// TODO: fix this later
		InsecureSkipVerify: true,
	})
	if err != nil {
		h.handler.Error(ctx, nil, err)
		return
	}

	s := newSocket(conn)
	h.handler.Open(ctx, s)
	defer h.handler.Close(s)

	for !s.Closed() {
		select {
		case <-ctx.Done():
			h.handler.Error(ctx, s, ctx.Err())
			break
		default:
			bytes, rErr := s.Read(ctx)
			if rErr != nil {
				h.handler.Error(ctx, s, rErr)
				continue
			}
			h.handler.Message(ctx, s, bytes)
		}
	}
}
