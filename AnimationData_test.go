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

import (
	"log"
	"testing"
)

func TestAnimationData(t *testing.T) {
	data := AnimationData()

	if data.Animation != "Color" {
		log.Print("Failed data.Animation check")
		t.Fail()
	} else if len(data.Colors) != 0 {
		log.Print("Failed data.Colors check")
		t.Fail()
	} else if data.Center != -1 {
		log.Print("Failed data.Center check")
		t.Fail()
	} else if data.Continuous != DEFAULT {
		log.Print("Failed data.Continuous check")
		t.Fail()
	} else if data.Delay != -1 {
		log.Print("Failed data.Delay check")
		t.Fail()
	} else if data.DelayMod != 1.0 {
		log.Print("Failed data.DelayMod check")
		t.Fail()
	} else if data.Direction != FORWARD {
		log.Print("Failed data.Direction check")
		t.Fail()
	} else if data.Distance != -1 {
		log.Print("Failed data.Distance check")
		t.Fail()
	} else if data.Id != "" {
		log.Print("Failed data.Id check")
		t.Fail()
	} else if data.Section != "" {
		log.Print("Failed data.Section check")
		t.Fail()
	} else if data.Spacing != -1 {
		log.Print("Failed data.Spacing check")
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
	if string(json) != `DATA:{"animation":"Meteor","center":50,"colors":[{"colors":[255,65280]},{"colors":[16711680]}],"continuous":false,"delay":10,"delayMod":1.5,"direction":"BACKWARD","distance":45,"id":"TEST","section":"SECT","spacing":5}` {
		t.Fail()
	}
}

func TestAnimationData_FromGoodJson(t *testing.T) {
	jsonStr := `DATA:{"animation":"Meteor","center":50,"colors":[{"colors":[255,65280]},{"colors":[16711680]}],"continuous":false,"delay":10,"delayMod":1.5,"direction":"BACKWARD","distance":45,"id":"TEST","section":"SECT","spacing":5}`

	data, _ := AnimationDataFromJson(jsonStr)

	if data.Animation != "Meteor" {
		log.Print("Failed data.Animation check")
		t.Fail()
	} else if len(data.Colors) != 2 {
		log.Print("Failed data.Colors check")
		t.Fail()
	} else if len(data.Colors[0].Colors) != 2 {
		log.Print("Failed data.Colors[0].Colors check")
		t.Fail()
	} else if len(data.Colors[1].Colors) != 1 {
		log.Print("Failed data.Colors[1].Colors check")
		t.Fail()
	} else if data.Colors[0].Colors[0] != 0xFF {
		log.Print("Failed data.Colors[0].Colors[0] check")
		t.Fail()
	} else if data.Colors[0].Colors[1] != 0xFF00 {
		log.Print("Failed data.Colors[0].Colors[1] check")
		t.Fail()
	} else if data.Colors[1].Colors[0] != 0xFF0000 {
		log.Print("Failed data.Colors[1].Colors[0] check")
		t.Fail()
	} else if data.Center != 50 {
		log.Print("Failed data.Center check")
		t.Fail()
	} else if data.Continuous != NONCONTINUOUS {
		log.Print("Failed data.Continuous check")
		t.Fail()
	} else if data.Delay != 10 {
		log.Print("Failed data.Delay check")
		t.Fail()
	} else if data.DelayMod != 1.5 {
		log.Print("Failed data.DelayMod check")
		t.Fail()
	} else if data.Direction != BACKWARD {
		log.Print("Failed data.Direction check")
		t.Fail()
	} else if data.Distance != 45 {
		log.Print("Failed data.Distance check")
		t.Fail()
	} else if data.Id != "TEST" {
		log.Print("Failed data.Id check")
		t.Fail()
	} else if data.Section != "SECT" {
		log.Print("Failed data.Section check")
		t.Fail()
	} else if data.Spacing != 5 {
		log.Print("Failed data.Spacing check")
		t.Fail()
	}
}

func TestAnimationData_FromBadJson(t *testing.T) {
	jsonStr := "{}"

	data, _ := AnimationDataFromJson(jsonStr)

	if data.Animation != "Color" {
		log.Print("Failed data.Animation check")
		t.Fail()
	} else if len(data.Colors) != 0 {
		log.Print("Failed data.Colors check")
		t.Fail()
	} else if data.Center != -1 {
		log.Print("Failed data.Center check")
		t.Fail()
	} else if data.Continuous != DEFAULT {
		log.Print("Failed data.Continuous check")
		t.Fail()
	} else if data.Delay != -1 {
		log.Print("Failed data.Delay check")
		t.Fail()
	} else if data.DelayMod != 1.0 {
		log.Print("Failed data.DelayMod check")
		t.Fail()
	} else if data.Direction != FORWARD {
		log.Print("Failed data.Direction check")
		t.Fail()
	} else if data.Distance != -1 {
		log.Print("Failed data.Distance check")
		t.Fail()
	} else if data.Id != "" {
		log.Print("Failed data.Id check")
		t.Fail()
	} else if data.Section != "" {
		log.Print("Failed data.Section check")
		t.Fail()
	} else if data.Spacing != -1 {
		log.Print("Failed data.Spacing check")
		t.Fail()
	}

}

func TestAnimationData_FromJson_Err(t *testing.T) {
	jsonStr := `{"animation":5}`

	_, err := AnimationDataFromJson(jsonStr)
	if err == nil {
		t.Fail()
	}
}

func TestAnimationData_ContinuousFromJson(t *testing.T) {
	// Tests for other continuous values

	jsonStr := `{"continuous":null}`
	data, _ := AnimationDataFromJson(jsonStr)
	if data.Continuous != DEFAULT {
		log.Print("Failed null -> DEFAULT")
		t.Fail()
	}

	jsonStr = `{"continuous":true}`
	data, _ = AnimationDataFromJson(jsonStr)
	if data.Continuous != CONTINUOUS {
		log.Print("Failed true -> CONTINUOUS")
		t.Fail()
	}

	jsonStr = `{"continuous":false}`
	data, _ = AnimationDataFromJson(jsonStr)
	if data.Continuous != NONCONTINUOUS {
		log.Print("Failed false -> NONCONTINUOUS")
		t.Fail()
	}

	jsonStr = `{"continuous":-1}`
	data, _ = AnimationDataFromJson(jsonStr)
	if data.Continuous != DEFAULT {
		log.Print("Failed -1 -> DEFAULT")
		t.Fail()
	}
}
