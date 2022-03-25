package client

type Message interface {
	Handle(handler Handler)
}
