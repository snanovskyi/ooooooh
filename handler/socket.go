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
	registry    *session.Registry
	ticker      *ticker.Ticker
	world       *game.World
	unmarshaler client.Unmarshaler
	marshaler   server.Marshaler
	mu          sync.RWMutex
	handler     map[socket.Socket]client.Handler
}

func NewSocketHandler(r *session.Registry, t *ticker.Ticker, w *game.World, u client.Unmarshaler, m server.Marshaler) *socketHandler {
	return &socketHandler{
		registry:    r,
		ticker:      t,
		world:       w,
		unmarshaler: u,
		marshaler:   m,
		handler:     make(map[socket.Socket]client.Handler),
	}
}

func (s *socketHandler) Open(ctx context.Context, sock socket.Socket) {
	s.mu.Lock()
	defer s.mu.Unlock()
	newSession := session.NewSession(ctx, sock, s.marshaler, game.NewPlayer(s.world))
	s.registry.Add(newSession)
	s.handler[newSession.Socket()] = NewClientHandler(s.ticker, newSession)
	s.ticker.NextTick(func() {
		s.world.Spawn(newSession.Player())
		newSession.Send(server.NewJoinGame(newSession.Player()))
	})
}

func (s *socketHandler) Message(_ context.Context, sock socket.Socket, bytes []byte) {
	message, err := s.unmarshaler.Unmarshal(bytes)
	if err != nil {
		log.Println(err)
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
	// Connection error
	if sock == nil {
		log.Println(err)
		return
	}

	// TODO: error handling
	log.Println(err)
	getSession := s.registry.Get(sock)
	getSession.Close(socket.StatusInternalError)
}
