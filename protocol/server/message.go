package server

type Message interface {
	Marshal(marshaler Marshaler) ([]byte, error)
}
