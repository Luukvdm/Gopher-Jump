package game

import (
	"github.com/gotk3/gotk3/cairo"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

type Window struct {
	Window *gtk.Window
	Canvas *gtk.DrawingArea
	Game Game
}

var prevFrame int64 = 0

func NewGameWindow() Window {
	var win, _ = gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	var game = NewGame()

	var da, _ = gtk.DrawingAreaNew()

	win.Add(da)
	win.SetTitle("Arrow keys")
	win.Connect("destroy", gtk.MainQuit)
	win.SetSizeRequest(1024, 720)
	win.SetResizable(false)
	win.ShowAll()

	win.AddTickCallback(tick)

	var gameWin = Window{win, da, game}
	gameWin.Canvas.SetSizeRequest(1024, 720)
	setupEventHandlers(gameWin)

	return gameWin
}

func tick(widget *gtk.Widget, clock *gdk.FrameClock) bool {
	widget.QueueDraw()

	return true
}

func setupEventHandlers(window Window) {
	// Draw
	window.Canvas.Connect("draw", func(da *gtk.DrawingArea, ctx *cairo.Context) {
		window.Game.Draw(ctx)
	})

	window.Window.Connect("key-press-event", func(win *gtk.Window, event *gdk.Event) {
		var keyEvent = &gdk.EventKey{Event: event}
		window.Game.ProcessKeyPress(keyEvent.KeyVal())
	})
}

