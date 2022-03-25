package math

type Vector struct {
	X float32
	Y float32
}

func (v *Vector) Add(other *Vector) {
	v.X += other.X
	v.Y += other.Y
}
