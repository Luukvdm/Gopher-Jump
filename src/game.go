package src

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/luukvdm/jumper/src/base_objects"
	"time"
)

type Game struct {
	FPS         int64
	player      *Player
	gameState   []*base_objects.AbstractObject
	currentTime int64
	accumulator int64
	dt          int64
	t           int64
	offset      base_objects.Vector
}

func NewGame() *Game {
	game := Game{FPS: 60}

	// Create some initial platforms
	// TODO create add game object func or something
	platform := NewPlatform(1, 50, 800)
	game.gameState = append(game.gameState, platform.AbstractObject)

	game.player = NewPlayer(0, 75, 500)
	game.currentTime = time.Now().UnixMilli()
	game.accumulator = 0
	game.dt = 1000 / game.FPS
	game.t = 0
	game.offset = base_objects.Vector{X: 0, Y: 0}

	return &game
}

func (g *Game) Update() {
	g.player.Update(g.gameState, g.offset)

	// Update the screen offset
	// if g.player.Location.Y <

	for _, gameObject := range g.gameState {
		gameObject.Update(g.gameState, g.offset)
	}
}

func (g *Game) Draw(ctx *cairo.Context) {
	g.player.Draw(ctx)
	for _, gameObject := range g.gameState {
		gameObject.Draw(ctx)
	}
}

func (g *Game) Tick(ctx *cairo.Context) {
	newTime := time.Now().UnixMilli()
	frameTime := newTime - g.currentTime

	g.currentTime = newTime
	g.accumulator += frameTime

	//fmt.Println("Accumulator: ", g.accumulator, " >= dt: ", g.dt)
	for g.accumulator >= g.dt {
		g.Update()
		g.t += g.dt
		g.accumulator -= g.dt
	}

	// alpha := g.accumulator / g.dt
	g.Draw(ctx)
}

func (g *Game) ProcessKeyPress(keyId uint, state gdk.ModifierType) (ok bool) {
	g.player.HandleKeyPress(keyId, state)
	/* for _, gameObject := range g.gameState {
		gameObject.HandleKeyPress(keyId, state)
	} */
	return true
}

func (g *Game) ProcessKeyRelease(keyId uint, state gdk.ModifierType) {
	g.player.HandleKeyRelease(keyId, state)
	/* for _, gameObject := range g.gameState {
		gameObject.HandleKeyRelease(keyId, state)
	} */
}
