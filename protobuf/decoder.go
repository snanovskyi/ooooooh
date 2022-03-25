package protobuf

import (
	"fmt"

	"github.com/snanovskyi/ooooooh/protocol/client"
	"google.golang.org/protobuf/proto"
)

type Decoder struct {
}

func (d *Decoder) Decode(bytes []byte) (client.Message, error) {
	message := &Message{}
	if err := proto.Unmarshal(bytes, message); err != nil {
		return nil, err
	}
	switch message.Opcode {
	case Message_CLIENT_PING:
		return d.ping(message)
	case Message_CLIENT_MOVE_PLAYER:
		return d.movePlayer(message)
	default:
		return nil, fmt.Errorf("unknown opcode '%d'", message.Opcode)
	}
}

func (d *Decoder) ping(message *Message) (client.Message, error) {
	return client.NewPing(message.Ping.Id), nil
}

func (d *Decoder) movePlayer(message *Message) (client.Message, error) {
	return client.NewMovePlayer(DecodeVector(message.MovePlayer.Direction)), nil
}
