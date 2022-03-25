package client

import (
	"github.com/snanovskyi/ooooooh/math"
)

type MovePlayer struct {
	direction *math.Vector
}

func NewMovePlayer(direction *math.Vector) *MovePlayer {
	return &MovePlayer{direction: direction}
}

func (m *MovePlayer) Direction() *math.Vector {
	return m.direction
}

func (m *MovePlayer) Handle(handler Handler) {
	handler.MovePlayer(m)
}
