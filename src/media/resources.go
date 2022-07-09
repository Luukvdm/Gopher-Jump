package media

import (
	_ "embed"
	"fmt"
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

func LoadPNG(data []byte) (*gdkpixbuf.Pixbuf, error) {
	l, err := gdkpixbuf.NewPixbufLoaderWithType("png")
	if err != nil {
		return nil, fmt.Errorf("NewLoaderWithType png: %w", err)
	}
	defer l.Close()

	if err := l.Write(data); err != nil {
		return nil, fmt.Errorf("PixbufLoader.Write: %w", err)
	}

	if err := l.Close(); err != nil {
		return nil, fmt.Errorf("PixbufLoader.Close: %w", err)
	}

	return l.Pixbuf(), nil
}
