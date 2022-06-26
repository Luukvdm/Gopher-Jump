package game

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/luukvdm/jumper/game/objects"
	"github.com/luukvdm/jumper/game/objects/abstractions"
	"time"
)

type Game struct {
	FPS         int64
	gameState   []*abstractions.AbstractObject
	currentTime int64
	accumulator int64
	dt          int64
	t           int64
}

func NewGame() *Game {
	game := Game{FPS: 60}

	// Create player
	// TODO create add game object func or something
	player := objects.NewPlayer(0, 75, 15)
	platform := objects.NewPlatform(1, 50, 500)

	game.gameState = append(game.gameState, player.AbstractObject)
	game.gameState = append(game.gameState, platform.AbstractObject)

	game.currentTime = time.Now().UnixMilli()
	game.accumulator = 0
	game.dt = 1000 / game.FPS
	game.t = 0

	return &game
}

func (g *Game) Update() {
	for _, gameObject := range g.gameState {
		gameObject.Update(g.gameState)
	}
}

func (g *Game) Draw(ctx *cairo.Context) {
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
	for _, gameObject := range g.gameState {
		gameObject.HandleKeyPress(keyId, state)
	}
	return true
}

func (g *Game) ProcessKeyRelease(keyId uint, state gdk.ModifierType) {
	for _, gameObject := range g.gameState {
		gameObject.HandleKeyRelease(keyId, state)
	}
}
