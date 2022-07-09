package main

import (
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"github.com/luukvdm/jumper/src/gui"
	"os"
)

const appID = "com.github.luukvdm.jumper"

func main() {
	gtk.Init()

	app := gtk.NewApplication(appID, 0)
	app.ConnectActivate(func() {
		win := gui.NewWindow(app)
		win.LoadWidget(gui.NewMenuWidget(win))
	})

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}
