package main

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/luukvdm/jumper/src"
	"github.com/luukvdm/jumper/src/gui"
	"os"
)

const appID = "com.github.luukvdm.jumper"

func main() {
	gtk.Init()

	app := gtk.NewApplication(appID, 0)
	game := src.NewGame()
	app.ConnectActivate(func() {
		gui.NewWindow(app, game)
	})

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}
