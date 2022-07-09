package gui

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Game interface {
	KeyHandler
	Update()
	DrawArea(ctx *cairo.Context, width, height int)
}

type GameWidget struct {
	game Game
	da   *gtk.DrawingArea
}

func NewGameWidget(parent JumperWindow, game Game) gtk.Widgetter {
	// Create GTK drawing area to draw the game on
	da := gtk.NewDrawingArea()
	da.SetSizeRequest(ScreenWidth, ScreenHeight)

	// I'm not super happy with having to use this struct, might refactor later
	widg := GameWidget{game, da}

	// Setup key controller
	parent.ConnectKeyEvents(game)

	// Setup draw loop
	da.AddTickCallback(widg.tick)
	da.SetDrawFunc(func(_ *gtk.DrawingArea, ctx *cairo.Context, width, height int) {
		// Having the first (DrawingArea) param in game.go means that gtk becomes a dependency for that package
		// And it doesn't even use it!
		game.DrawArea(ctx, width, height)
	})

	return da
}

func (widg GameWidget) tick(widGetter gtk.Widgetter, frameClock gdk.FrameClocker) bool {
	switch w := widGetter.(type) {
	case *gtk.DrawingArea:
		widg.game.Update()
		w.QueueDraw()
	}

	return true
}
