package objects

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/luukvdm/jumper/game/controls"
	"github.com/luukvdm/jumper/game/objects/abstractions"
)

type Player struct {
	*abstractions.AbstractObject
	unitSize float64
}

func NewPlayer() *Player {
	var playerSize float64 = 50

	var playerObject = abstractions.NewAbstractObject(abstractions.Vector{X: 5, Y: 5}, 10)
	var player = Player{playerObject, playerSize}

	player.AbstractObject.IAbstractObject = &player
	return &player
}

const (
	movementStep = 10
)

func (player *Player) Draw(ctx *cairo.Context)  {
	var abstObj = player.AbstractObject
	ctx.SetSourceRGB(0, 0, 0)
	ctx.Rectangle(abstObj.Location.X, abstObj.Location.Y, player.unitSize, player.unitSize)
	ctx.Fill()
}

func (player *Player) Update() {
	player.ApplyGravity()
	player.Velocity.Add(player.Acceleration)
	player.Location.Add(player.Velocity)
	if player.Location.Y < 0 {
		player.Location.Y = 0
	}

	if player.Location.Y + player.unitSize > 720 {
		player.Location.Y = 720 - player.unitSize
	}

	// player.Acceleration.MultiplyByScalar(0)
}

func (player *Player) HandleKeyPress(keyId uint) {
	switch keyId {
	case controls.KeyLeft:
		player.Location.X -= movementStep
		break
	case controls.KeyUp:
		player.Location.Y -= movementStep
		break
	case controls.KeyRight:
		player.Location.X += movementStep
		break
	case controls.KeyDown:
		player.Location.Y += movementStep
		break
	}
}