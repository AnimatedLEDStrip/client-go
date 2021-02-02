package animatedledstrip

type location struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

func Location(x float64, y float64, z float64) *location {
	return &location{X: x, Y: y, Z: z}
}
