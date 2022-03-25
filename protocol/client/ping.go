package client

type Ping struct {
	id uint32
}

func NewPing(id uint32) *Ping {
	return &Ping{id: id}
}

func (p *Ping) ID() uint32 {
	return p.id
}

func (p *Ping) Handle(handler Handler) {
	handler.Ping(p)
}
