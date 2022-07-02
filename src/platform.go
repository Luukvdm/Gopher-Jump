package src

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/luukvdm/jumper/src/base_objects"
)

type Platform struct {
	*base_objects.AbstractObject
}

func NewPlatform(objId int, x float64, y float64) *Platform {
	var platformWidth float64 = 200
	var platformHeight float64 = 25

	platformObject := base_objects.NewAbstractObject(objId, base_objects.Vector{X: x, Y: y}, platformWidth, platformHeight, 10, true, false)
	platform := Platform{platformObject}

	platform.AbstractObject.IAbstractObject = &platform
	return &platform
}
func (platform *Platform) Draw(ctx *cairo.Context) {
	abstObj := platform.AbstractObject
	ctx.SetSourceRGB(0, 0, 0)
	ctx.Rectangle(abstObj.Location.X, abstObj.Location.Y, platform.Width, platform.Height)
	ctx.Fill()
	ctx.SetSourceRGB(255, 0, 0)
	ctx.Rectangle(abstObj.Location.X, abstObj.Location.Y, platform.Width, 5)
	ctx.Fill()
	ctx.SetSourceRGB(0, 0, 0)
}

func (platform *Platform) Update(objects []*base_objects.AbstractObject, offset base_objects.Vector) {
	platform.Location.Y += offset.Y
}
func (platform *Platform) HandleKeyPress(keyId uint, state gdk.ModifierType) {
}
func (platform *Platform) HandleKeyRelease(keyId uint, state gdk.ModifierType) {
}
