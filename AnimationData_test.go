/*
 *  Copyright (c) 2019-2020 AnimatedLEDStrip
 *
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
 *  of this software and associated documentation files (the "Software"), to deal
 *  in the Software without restriction, including without limitation the rights
 *  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 *  copies of the Software, and to permit persons to whom the Software is
 *  furnished to do so, subject to the following conditions:
 *
 *  The above copyright notice and this permission notice shall be included in
 *  all copies or substantial portions of the Software.
 *
 *  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 *  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 *  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 *  THE SOFTWARE.
 */

package animatedledstrip

import "testing"

func TestAnimationData(t *testing.T) {
	data := AnimationData()

	if data.Animation != "Color" {
		t.Fail()
	} else if len(data.Colors) != 0 {
		t.Fail()
	} else if data.Center != -1 {
		t.Fail()
	} else if data.Continuous != DEFAULT {
		t.Fail()
	} else if data.Delay != -1 {
		t.Fail()
	} else if data.DelayMod != 1.0 {
		t.Fail()
	} else if data.Direction != FORWARD {
		t.Fail()
	} else if data.Distance != -1 {
		t.Fail()
	} else if data.Id != "" {
		t.Fail()
	} else if data.Section != "" {
		t.Fail()
	} else if data.Spacing != -1 {
		t.Fail()
	}
}

func TestAnimationData_SetAnimation(t *testing.T) {
	data := AnimationData()
	data.SetAnimation("Bounce")

	if data.Animation != "Bounce" {
		t.Fail()
	}
}

func TestAnimationData_AddColor(t *testing.T) {
	cc := ColorContainer{}
	cc.AddColor(0xFF)

	data := AnimationData()
	data.AddColor(&cc)

	if len(data.Colors) != 1 {
		t.Fail()
	}
}

func TestAnimationData_SetCenter(t *testing.T) {
	data := AnimationData()
	data.SetCenter(50)

	if data.Center != 50 {
		t.Fail()
	}
}

func TestAnimationData_SetContinuous(t *testing.T) {
	data := AnimationData()
	data.SetContinuous(CONTINUOUS)

	if data.Continuous != CONTINUOUS {
		t.Fail()
	}
}

func TestAnimationData_SetDelay(t *testing.T) {
	data := AnimationData()
	data.SetDelay(200)

	if data.Delay != 200 {
		t.Fail()
	}
}

func TestAnimationData_SetDelayMod(t *testing.T) {
	data := AnimationData()
	data.SetDelayMod(2.0)

	if data.DelayMod != 2.0 {
		t.Fail()
	}
}

func TestAnimationData_SetDirection(t *testing.T) {
	data := AnimationData()
	data.SetDirection(BACKWARD)

	if data.Direction != BACKWARD {
		t.Fail()
	}
}

func TestAnimationData_SetDistance(t *testing.T) {
	data := AnimationData()
	data.SetDistance(35)

	if data.Distance != 35 {
		t.Fail()
	}
}

func TestAnimationData_SetID(t *testing.T) {
	data := AnimationData()
	data.SetID("TEST")

	if data.Id != "TEST" {
		t.Fail()
	}
}

func TestAnimationData_SetSection(t *testing.T) {
	data := AnimationData()
	data.SetSection("SECT")

	if data.Section != "SECT" {
		t.Fail()
	}
}

func TestAnimationData_SetSpacing(t *testing.T) {
	data := AnimationData()
	data.SetSpacing(4)

	if data.Spacing != 4 {
		t.Fail()
	}
}

func TestAnimationData_Json(t *testing.T) {
	data := AnimationData()
	data.SetAnimation("Meteor")
	data.SetCenter(50)
	data.SetContinuous(NONCONTINUOUS)
	data.SetDelay(10)
	data.SetDelayMod(1.5)
	data.SetDirection(BACKWARD)
	data.SetDistance(45)
	data.SetID("TEST")
	data.SetSection("SECT")
	data.SetSpacing(5)

	cc := ColorContainer{}
	cc.AddColor(0xFF).AddColor(0xFF00)
	cc2 := ColorContainer{}
	cc2.AddColor(0xFF0000)
	data.AddColor(&cc)
	data.AddColor(&cc2)

	json := data.Json()
	if json != `DATA:{"animation":"Meteor","colors":[{"colors":[255,65280]},{"colors":[16711680]}],"center":50,"continuous":false,"baseDelay":10,"delayMod":1.500000,"direction":"BACKWARD","distance":45,"id":"TEST","section":"SECT","spacing":5}` {
		t.Fail()
	}
}

func TestAnimationData_FromGoodJson(t *testing.T) {
	// Good JSON test

	jsonStr := `DATA:{"animation":"Meteor","colors":[{"colors":[255,65280]},{"colors":[16711680]}],"center":50,"continuous":false,"baseDelay":10,"delayMod":1.500000,"direction":"BACKWARD","distance":45,"id":"TEST","section":"SECT","spacing":5}`

	data := AnimationDataFromJson(jsonStr)

	if data.Animation != "Meteor" {
		t.Fail()
	} else if len(data.Colors) != 2 {
		t.Fail()
	} else if len(data.Colors[0].Colors) != 2 {
		t.Fail()
	} else if len(data.Colors[1].Colors) != 1 {
		t.Fail()
	} else if data.Colors[0].Colors[0] != 0xFF {
		t.Fail()
	} else if data.Colors[0].Colors[1] != 0xFF00 {
		t.Fail()
	} else if data.Colors[1].Colors[0] != 0xFF0000 {
		t.Fail()
	} else if data.Center != 50 {
		t.Fail()
	} else if data.Continuous != NONCONTINUOUS {
		t.Fail()
	} else if data.Delay != 10 {
		t.Fail()
	} else if data.DelayMod != 1.5 {
		t.Fail()
	} else if data.Direction != BACKWARD {
		t.Fail()
	} else if data.Distance != 45 {
		t.Fail()
	} else if data.Id != "TEST" {
		t.Fail()
	} else if data.Section != "SECT" {
		t.Fail()
	} else if data.Spacing != 5 {
		t.Fail()
	}
}

func TestAnimationData_FromBadJson(t *testing.T) {
	// Bad JSON test

	jsonStr := "{}"

	data := AnimationDataFromJson(jsonStr)

	if data.Animation != "Color" {
		t.Fail()
	} else if len(data.Colors) != 0 {
		t.Fail()
	} else if data.Center != -1 {
		t.Fail()
	} else if data.Continuous != DEFAULT {
		t.Fail()
	} else if data.Delay != -1 {
		t.Fail()
	} else if data.DelayMod != 1.0 {
		t.Fail()
	} else if data.Direction != FORWARD {
		t.Fail()
	} else if data.Distance != -1 {
		t.Fail()
	} else if data.Id != "" {
		t.Fail()
	} else if data.Section != "" {
		t.Fail()
	} else if data.Spacing != -1 {
		t.Fail()
	}

}

func TestAnimationData_ContinuousFromJson(t *testing.T) {
	// Tests for other continuous values

	jsonStr := `{"continuous":null}`

	data := AnimationDataFromJson(jsonStr)

	if data.Continuous != DEFAULT {
		t.Fail()
	}

	jsonStr = `{"continuous":true}`

	data = AnimationDataFromJson(jsonStr)

	if data.Continuous != CONTINUOUS {
		t.Fail()
	}

	jsonStr = `{"continuous":false}`

	data = AnimationDataFromJson(jsonStr)

	if data.Continuous != NONCONTINUOUS {
		t.Fail()
	}

	jsonStr = `{"continuous":-1}`

	data = AnimationDataFromJson(jsonStr)

	if data.Continuous != DEFAULT {
		t.Fail()
	}
}
