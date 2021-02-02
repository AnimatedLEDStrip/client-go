package animatedledstrip

type rotation struct {
	RotationType  string   `json:"type"`
	XRotation     float64  `json:"xRotation"`
	YRotation     float64  `json:"yRotation"`
	ZRotation     float64  `json:"zRotation"`
	RotationOrder []string `json:"rotationOrder"`
}

func DegreesRotation(xRotation float64, yRotation float64, zRotation float64, rotationOrder []string) *rotation {
	return &rotation{
		RotationType:  "DegreesRotation",
		XRotation:     xRotation,
		YRotation:     yRotation,
		ZRotation:     zRotation,
		RotationOrder: rotationOrder,
	}
}

func RadiansRotation(xRotation float64, yRotation float64, zRotation float64, rotationOrder []string) *rotation {
	return &rotation{
		RotationType:  "RadiansRotation",
		XRotation:     xRotation,
		YRotation:     yRotation,
		ZRotation:     zRotation,
		RotationOrder: rotationOrder,
	}
}
