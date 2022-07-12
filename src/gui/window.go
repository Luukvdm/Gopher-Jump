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

type GJumpWindow interface {
	LoadWidget(widget gtk.Widgetter)
	ConnectKeyEvents(handler KeyHandler)
	QuitApp()
}

func NewWindow(app *gtk.Application) GJumpWindow {
	// Create GTK window
	appWin := gtk.NewApplicationWindow(app)
	appWin.SetTitle("Gopher-Jump")
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

func (window *Window) QuitApp() {
	window.Application().Quit()
}
