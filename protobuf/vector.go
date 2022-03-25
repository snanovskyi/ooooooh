package protobuf

import (
	"github.com/snanovskyi/ooooooh/math"
)

func DecodeVector(vector *Vector) *math.Vector {
	return &math.Vector{
		X: vector.X,
		Y: vector.Y,
	}
}

func EncodeVector(vector *math.Vector) *Vector {
	return &Vector{
		X: vector.X,
		Y: vector.Y,
	}
}
