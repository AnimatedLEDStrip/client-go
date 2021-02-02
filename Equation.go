package animatedledstrip

type equation struct {
	Coefficients []float64 `json:"coefficients"`
}

func Equation(coefficients []float64) *equation {
	return &equation{Coefficients: coefficients}
}
