package game

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Window struct {
	Window *gtk.ApplicationWindow
	Canvas *gtk.DrawingArea
	Game   *Game
}

func NewGameWindow(app *gtk.Application) *Window {
	// Create GTK drawing area to draw the game on
	var da = gtk.NewDrawingArea()
	da.AddTickCallback(tick)

	// Create GTK window
	var win = gtk.NewApplicationWindow(app)
	win.SetChild(da)
	win.SetTitle("Jumper")
	win.SetSizeRequest(1024, 720)
	win.SetResizable(false)
	win.AddTickCallback(tick)

	// Create game instance
	var game = NewGame()
	var gameWin = Window{win, da, game}
	gameWin.Canvas.SetSizeRequest(1024, 720)

	// Setup key controller
	keyCtrl := gtk.NewEventControllerKey()
	setupKeyEventHandlers(gameWin, keyCtrl)
	gameWin.Window.AddController(keyCtrl)
	// Setup other event handlers
	setupEventHandlers(gameWin)

	gameWin.Window.Show()

	return &gameWin
}

func tick(widgetter gtk.Widgetter, frameClock gdk.FrameClocker) bool {
	switch w := widgetter.(type) {
	case *gtk.DrawingArea:
		w.QueueDraw()
	}

	return true
}

func setupEventHandlers(window Window) {
	// Draw
	window.Canvas.SetDrawFunc(func(da *gtk.DrawingArea, ctx *cairo.Context, width, height int) {
		window.Game.Draw(ctx)
	})
}

func setupKeyEventHandlers(window Window, controller *gtk.EventControllerKey) {
	controller.ConnectKeyPressed(func(keyval, keycode uint, state gdk.ModifierType) (ok bool) {
		return window.Game.ProcessKeyPress(keycode, state)
	})
	controller.ConnectKeyReleased(func(keyval, keycode uint, state gdk.ModifierType) {
		window.Game.ProcessKeyRelease(keycode, state)
	})
}
