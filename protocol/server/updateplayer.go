package server

import (
	"github.com/snanovskyi/ooooooh/game"
)

type UpdatePlayer struct {
	player *game.Player
}

func NewUpdatePlayer(player *game.Player) *UpdatePlayer {
	return &UpdatePlayer{player: player}
}

func (u *UpdatePlayer) Player() *game.Player {
	return u.player
}

func (u *UpdatePlayer) Marshal(marshaler Marshaler) ([]byte, error) {
	return marshaler.UpdatePlayer(u)
}
