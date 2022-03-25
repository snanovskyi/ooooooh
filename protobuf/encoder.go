package protobuf

import (
	"github.com/snanovskyi/ooooooh/math"
	"github.com/snanovskyi/ooooooh/protocol/server"
	"google.golang.org/protobuf/proto"
)

type encoder struct {
}

func NewEncoder() *encoder {
	return &encoder{}
}

func (e *encoder) Pong(pong *server.Pong) ([]byte, error) {
	return proto.Marshal(&Message{
		Opcode: Message_SERVER_PONG,
		Pong: &Pong{
			Id: pong.ID(),
		},
	})
}

func (e *encoder) JoinGame(joinGame *server.JoinGame) ([]byte, error) {
	entities := joinGame.Player().World().Entities()
	players := make([]*JoinGame_Player, len(entities))
	for i, entity := range entities {
		players[i] = &JoinGame_Player{
			Id:        entity.ID(),
			Position:  EncodeVector(entity.Position()),
			Direction: EncodeVector(&math.Vector{}),
			Velocity:  0,
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

func (e *encoder) DestroyEntity(destroyEntity *server.DestroyEntity) ([]byte, error) {
	return proto.Marshal(&Message{
		Opcode: Message_SERVER_DESTROY_ENTITY,
		DestroyEntity: &DestroyEntity{
			Id: destroyEntity.Entity().ID(),
		},
	})
}

func (e *encoder) SpawnPlayer(spawnPlayer *server.SpawnPlayer) ([]byte, error) {
	return proto.Marshal(&Message{
		Opcode: Message_SERVER_SPAWN_PLAYER,
		SpawnPlayer: &SpawnPlayer{
			Id:       spawnPlayer.Player().ID(),
			Position: EncodeVector(spawnPlayer.Player().Position()),
		},
	})
}

func (e *encoder) UpdatePlayer(updatePlayer *server.UpdatePlayer) ([]byte, error) {
	return proto.Marshal(&Message{
		Opcode: Message_SERVER_UPDATE_PLAYER,
		UpdatePlayer: &UpdatePlayer{
			Id:        updatePlayer.Player().ID(),
			Position:  EncodeVector(updatePlayer.Player().Position()),
			Direction: EncodeVector(updatePlayer.Player().Direction()),
			Velocity:  updatePlayer.Player().Velocity(),
		},
	})
}
