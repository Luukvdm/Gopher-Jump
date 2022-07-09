package base_objects

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

type IAbstractObject interface {
	Draw(ctx *cairo.Context, offset Vector)
	Update(objects []*AbstractObject, offset Vector, screenWidth, screenHeight float64)
	HandleKeyPress(keyId uint, state gdk.ModifierType)
	HandleKeyRelease(keyId uint, state gdk.ModifierType)
}

type AbstractObject struct {
	IAbstractObject
	Id           int
	Location     Vector
	Width        float64
	Height       float64
	Velocity     Vector
	Acceleration Vector
	Gravity      Vector
	Mass         float64
	IsPlatform   bool // Platform is an object the player (or some other entity) can stand on
	Collides     bool // Does this collide with other game objects (can they pass through it)
}

func NewAbstractObject(id int, location Vector, width float64, height float64, mass float64, isPlatform bool, collides bool) *AbstractObject {
	return &AbstractObject{
		Id:           id,
		Location:     location,
		Width:        width,
		Height:       height,
		Velocity:     Vector{},
		Acceleration: Vector{},
		Gravity:      Vector{Y: -0.2 * mass},
		Mass:         mass,
		IsPlatform:   isPlatform,
		Collides:     collides,
	}
}

func (obj *AbstractObject) OffsetLoc(offset Vector) Vector {
	return VectorAdd(obj.Location, offset)
}

func (obj *AbstractObject) ApplyForce(force Vector) {
	m := obj.Mass * 0.5
	f := VectorDivideByScalar(force, m)
	obj.Acceleration.Add(f)
}

func (obj *AbstractObject) ApplyGravity() {
	obj.ApplyForce(obj.Gravity)
}

func (obj *AbstractObject) ApplyFriction() {
	friction := Vector{X: obj.Velocity.X, Y: obj.Velocity.Y}
	friction.MultiplyByScalar(-1)
	friction.Normalize()
	friction.MultiplyByScalar(0.01)
	obj.ApplyForce(friction)
}

func (obj *AbstractObject) BounceVertical() {
	obj.Velocity.Y *= -1
}
