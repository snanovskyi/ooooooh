package server

import (
	"github.com/snanovskyi/ooooooh/game"
)

type JoinGame struct {
	player *game.Player
}

func NewJoinGame(player *game.Player) *JoinGame {
	return &JoinGame{player: player}
}

func (j *JoinGame) Player() *game.Player {
	return j.player
}

func (j *JoinGame) Encode(encoder Encoder) ([]byte, error) {
	return encoder.JoinGame(j)
}
