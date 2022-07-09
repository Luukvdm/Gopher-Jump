package gui

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type Window struct {
	*gtk.ApplicationWindow
}

const (
	ScreenWidth  = 720
	ScreenHeight = 1024
)

type JumperWindow interface {
	LoadWidget(widget gtk.Widgetter)
	ConnectKeyEvents(handler KeyHandler)
}

func NewWindow(app *gtk.Application) JumperWindow {
	// Create GTK window
	appWin := gtk.NewApplicationWindow(app)
	appWin.SetTitle("Jumper")
	appWin.SetSizeRequest(ScreenWidth, ScreenHeight)
	appWin.SetResizable(false)
	appWin.Show()

	// Create gameCallbacks instance
	win := Window{appWin}

	return &win
}

func (window *Window) LoadWidget(widget gtk.Widgetter) {
	window.SetChild(widget)
	window.SetDefaultWidget(widget)
	window.SetFocus(widget)
}

func (window *Window) ConnectKeyEvents(handler KeyHandler) {
	c := CreateKeyController(handler)
	window.AddController(c)
}
