package session

import (
	"context"
	"log"

	"github.com/snanovskyi/ooooooh/game"
	"github.com/snanovskyi/ooooooh/protocol/server"
	"github.com/snanovskyi/ooooooh/socket"
)

type Session struct {
	context context.Context
	socket  socket.Socket
	encoder server.Encoder
	player  *game.Player
}

func NewSession(ctx context.Context, s socket.Socket, e server.Encoder, p *game.Player) *Session {
	return &Session{
		context: ctx,
		socket:  s,
		encoder: e,
		player:  p,
	}
}

func (s *Session) Socket() socket.Socket {
	return s.socket
}

func (s *Session) Player() *game.Player {
	return s.player
}

func (s *Session) Send(m server.Message) {
	// TODO: error handling

	bytes, err := m.Encode(s.encoder)
	if err != nil {
		log.Println(err)
		return
	}
	err = s.socket.Write(s.context, bytes)
	if err != nil {
		log.Println(err)
		return
	}
}

func (s *Session) Close(status socket.CloseStatus) error {
	return s.socket.Close(status)
}
