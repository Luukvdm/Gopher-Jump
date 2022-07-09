package gui

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Game interface {
	KeyHandler
	Tick(ctx *cairo.Context)
}

func NewGameWidget(parent JumperWindow, game Game) gtk.Widgetter {
	// Create GTK drawing area to draw the game on
	da := gtk.NewDrawingArea()
	da.AddTickCallback(tick)
	da.SetSizeRequest(ScreenWidth, ScreenHeight)

	// Setup key controller
	parent.ConnectKeyEvents(game)

	// Setup draw loop
	da.SetDrawFunc(func(drawingArea *gtk.DrawingArea, ctx *cairo.Context, width, height int) {
		game.Tick(ctx)
	})

	return da
}

func tick(widGetter gtk.Widgetter, frameClock gdk.FrameClocker) bool {
	switch w := widGetter.(type) {
	case *gtk.DrawingArea:
		w.QueueDraw()
	}

	return true
}
