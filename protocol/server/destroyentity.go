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

func (d *DestroyEntity) Marshal(marshaler Marshaler) ([]byte, error) {
	return marshaler.DestroyEntity(d)
}
