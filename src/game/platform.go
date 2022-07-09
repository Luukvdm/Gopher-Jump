package game

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	base_objects2 "github.com/luukvdm/jumper/src/game/base_objects"
)

type Platform struct {
	*base_objects2.AbstractObject
}

func NewPlatform(objId int, x float64, y float64) *Platform {
	var platformWidth float64 = 150
	var platformHeight float64 = 25

	platformObject := base_objects2.NewAbstractObject(objId, base_objects2.Vector{X: x, Y: y}, platformWidth, platformHeight, 10, true, false)
	platform := Platform{platformObject}

	platform.AbstractObject.IAbstractObject = &platform
	return &platform
}
func (platform *Platform) Draw(ctx *cairo.Context, offset base_objects2.Vector) {
	abstObj := platform.AbstractObject
	ctx.SetSourceRGB(0, 0, 0)
	ctx.Rectangle(abstObj.Location.X /*-offset.X*/, abstObj.Location.Y /*-offset.Y*/, platform.Width, platform.Height)
	ctx.Fill()
}

func (platform *Platform) Update(objects []*base_objects2.AbstractObject, offset base_objects2.Vector, screenWidth, screenHeight float64) {
	// platform.Location.Y += offsetTarget.Y
}
func (platform *Platform) HandleKeyPress(keyId uint, state gdk.ModifierType) {
}
func (platform *Platform) HandleKeyRelease(keyId uint, state gdk.ModifierType) {
}
