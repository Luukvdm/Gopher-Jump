package objects

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/luukvdm/jumper/game/controls"
	"github.com/luukvdm/jumper/game/objects/abstractions"
)

const (
	movementStep = 10
	maxSpeed     = 25
	jumpVelocity = 13
	mass         = 50
	playerWidth  = 50
	playerHeight = 50
)

type Player struct {
	*abstractions.AbstractObject
	isMovingRight bool
	isMovingLeft  bool
}

func NewPlayer(objId int, locX float64, locY float64) *Player {

	var playerObject = abstractions.NewAbstractObject(objId, abstractions.Vector{X: locX, Y: locY}, playerWidth, playerHeight, mass, false, false)
	var player = Player{
		AbstractObject: playerObject,
		isMovingLeft:   false,
		isMovingRight:  false,
	}

	player.AbstractObject.IAbstractObject = &player
	return &player
}

func (player *Player) Draw(ctx *cairo.Context) {
	ctx.SetSourceRGB(255, 0, 0)
	ctx.Rectangle(player.Location.X, player.Location.Y, player.Width, player.Height)
	ctx.Fill()
}

func (player *Player) Update(objects []*abstractions.AbstractObject) {
	oldLocation := player.Location
	bounced := false

	if player.Location.Y+player.Height > 720 {
		// Player is on the floor
		player.Location.Y = 720 - player.Height
		player.Jump()
		bounced = true
	} else if player.Location.Y < 0 {
		// Player is on the ceiling
		player.Location.Y = 0
		player.BounceVertical()
		bounced = true
	}

	if player.isMovingLeft && !player.isMovingRight {
		player.Location.X -= movementStep
	}
	if player.isMovingRight && !player.isMovingLeft {
		player.Location.X += movementStep
	}

	player.ApplyGravity()
	player.Velocity.Add(player.Acceleration)
	player.Velocity.Limit(maxSpeed)
	newLoc := abstractions.VectorAdd(player.Location, player.Velocity)

	// Check if moving from the old location to the new location collides with any other game object
	for _, gameObject := range objects {
		// Don't collide with yourself
		if gameObject.Id == player.Id {
			continue
		}

		if gameObject.Collides {
			// TODO
		} else if gameObject.IsPlatform {
			// Ignore platforms that aren't below the player
			if gameObject.Location.X > (newLoc.X+(player.Width)) || (gameObject.Width+gameObject.Location.X) < newLoc.X {
				// TODO continue the for loop in a cleaner way
				continue
			}

			// Ignore platforms that aren't below the player
			// Also don't do anything if the player has already bounced or is going up
			if !bounced && player.Velocity.Y > 0 && gameObject.Location.Y < (newLoc.Y+player.Height) {
				// Check if the player is currently on a platform
				if (gameObject.Location.Y) >= (oldLocation.Y+player.Height) && (gameObject.Location.Y) <= (newLoc.Y+player.Height) {
					player.Jump()
					player.Location.Y = gameObject.Location.Y - player.Height
					bounced = true
				}
			}
		}
	}

	player.Location.Add(player.Velocity)
	// Clear acceleration
	player.Acceleration.MultiplyByScalar(0)
}

func (player *Player) Jump() {
	player.BounceVertical()
	if player.Velocity.Y <= 0 {
		player.Velocity.Y = -jumpVelocity
	}
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
