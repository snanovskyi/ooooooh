package session

import (
	"context"
	"log"

	"github.com/snanovskyi/ooooooh/game"
	"github.com/snanovskyi/ooooooh/protocol/server"
	"github.com/snanovskyi/ooooooh/socket"
)

type Session struct {
	context   context.Context
	socket    socket.Socket
	marshaler server.Marshaler
	player    *game.Player
}

func NewSession(ctx context.Context, s socket.Socket, m server.Marshaler, p *game.Player) *Session {
	return &Session{
		context:   ctx,
		socket:    s,
		marshaler: m,
		player:    p,
	}
}

func (s *Session) Socket() socket.Socket {
	return s.socket
}

func (s *Session) Player() *game.Player {
	return s.player
}

func (s *Session) Send(message server.Message) {
	bytes, err := message.Marshal(s.marshaler)
	if err != nil {
		log.Println(err)
		s.Close(socket.StatusProtocolError)
		return
	}
	err = s.socket.Write(s.context, bytes)
	if err != nil {
		log.Println(err)
		s.Close(socket.StatusInternalError)
		return
	}
}

func (s *Session) Close(status socket.CloseStatus) error {
	return s.socket.Close(status)
}
