package abstractions

import (
	"math"
)

type Vector struct {
	X, Y float64
}

func (a *Vector) Add(b Vector) {
	a.X = a.X + b.X
	a.Y = a.Y + b.Y
}

func VectorAdd(a Vector, b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a *Vector) Sub(b Vector) {
	a.X = a.X - b.X
	a.Y = a.Y - b.Y
}

func VectorSub(a Vector, b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a *Vector) MultiplyByScalar(s float64) {
	a.X = a.X * s
	a.Y = a.Y * s
}

func VectorMultiplyByScalar(vec Vector, s float64) Vector {
	return Vector{
		X: vec.X * s,
		Y: vec.Y * s,
	}
}

func (a Vector) DivideByScalar(s float64) {
	a.X = a.X / s
	a.Y = a.Y / s
}

func VectorDivideByScalar(vec Vector, s float64) Vector {
	return Vector{
		X: vec.X / s,
		Y: vec.Y / s,
	}
}

func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y
}

func (a Vector) Length() float64 {
	return math.Sqrt(a.Dot(a))
}

func (a Vector) Normalize() {
	a.MultiplyByScalar(1. / a.Length())
}


