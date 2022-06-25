package objects

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/luukvdm/jumper/game/objects/abstractions"
)

type Platform struct {
	*abstractions.AbstractObject
	width float64
}

func NewPlatform(x float64, y float64) *Platform {
	var platformWidth float64 = 200

	var platformObject = abstractions.NewAbstractObject(abstractions.Vector{X: x, Y: y}, 10)
	var platform = Platform{platformObject, platformWidth}

	platform.AbstractObject.IAbstractObject = &platform
	return &platform
}
func (platform *Platform) Draw(ctx *cairo.Context) {
	var abstObj = platform.AbstractObject
	ctx.SetSourceRGB(0, 0, 0)
	ctx.Rectangle(abstObj.Location.X, abstObj.Location.Y, platform.width, 25)
	ctx.Fill()
}

func (platform *Platform) Update() {
}
func (platform *Platform) HandleKeyPress(keyId uint, state gdk.ModifierType) {
}
func (platform *Platform) HandleKeyRelease(keyId uint, state gdk.ModifierType) {
}
