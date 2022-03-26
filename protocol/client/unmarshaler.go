package client

type Unmarshaler interface {
	Unmarshal(bytes []byte) (Message, error)
}
