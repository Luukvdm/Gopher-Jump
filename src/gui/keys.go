package gui

import (
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

type KeyHandler interface {
	ProcessKeyPress(keyId uint, state gdk.ModifierType) (ok bool)
	ProcessKeyRelease(keyId uint, state gdk.ModifierType)
}

func CreateKeyController(handler KeyHandler) *gtk.EventControllerKey {
	keyCtrl := gtk.NewEventControllerKey()
	setupKeyEventHandlers(handler, keyCtrl)
	return keyCtrl
}

func setupKeyEventHandlers(handler KeyHandler, controller *gtk.EventControllerKey) {
	controller.ConnectKeyPressed(func(keyVal, keycode uint, state gdk.ModifierType) (ok bool) {
		return handler.ProcessKeyPress(keycode, state)
	})
	controller.ConnectKeyReleased(func(keyVal, keycode uint, state gdk.ModifierType) {
		handler.ProcessKeyRelease(keycode, state)
	})
}
