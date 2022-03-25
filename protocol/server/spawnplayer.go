package server

import (
	"github.com/snanovskyi/ooooooh/game"
)

type SpawnPlayer struct {
	player *game.Player
}

func NewSpawnPlayer(player *game.Player) *SpawnPlayer {
	return &SpawnPlayer{player: player}
}

func (s *SpawnPlayer) Player() *game.Player {
	return s.player
}

func (s *SpawnPlayer) Encode(encoder Encoder) ([]byte, error) {
	return encoder.SpawnPlayer(s)
}
