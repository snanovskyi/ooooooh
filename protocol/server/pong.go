package server

type Pong struct {
	id uint32
}

func NewPong(id uint32) *Pong {
	return &Pong{id: id}
}

func (p *Pong) ID() uint32 {
	return p.id
}

func (p *Pong) Marshal(marshaler Marshaler) ([]byte, error) {
	return marshaler.Pong(p)
}
