package handler

import (
	"github.com/snanovskyi/ooooooh/protocol/client"
	"github.com/snanovskyi/ooooooh/protocol/server"
	"github.com/snanovskyi/ooooooh/session"
	"github.com/snanovskyi/ooooooh/ticker"
)

type clientHandler struct {
	ticker  *ticker.Ticker
	session *session.Session
}

func NewClientHandler(t *ticker.Ticker, s *session.Session) *clientHandler {
	return &clientHandler{
		ticker:  t,
		session: s,
	}
}

func (c *clientHandler) Ping(ping *client.Ping) {
	c.session.Send(server.NewPong(ping.ID()))
}

func (c *clientHandler) MovePlayer(movePlayer *client.MovePlayer) {
	c.ticker.NextTick(func() {
		c.session.Player().Move(movePlayer.Direction())
	})
}
