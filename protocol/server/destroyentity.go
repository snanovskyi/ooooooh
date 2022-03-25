package server

import (
	"github.com/snanovskyi/ooooooh/game"
)

type DestroyEntity struct {
	entity game.Entity
}

func NewDestroyEntity(entity game.Entity) *DestroyEntity {
	return &DestroyEntity{entity: entity}
}

func (d *DestroyEntity) Entity() game.Entity {
	return d.entity
}

func (d *DestroyEntity) Encode(encoder Encoder) ([]byte, error) {
	return encoder.DestroyEntity(d)
}
