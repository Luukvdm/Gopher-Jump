package main

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/luukvdm/jumper/game"
)

func main() {
	gtk.Init(nil)

	var _ = game.NewGameWindow()

	gtk.Main()
}
