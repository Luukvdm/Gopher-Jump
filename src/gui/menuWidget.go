package gui

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/diamondburned/gotk4/pkg/pango"
	"github.com/luukvdm/gopher-jump/src/game"
	"github.com/luukvdm/gopher-jump/src/media"
)

func NewMenuWidget(parent GJumpWindow) gtk.Widgetter {
	// Create GTK drawing area to draw the g on
	box := gtk.NewBox(gtk.OrientationVertical, 50)
	box.SetMarginTop(50)
	box.SetMarginStart(50)
	box.SetMarginEnd(50)

	gopher := media.GJumpPixelBuffs.Gopher
	// gopher.ScaleSimple(60, 75, gdkpixbuf.InterpBilinear)
	img := gtk.NewImageFromPixbuf(gopher)
	img.SetSizeRequest(120, 150)

	title := gtk.NewLabel("Gopher-Jump")
	title.SetWrap(true)
	title.SetWrapMode(pango.WrapWordChar)

	startBtn := gtk.NewButton()
	startBtn.SetLabel("Start")
	startBtn.ConnectClicked(func() {
		g := game.NewGame()
		parent.LoadWidget(NewGameWidget(parent, g))
	})

	stopBtn := gtk.NewButton()
	stopBtn.SetLabel("Exit")
	stopBtn.ConnectClicked(func() {
		parent.QuitApp()
	})

	box.Append(img)
	box.Append(title)
	box.Append(startBtn)
	box.Append(stopBtn)

	return box
}
