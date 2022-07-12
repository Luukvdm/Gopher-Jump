package gui

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type gameCallbacks interface {
	KeyHandler
	Update(screenWidth, screenHeight int)
	DrawArea(ctx *cairo.Context, width, height int)
}

type GameWidget struct {
	game gameCallbacks
}

func NewGameWidget(parent GJumpWindow, g gameCallbacks) gtk.Widgetter {
	// Create GTK drawing area to draw the g on
	da := gtk.NewDrawingArea()
	da.SetSizeRequest(ScreenWidth, ScreenHeight)

	// Setup key controller
	parent.ConnectKeyEvents(g)

	// I'm not super happy with having to use this struct, might refactor later
	widget := GameWidget{g}
	da.AddTickCallback(widget.tick)
	da.SetDrawFunc(func(_ *gtk.DrawingArea, ctx *cairo.Context, width, height int) {
		// Having the first (DrawingArea) param in g.go means that gtk becomes a dependency for that package
		// And it doesn't even use it!
		g.DrawArea(ctx, width, height)
	})

	return da
}

func (widget GameWidget) tick(widGetter gtk.Widgetter, frameClock gdk.FrameClocker) bool {
	switch w := widGetter.(type) {
	case *gtk.DrawingArea:
		widget.game.Update(ScreenWidth, ScreenHeight)
		w.QueueDraw()
	}

	return true
}
