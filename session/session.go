package session

import (
	"context"
	"log"

	"github.com/snanovskyi/ooooooh/game"
	"github.com/snanovskyi/ooooooh/protocol"
	"github.com/snanovskyi/ooooooh/protocol/server"
	"github.com/snanovskyi/ooooooh/socket"
)

type Session struct {
	context context.Context
	socket  socket.Socket
	codec   *protocol.Codec
	player  *game.Player
}

func NewSession(ctx context.Context, s socket.Socket, c *protocol.Codec, p *game.Player) *Session {
	return &Session{
		context: ctx,
		socket:  s,
		codec:   c,
		player:  p,
	}
}

func (s *Session) Context() context.Context {
	return s.context
}

func (s *Session) Socket() socket.Socket {
	return s.socket
}

func (s *Session) Player() *game.Player {
	return s.player
}

func (s *Session) Send(m server.Message) {
	// TODO: error handling

	bytes, err := s.codec.Encode(m)
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

func (s *Session) Close() error {
	return s.socket.Close()
}
