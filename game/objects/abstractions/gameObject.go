package abstractions

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

type IAbstractObject interface {
	Draw(ctx *cairo.Context)
	Update()
	HandleKeyPress(keyId uint, state gdk.ModifierType)
}

type AbstractObject struct {
	IAbstractObject
	Location     Vector
	Velocity     Vector
	Acceleration Vector
	Gravity      Vector
	Mass         float64
}

func NewAbstractObject(location Vector, mass float64) *AbstractObject {
	return &AbstractObject{
		Location:     location,
		Velocity:     Vector{},
		Acceleration: Vector{},
		Gravity:      Vector{Y: 0.5 * mass},
		Mass:         mass,
	}
}

func (obj *AbstractObject) ApplyForce(force Vector) {
	var m = obj.Mass * 0.5
	var f = VectorDivideByScalar(force, m)
	obj.Acceleration.Add(f)
}

func (obj *AbstractObject) ApplyGravity() {
	obj.ApplyForce(obj.Gravity)
}

func (obj *AbstractObject) BounceVertical() {
	obj.Velocity.Y *= -1
}
