package protobuf

import (
	"github.com/snanovskyi/ooooooh/game"
	"github.com/snanovskyi/ooooooh/math"
	"github.com/snanovskyi/ooooooh/protocol/server"
	"google.golang.org/protobuf/proto"
)

type Marshaler struct{}

func (m *Marshaler) Pong(pong *server.Pong) ([]byte, error) {
	return proto.Marshal(&Message{
		Opcode: Message_SERVER_PONG,
		Pong: &Pong{
			Id: pong.ID(),
		},
	})
}

func (m *Marshaler) JoinGame(joinGame *server.JoinGame) ([]byte, error) {
	entities := joinGame.Player().World().Entities()
	players := make([]*JoinGame_Player, len(entities))
	for i, entity := range entities {
		// TODO: fix this later
		p := entity.(*game.Player)
		players[i] = &JoinGame_Player{
			Id:        p.ID(),
			Position:  m.encodeVector(p.Position()),
			Direction: m.encodeVector(p.Direction()),
			Velocity:  p.Velocity(),
		}
	}
	return proto.Marshal(&Message{
		Opcode: Message_SERVER_JOIN_GAME,
		JoinGame: &JoinGame{
			PlayerId: joinGame.Player().ID(),
			Players:  players,
		},
	})
}

func (m *Marshaler) DestroyEntity(destroyEntity *server.DestroyEntity) ([]byte, error) {
	return proto.Marshal(&Message{
		Opcode: Message_SERVER_DESTROY_ENTITY,
		DestroyEntity: &DestroyEntity{
			Id: destroyEntity.Entity().ID(),
		},
	})
}

func (m *Marshaler) SpawnPlayer(spawnPlayer *server.SpawnPlayer) ([]byte, error) {
	return proto.Marshal(&Message{
		Opcode: Message_SERVER_SPAWN_PLAYER,
		SpawnPlayer: &SpawnPlayer{
			Id:       spawnPlayer.Player().ID(),
			Position: m.encodeVector(spawnPlayer.Player().Position()),
		},
	})
}

func (m *Marshaler) UpdatePlayer(updatePlayer *server.UpdatePlayer) ([]byte, error) {
	return proto.Marshal(&Message{
		Opcode: Message_SERVER_UPDATE_PLAYER,
		UpdatePlayer: &UpdatePlayer{
			Id:        updatePlayer.Player().ID(),
			Position:  m.encodeVector(updatePlayer.Player().Position()),
			Direction: m.encodeVector(updatePlayer.Player().Direction()),
			Velocity:  updatePlayer.Player().Velocity(),
		},
	})
}

func (m *Marshaler) encodeVector(vector *math.Vector) *Vector {
	return &Vector{
		X: vector.X,
		Y: vector.Y,
	}
}
