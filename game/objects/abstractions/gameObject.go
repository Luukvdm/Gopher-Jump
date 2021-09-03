package abstractions

import (
	"github.com/gotk3/gotk3/cairo"
)

type IAbstractObject interface {
	Draw(ctx *cairo.Context)
	Update()
	HandleKeyPress(keyId uint)
}

type AbstractObject struct {
	IAbstractObject
	Location Vector
	Velocity Vector
	Acceleration Vector
	Gravity Vector
	Mass float64
}

func NewAbstractObject(location Vector, mass float64) *AbstractObject {
	return &AbstractObject{
		Location: location,
		Velocity: Vector{},
		Acceleration: Vector{},
		Gravity: Vector{Y: 0.01*mass},
		Mass: mass,
	}
}

func (obj *AbstractObject) ApplyForce(force Vector) {
	var f = VectorDivideByScalar(force, obj.Mass)
	obj.Acceleration.Add(f)
}

func (obj *AbstractObject) ApplyGravity() {
	obj.ApplyForce(obj.Gravity)
}

func (obj *AbstractObject) BounceVertical() {
	obj.Velocity.Y *= -1
}