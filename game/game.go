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
	gameState   []abstractions.IAbstractObject
	currentTime int64
	accumulator int64
	dt          int64
	t           int64
}

func NewGame() *Game {
	var game = Game{FPS: 60}

	// Create player
	// TODO create add game object func or something
	var player = objects.NewPlayer(75, 15)
	var platform = objects.NewPlatform(50, 500)

	game.gameState = append(game.gameState, player)
	game.gameState = append(game.gameState, platform)

	game.currentTime = time.Now().UnixMilli()
	game.accumulator = 0
	game.dt = 1000 / game.FPS
	game.t = 0

	return &game
}

func (g *Game) Update() {
	for _, gameObject := range g.gameState {
		gameObject.Update()
	}
}
func (g *Game) Draw(ctx *cairo.Context) {
	var newTime = time.Now().UnixMilli()
	var frameTime = newTime - g.currentTime

	g.currentTime = newTime
	g.accumulator += frameTime

	//fmt.Println("Accumulator: ", g.accumulator, " >= dt: ", g.dt)
	for g.accumulator >= g.dt {
		g.Update()
		g.t += g.dt
		g.accumulator -= g.dt
	}

	// var alpha = g.accumulator / g.dt

	for _, gameObject := range g.gameState {
		gameObject.Draw(ctx)
	}
}

func (g *Game) ProcessKeyPress(keyId uint, state gdk.ModifierType) (ok bool) {
	for _, gameObject := range g.gameState {
		gameObject.HandleKeyPress(keyId, state)
	}
	return true
}
