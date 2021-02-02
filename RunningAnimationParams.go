package animatedledstrip

type runningAnimationParams struct {
	AnimationName  string                    `json:"animationName"`
	Colors         []*preparedColorContainer `json:"colors"`
	Id             string                    `json:"id"`
	Section        string                    `json:"section"`
	RunCount       int                       `json:"runCount"`
	IntParams      map[string]int            `json:"intParams"`
	DoubleParams   map[string]float64        `json:"doubleParams"`
	StringParams   map[string]string         `json:"stringParams"`
	LocationParams map[string]*location      `json:"locationParams"`
	DistanceParams map[string]*distance      `json:"distanceParams"`
	RotationParams map[string]*rotation      `json:"rotationParams"`
	EquationParams map[string]*equation      `json:"equationParams"`
	SourceParams   *animationToRunParams     `json:"sourceParams"`
}

func RunningAnimationParams(animationName string, colors []*preparedColorContainer, id string, section string,
	runCount int, intParams map[string]int, doubleParams map[string]float64, stringParams map[string]string,
	locationParams map[string]*location, distanceParams map[string]*distance, rotationParams map[string]*rotation,
	equationParams map[string]*equation, sourceParams *animationToRunParams) *runningAnimationParams {
	return &runningAnimationParams{
		AnimationName:  animationName,
		Colors:         colors,
		Id:             id,
		Section:        section,
		RunCount:       runCount,
		IntParams:      intParams,
		DoubleParams:   doubleParams,
		StringParams:   stringParams,
		LocationParams: locationParams,
		DistanceParams: distanceParams,
		RotationParams: rotationParams,
		EquationParams: equationParams,
		SourceParams:   sourceParams,
	}
}
