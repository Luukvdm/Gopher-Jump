package game

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/luukvdm/jumper/src/game/base_objects"
)

const (
	bounceMargin     = 25
	maxPlatformSpeed = 25
)

type Platform struct {
	*base_objects.AbstractObject
	color base_objects.RGBA
}

func NewPlatform(objId int, x, y float64) *Platform {
	return newPlatform(objId, x, y, base_objects.NewRGBA(0, 0, 0, 255))
}

func NewMovingPlatform(objId int, x, y float64) *Platform {
	p := newPlatform(objId, x, y, base_objects.NewRGBA(0, 0, 255, 255))
	// p.Velocity.X = 3
	p.Acceleration.X = 3
	return p
}

func newPlatform(objId int, x float64, y float64, color base_objects.RGBA) *Platform {
	var platformWidth float64 = 150
	var platformHeight float64 = 25

	platformObject := base_objects.NewAbstractObject(objId, base_objects.Vector{X: x, Y: y}, platformWidth, platformHeight, 10, true, false)
	platform := Platform{platformObject, color}

	platform.AbstractObject.IAbstractObject = &platform
	return &platform
}
func (platform *Platform) Draw(ctx *cairo.Context, offset base_objects.Vector) {
	abstObj := platform.AbstractObject
	ctx.SetSourceRGB(0, 0, 0)
	ctx.SetSourceRGBA(platform.color.R, platform.color.G, platform.color.B, platform.color.A)
	ctx.Rectangle(abstObj.Location.X /*-offset.X*/, abstObj.Location.Y /*-offset.Y*/, platform.Width, platform.Height)
	ctx.Fill()
}

func (platform *Platform) Update(objects []*base_objects.AbstractObject, offset base_objects.Vector, screenWidth, screenHeight float64) {
	platform.Velocity.Add(platform.Acceleration)
	platform.Velocity.Limit(maxPlatformSpeed)
	platform.Location.Add(platform.Velocity)
	platform.Acceleration.MultiplyByScalar(0)

	if platform.Location.X+platform.Width > screenWidth-bounceMargin || platform.Location.X < bounceMargin {
		platform.BounceHorizontal()
	}
}
func (platform *Platform) HandleKeyPress(keyId uint, state gdk.ModifierType) {
}
func (platform *Platform) HandleKeyRelease(keyId uint, state gdk.ModifierType) {
}
