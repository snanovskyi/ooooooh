package protobuf

import (
	"fmt"

	"github.com/snanovskyi/ooooooh/math"
	"github.com/snanovskyi/ooooooh/protocol/client"
	"google.golang.org/protobuf/proto"
)

type Unmarshaler struct{}

func (u *Unmarshaler) Unmarshal(bytes []byte) (client.Message, error) {
	message := &Message{}
	if err := proto.Unmarshal(bytes, message); err != nil {
		return nil, err
	}
	switch message.Opcode {
	case Message_CLIENT_PING:
		return u.ping(message)
	case Message_CLIENT_MOVE_PLAYER:
		return u.movePlayer(message)
	default:
		return nil, fmt.Errorf("unknown opcode '%d'", message.Opcode)
	}
}

func (u *Unmarshaler) ping(message *Message) (client.Message, error) {
	return client.NewPing(message.Ping.Id), nil
}

func (u *Unmarshaler) movePlayer(message *Message) (client.Message, error) {
	return client.NewMovePlayer(u.decodeVector(message.MovePlayer.Direction)), nil
}

func (u *Unmarshaler) decodeVector(vector *Vector) *math.Vector {
	return &math.Vector{
		X: vector.X,
		Y: vector.Y,
	}
}
