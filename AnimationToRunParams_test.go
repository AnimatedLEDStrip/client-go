package animatedledstrip

import (
	"encoding/json"
	"testing"
)

func TestAnimationToRunParams(t *testing.T) {
	//a := AnimationToRunParams("", []*colorContainer{}, "", "", 0, map[string]int{}, map[string]float64{}, map[string]string{}, map[string]*location{}, map[string]*distance{}, map[string]*rotation{"x": RadiansRotation(5.0, 4.0, 3.0, []string{})}, map[string]*equation{})
	//println(a)
	//r, e := json.Marshal(a)
	//println(string(r))
	//println(e)

	c := ALSHttpClient("10.0.0.91")
	var p []int
	p = append(p, 1)
	p = append(p, 2)
	//s := Section("section9", p, "")
	//ns, _ := c.CreateNewSection(s)
	//println(ns.Name)
	//n, _ := c.GetSections()
	//println(n[7].Name)
	var col []*colorContainer
	var colInts []int
	colInts = append(colInts, 0xFF)
	col = append(col, ColorContainer(colInts))
	a := AnimationToRunParams("ripple", col, "", "", 5, map[string]int{}, map[string]float64{}, map[string]string{}, map[string]*location{}, map[string]*distance{}, map[string]*rotation{}, map[string]*equation{})
	b, e := json.Marshal(a)
	if e != nil {
		println(e.Error())
	}
	println(string(b))

	n, err := c.StartAnimation(a)
	println(n)
	if err != nil {
		println(err.Error())
	}
	//i, _:= c.GetAnimationInfo("fireworks")
	//println(location(*(i.DistanceParams[0].Default)))
	//p := s.GetRunningAnimationParams("6898263")
	//println(p.AnimationName)
	//println(p.Colors[0].Colors[0])
	//println(p.RotationParams["rotation"].RotationType)
}
