package handler

import (
	"context"
	"log"
	"sync"

	"github.com/snanovskyi/ooooooh/game"
	"github.com/snanovskyi/ooooooh/protocol/client"
	"github.com/snanovskyi/ooooooh/protocol/server"
	"github.com/snanovskyi/ooooooh/session"
	"github.com/snanovskyi/ooooooh/socket"
	"github.com/snanovskyi/ooooooh/ticker"
)

type socketHandler struct {
	registry *session.Registry
	ticker   *ticker.Ticker
	world    *game.World
	decoder  client.Decoder
	encoder  server.Encoder
	mu       sync.RWMutex
	handler  map[socket.Socket]client.Handler
}

func NewSocketHandler(r *session.Registry, t *ticker.Ticker, w *game.World, d client.Decoder, e server.Encoder) *socketHandler {
	return &socketHandler{
		registry: r,
		ticker:   t,
		world:    w,
		decoder:  d,
		encoder:  e,
		handler:  make(map[socket.Socket]client.Handler),
	}
}

func (s *socketHandler) Open(ctx context.Context, sock socket.Socket) {
	s.mu.Lock()
	defer s.mu.Unlock()
	newSession := session.NewSession(ctx, sock, s.encoder, game.NewPlayer(s.world))
	s.registry.Add(newSession)
	s.handler[newSession.Socket()] = NewClientHandler(s.ticker, newSession)
	s.ticker.NextTick(func() {
		s.world.Spawn(newSession.Player())
		newSession.Send(server.NewJoinGame(newSession.Player()))
	})
}

func (s *socketHandler) Message(_ context.Context, sock socket.Socket, bytes []byte) {
	message, err := s.decoder.Decode(bytes)
	if err != nil {
		s.registry.Get(sock).Close(socket.StatusProtocolError)
		return
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	message.Handle(s.handler[sock])
}

func (s *socketHandler) Close(sock socket.Socket) {
	s.mu.Lock()
	defer s.mu.Unlock()
	getSession := s.registry.Get(sock)
	s.registry.Delete(sock)
	delete(s.handler, getSession.Socket())
	s.ticker.NextTick(func() {
		s.world.Destroy(getSession.Player())
	})
}

func (s *socketHandler) Error(_ context.Context, sock socket.Socket, err error) {
	// TODO: error handling
	log.Println(err)
	getSession := s.registry.Get(sock)
	getSession.Close(socket.StatusInternalError)
}
