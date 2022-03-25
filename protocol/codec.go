package protocol

import (
	"github.com/snanovskyi/ooooooh/protocol/client"
	"github.com/snanovskyi/ooooooh/protocol/server"
)

type Codec struct {
	decoder client.Decoder
	encoder server.Encoder
}

func NewCodec(d client.Decoder, e server.Encoder) *Codec {
	return &Codec{
		decoder: d,
		encoder: e,
	}
}

func (c *Codec) Decode(bytes []byte) (client.Message, error) {
	return c.decoder.Decode(bytes)
}

func (c *Codec) Encode(message server.Message) ([]byte, error) {
	return message.Encode(c.encoder)
}
