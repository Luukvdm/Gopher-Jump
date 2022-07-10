package base_objects

type RGBA struct {
	R, G, B, A float64
}

func NewRGBA(R, G, B, A float64) RGBA {
	return RGBA{R: R, G: G, B: B, A: A}
}
