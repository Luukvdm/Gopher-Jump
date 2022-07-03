package src

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/luukvdm/jumper/src/base_objects"
	"github.com/luukvdm/jumper/src/gui"
	"strconv"
	"time"
)

type Game struct {
	FPS          int64
	player       *Player
	gameState    []*base_objects.AbstractObject
	currentTime  int64
	accumulator  int64
	dt           int64
	t            int64
	offsetTarget base_objects.Vector
	offset       base_objects.Vector
}

func NewGame() *Game {
	game := Game{FPS: 60}

	// Create some initial platforms
	// TODO create add game object func or something
	platformA := NewPlatform(1, 50, 200)
	platformB := NewPlatform(2, 400, 400)
	game.gameState = append(game.gameState, platformA.AbstractObject)
	game.gameState = append(game.gameState, platformB.AbstractObject)

	game.player = NewPlayer(0, 75, 500)
	game.currentTime = time.Now().UnixMilli()
	game.accumulator = 0
	game.dt = 1000 / game.FPS
	game.t = 0

	return &game
}

func (g *Game) Update() {
	g.player.Update(g.gameState, g.offset)

	scrollBorder := (gui.ScreenHeight + g.offset.Y) * 0.5
	// log.Printf("border: %f player: %f", scrollBorder, g.player.Location.Y)
	if g.player.Location.Y > scrollBorder {
		// Scroll the screen up
		g.offsetTarget.Y = g.player.Location.Y - scrollBorder
		// log.Printf("player height: %f scroll border: %f scroll with: %f", g.player.Location.Y, scrollBorder, (scrollBorder - g.player.Location.Y))
	}

	if g.offsetTarget.Y > g.offset.Y {
		g.offset.Y += 5
	}

	// TODO filter out platforms
	/* lastPlatform := g.gameState[len(g.gameState)-1]
	if lastPlatform.Location.Y-g.offset.Y < 400 {
		platform := NewPlatform(2, 400, 800+g.offset.Y)
		g.gameState = append(g.gameState, platform.AbstractObject)
	} */

	for _, gameObject := range g.gameState {
		gameObject.Update(g.gameState, g.offset)
	}
}

func (g *Game) Draw(ctx *cairo.Context) {
	// Move 0,0 point to bottom left, instead of top left
	ctx.Transform(&cairo.Matrix{
		Xx: 1,
		Yx: 0,
		Xy: 0,
		Yy: -1,
		X0: 0,
		Y0: gui.ScreenHeight + g.offset.Y,
	})
	// :(
	// https://github.com/diamondburned/gotk4/blob/5e908130e58f7314673b10f0c96a0662fcc5a1fa/pkg/cairo/text.go#L39

	for i := 0; i < gui.ScreenWidth; i += 50 {
		ctx.MoveTo(float64(i), 10)
		ctx.ShowText(strconv.Itoa(i))
	}
	for i := 0; i < gui.ScreenHeight; i += 50 {
		ctx.MoveTo(5, float64(i))
		ctx.ShowText(strconv.Itoa(i))
	}
	g.player.Draw(ctx, g.offset)
	for _, gameObject := range g.gameState {
		gameObject.Draw(ctx, g.offset)
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
