package game

import (
	"github.com/snanovskyi/ooooooh/math"
)

type Entity interface {
	World() *World
	ID() uint32
	Position() *math.Vector
	Spawn()
	Destroy()
	Update()
}
