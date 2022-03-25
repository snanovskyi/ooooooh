package server

type Message interface {
	Encode(encoder Encoder) ([]byte, error)
}
