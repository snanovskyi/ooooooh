package handler

import (
	"github.com/snanovskyi/ooooooh/game"
	"github.com/snanovskyi/ooooooh/protocol/server"
	"github.com/snanovskyi/ooooooh/session"
)

type gameHandler struct {
	registry *session.Registry
}

func NewGameHandler(r *session.Registry) *gameHandler {
	return &gameHandler{registry: r}
}

func (g *gameHandler) DestroyEntity(entity game.Entity) {
	g.registry.Broadcast(server.NewDestroyEntity(entity))
}

func (g *gameHandler) SpawnPlayer(player *game.Player) {
	g.registry.Broadcast(server.NewSpawnPlayer(player))
}

func (g *gameHandler) UpdatePlayer(player *game.Player) {
	g.registry.Broadcast(server.NewUpdatePlayer(player))
}
