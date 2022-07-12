package game

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/luukvdm/gopher-jump/src/game/base_objects"
	"math/rand"
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

func (g *Game) UpdateState(screenWidth, screenHeight int) {
	screen := g.getScreenInfo(screenWidth, screenHeight)

	g.player.Update(g.gameState, screen)

	scrollBorder := screen.Top - (screen.Height / 2)
	if g.player.Location.Y > scrollBorder {
		// Scroll the screen up
		g.offsetTarget.Y = g.player.Location.Y - scrollBorder + screen.Bottom
	}

	if g.offsetTarget.Y > g.offset.Y {
		g.offset.Y += 5
	}

	// TODO get only platforms
	lastPlatform := g.gameState[len(g.gameState)-1]
	if lastPlatform.Location.Y < screen.Top-200 {
		g.spawnNewPlatform(lastPlatform, screenWidth, screenHeight)
	}

	for _, gameObject := range g.gameState {
		gameObject.Update(g.gameState, screen)
	}

	// Check if the player lost
	if g.player.Location.Y+g.player.Height < screen.Top {

		// win.LoadWidget(gui.NewMenuWidget(win))
	}
}
func (g *Game) DrawArea(ctx *cairo.Context, screenWidth, screenHeight int) {
	screen := g.getScreenInfo(screenWidth, screenHeight)

	// Move 0,0 point to bottom left, instead of top left
	ctx.Transform(&cairo.Matrix{
		Xx: 1,
		Yx: 0,
		Xy: 0,
		Yy: -1,
		X0: 0,
		Y0: screen.Top,
	})
	// :(
	// https://github.com/diamondburned/gotk4/blob/5e908130e58f7314673b10f0c96a0662fcc5a1fa/pkg/cairo/text.go#L39

	for _, gameObject := range g.gameState {
		gameObject.Draw(ctx, screen)
	}
	g.player.Draw(ctx, screen)

	// Because cairo_set_font_matrix isn't implemented yet we need to transform the entire context back
	// otherwise the text would be upside down
	ctx.Transform(&cairo.Matrix{Xx: 1, Yy: -1, Y0: g.offset.Y + screen.Height})
	ctx.SetSourceRGB(255, 0, 0)
	ctx.MoveTo(screen.Width/2, 100-10)
	ctx.SetFontSize(28)
	ctx.ShowText(strconv.Itoa(int(g.offset.Y) / 100))
}

func (g *Game) Update(screenWidth, screenHeight int) {
	newTime := time.Now().UnixMilli()
	frameTime := newTime - g.currentTime

	g.currentTime = newTime
	g.accumulator += frameTime

	for g.accumulator >= g.dt {
		g.UpdateState(screenWidth, screenHeight)
		g.t += g.dt
		g.accumulator -= g.dt
	}
}

func (g *Game) ProcessKeyPress(keyId uint, state gdk.ModifierType) (ok bool) {
	g.player.HandleKeyPress(keyId, state)
	return true
}

func (g *Game) ProcessKeyRelease(keyId uint, state gdk.ModifierType) {
	g.player.HandleKeyRelease(keyId, state)
}

func (g *Game) spawnNewPlatform(lastPlatform *base_objects.AbstractObject, screenWidth, screenHeight int) {
	// TODO make the platform width a variable for the platform object maybe
	platformWidth := 200
	max := screenWidth - platformWidth
	x := rand.Intn(max)

	var platform *Platform
	diffMultiplier := int(g.offset.Y / 1000)
	if rand.Intn(100) < diffMultiplier {
		platform = NewMovingPlatform(2, float64(x), lastPlatform.Location.Y+200)
	} else {
		platform = NewPlatform(2, float64(x), lastPlatform.Location.Y+200)
	}

	g.gameState = append(g.gameState, platform.AbstractObject)
}

func (g *Game) getScreenInfo(screenWidth, screenHeight int) base_objects.ScreenInfo {
	h := float64(screenHeight)
	w := float64(screenWidth)

	return base_objects.ScreenInfo{
		Left:   g.offset.X,
		Right:  g.offset.X + w,
		Bottom: g.offset.Y,
		Top:    g.offset.Y + h,
		Width:  w,
		Height: h,
	}
}
