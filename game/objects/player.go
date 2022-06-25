package objects

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/luukvdm/jumper/game/controls"
	"github.com/luukvdm/jumper/game/objects/abstractions"
)

type Player struct {
	*abstractions.AbstractObject
	unitSize      float64
	isMovingRight bool
	isMovingLeft  bool
}

func NewPlayer(locX float64, locY float64) *Player {
	var playerSize float64 = 50

	var playerObject = abstractions.NewAbstractObject(abstractions.Vector{X: locX, Y: locY}, 50)
	var player = Player{playerObject, playerSize, false, false}

	player.AbstractObject.IAbstractObject = &player
	return &player
}

const (
	movementStep = 10
)

func (player *Player) Draw(ctx *cairo.Context) {
	var abstObj = player.AbstractObject
	ctx.SetSourceRGB(255, 0, 0)
	ctx.Rectangle(abstObj.Location.X, abstObj.Location.Y, player.unitSize, player.unitSize)
	ctx.Fill()
}

func (player *Player) Update() {
	if player.Location.Y < 0 {
		player.Location.Y = 0
		player.BounceVertical()
	}

	if player.Location.Y+player.unitSize > 720 {
		player.Location.Y = 720 - player.unitSize
		player.BounceVertical()
		// player.ApplyForce(abstractions.Vector{Y: -2})
	}

	if player.isMovingLeft && !player.isMovingRight {
		player.Location.X -= movementStep
	}
	if player.isMovingRight && !player.isMovingLeft {
		player.Location.X += movementStep
	}

	player.ApplyGravity()
	player.Velocity.Add(player.Acceleration)
	player.Location.Add(player.Velocity)
	// Clear acceleration
	// log.Printf("player acceleration: %f\n", player.Velocity)
	player.Acceleration.MultiplyByScalar(0)
}

func (player *Player) HandleKeyPress(keyId uint, state gdk.ModifierType) {
	switch keyId {
	case controls.KeyLeft:
		player.isMovingLeft = true
		break
		/*
			case controls.KeyUp:
				player.Location.Y -= movementStep
				break
		*/
	case controls.KeyRight:
		player.isMovingRight = true
		break
		/*
			case controls.KeyDown:
				player.Location.Y += movementStep
				break
		*/
	}
}
func (player *Player) HandleKeyRelease(keyId uint, state gdk.ModifierType) {
	switch keyId {
	case controls.KeyLeft:
		player.isMovingLeft = false
		break
	case controls.KeyRight:
		player.isMovingRight = false
		break
	}
}
