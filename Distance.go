package animatedledstrip

type distance struct {
	DistanceType string  `json:"type"`
	X            float64 `json:"x"`
	Y            float64 `json:"y"`
	Z            float64 `json:"z"`
}

func AbsoluteDistance(x float64, y float64, z float64) *distance {
	return &distance{
		DistanceType: "AbsoluteDistance",
		X:            x,
		Y:            y,
		Z:            z,
	}
}

func PercentDistance(x float64, y float64, z float64) *distance {
	return &distance{
		DistanceType: "PercentDistance",
		X:            x,
		Y:            y,
		Z:            z,
	}
}
