package media

import (
	_ "embed"
	"github.com/diamondburned/gotk4/pkg/gdkpixbuf/v2"
	"log"
)

//go:embed resources/gopher.png
var GopherPNG []byte

type PixelBuffs struct {
	Gopher *gdkpixbuf.Pixbuf
}

var JumperPixelBuffs PixelBuffs

func init() {
	gopher, err := LoadPNG(GopherPNG)
	if err != nil {
		log.Fatalln("failed to load gopher.png:", err)
	}

	JumperPixelBuffs = PixelBuffs{
		gopher,
	}
}
