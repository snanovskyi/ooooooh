package client

type Decoder interface {
	Decode(bytes []byte) (Message, error)
}
