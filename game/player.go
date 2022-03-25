package game

import (
	"github.com/snanovskyi/ooooooh/math"
)

type Player struct {
	world     *World
	id        uint32
	position  *math.Vector
	direction *math.Vector
	velocity  float32
}

func NewPlayer(world *World) *Player {
	return &Player{
		world:     world,
		id:        world.NewID(),
		direction: &math.Vector{},
		velocity:  30,
	}
}

func (p *Player) World() *World {
	return p.world
}

func (p *Player) ID() uint32 {
	return p.id
}

func (p *Player) Position() *math.Vector {
	return p.position
}

func (p *Player) Direction() *math.Vector {
	return p.direction
}

func (p *Player) Velocity() float32 {
	return p.velocity
}

func (p *Player) Spawn() {
	p.position = &math.Vector{}
	p.world.Handler().SpawnPlayer(p)
}

func (p *Player) Destroy() {
	p.position = nil
}

func (p *Player) Update() {
	p.position.Add(&math.Vector{
		X: p.direction.X * p.velocity,
		Y: p.direction.Y * p.velocity,
	})
}

func (p *Player) Move(direction *math.Vector) {
	p.direction = direction
	p.world.Handler().UpdatePlayer(p)
}
