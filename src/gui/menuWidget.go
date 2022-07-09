package gui

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/luukvdm/jumper/src/game"
)

func NewMenuWidget(parent JumperWindow) gtk.Widgetter {
	// Create GTK drawing area to draw the g on
	box := gtk.NewBox(gtk.OrientationVertical, 25)

	startBtn := gtk.NewButton()
	startBtn.SetLabel("Start")
	startBtn.ConnectClicked(func() {
		g := game.NewGame()
		parent.LoadWidget(NewGameWidget(parent, g))
	})

	box.Append(startBtn)

	return box
}
