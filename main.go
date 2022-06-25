package main

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/luukvdm/jumper/game"
	"os"
)

const appID = "com.github.luukvdm.jumper"

func main() {
	gtk.Init()

	app := gtk.NewApplication(appID, 0)
	app.ConnectActivate(func() { game.NewGameWindow(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}
