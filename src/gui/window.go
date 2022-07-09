package gui

import (
	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Window struct {
	Window *gtk.ApplicationWindow
	Canvas *gtk.DrawingArea
	Game   IGame
}

type IGame interface {
	Tick(ctx *cairo.Context)
	ProcessKeyPress(keyId uint, state gdk.ModifierType) (ok bool)
	ProcessKeyRelease(keyId uint, state gdk.ModifierType)
}

const (
	ScreenWidth  = 720
	ScreenHeight = 1024
)

func NewWindow(app *gtk.Application, game IGame) *Window {
	// Create GTK drawing area to draw the game on
	da := gtk.NewDrawingArea()
	da.AddTickCallback(tick)

	// Create GTK window
	appWin := gtk.NewApplicationWindow(app)
	appWin.SetChild(da)
	appWin.SetTitle("Jumper")
	appWin.SetSizeRequest(ScreenWidth, ScreenHeight)
	appWin.SetResizable(false)
	// appWin.AddTickCallback(tick)

	// Create game instance
	win := Window{appWin, da, game}
	win.Canvas.SetSizeRequest(ScreenWidth, ScreenHeight)

	// Setup key controller
	keyCtrl := gtk.NewEventControllerKey()
	setupKeyEventHandlers(win, keyCtrl)
	win.Window.AddController(keyCtrl)
	// Setup other event handlers
	setupEventHandlers(win)

	win.Window.Show()

	return &win
}

func tick(widGetter gtk.Widgetter, frameClock gdk.FrameClocker) bool {
	switch w := widGetter.(type) {
	case *gtk.DrawingArea:
		w.QueueDraw()
	}

	return true
}

func setupEventHandlers(window Window) {
	window.Canvas.SetDrawFunc(func(da *gtk.DrawingArea, ctx *cairo.Context, width, height int) {
		window.Game.Tick(ctx)
	})
}

func setupKeyEventHandlers(window Window, controller *gtk.EventControllerKey) {
	controller.ConnectKeyPressed(func(keyVal, keycode uint, state gdk.ModifierType) (ok bool) {
		return window.Game.ProcessKeyPress(keycode, state)
	})
	controller.ConnectKeyReleased(func(keyVal, keycode uint, state gdk.ModifierType) {
		window.Game.ProcessKeyRelease(keycode, state)
	})
}
