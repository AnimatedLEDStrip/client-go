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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAnimationData(t *testing.T) {
	data := AnimationData()

	assert.Equal(t, "Color", data.Animation)
	assert.Len(t, data.Colors, 0)
	assert.Equal(t, -1, data.Center)
	assert.Equal(t, DEFAULT, data.Continuous)
	assert.Equal(t, -1, data.Delay)
	assert.Equal(t, 1.0, data.DelayMod)
	assert.Equal(t, FORWARD, data.Direction)
	assert.Equal(t, -1, data.Distance)
	assert.Equal(t, "", data.Id)
	assert.Equal(t, "", data.Section)
	assert.Equal(t, -1, data.Spacing)
}

func TestAnimationData_SetAnimation(t *testing.T) {
	data := AnimationData()
	data.SetAnimation("Bounce")

	assert.Equal(t, "Bounce", data.Animation)
}

//func TestAnimationData_AddColor(t *testing.T) {
//	cc := ColorContainer()
//	cc.AddColor(0xFF)
//
//	data := AnimationData()
//	data.AddColor(&cc)
//
//	assert.Len(t, data.Colors, 1)
//	assert.Equal(t, 0xFF, data.Colors[0].Colors[0])
//}

func TestAnimationData_SetCenter(t *testing.T) {
	data := AnimationData()
	data.SetCenter(50)

	assert.Equal(t, 50, data.Center)
}

func TestAnimationData_SetContinuous(t *testing.T) {
	data := AnimationData()
	data.SetContinuous(CONTINUOUS)

	assert.Equal(t, CONTINUOUS, data.Continuous)
}

func TestAnimationData_SetDelay(t *testing.T) {
	data := AnimationData()
	data.SetDelay(200)

	assert.Equal(t, 200, data.Delay)
}

func TestAnimationData_SetDelayMod(t *testing.T) {
	data := AnimationData()
	data.SetDelayMod(2.0)

	assert.Equal(t, 2.0, data.DelayMod)
}

func TestAnimationData_SetDirection(t *testing.T) {
	data := AnimationData()
	data.SetDirection(BACKWARD)

	assert.Equal(t, BACKWARD, data.Direction)
}

func TestAnimationData_SetDistance(t *testing.T) {
	data := AnimationData()
	data.SetDistance(35)

	assert.Equal(t, 35, data.Distance)
}

func TestAnimationData_SetID(t *testing.T) {
	data := AnimationData()
	data.SetID("TEST")

	assert.Equal(t, "TEST", data.Id)
}

func TestAnimationData_SetSection(t *testing.T) {
	data := AnimationData()
	data.SetSection("SECT")

	assert.Equal(t, "SECT", data.Section)
}

func TestAnimationData_SetSpacing(t *testing.T) {
	data := AnimationData()
	data.SetSpacing(4)

	assert.Equal(t, 4, data.Spacing)
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

	//cc := ColorContainer{}
	//cc.AddColor(0xFF).AddColor(0xFF00)
	//cc2 := ColorContainer{}
	//cc2.AddColor(0xFF0000)
	//data.AddColor(&cc)
	//data.AddColor(&cc2)

	json := data.Json()

	assert.Equal(t,
		`DATA:{"animation":"Meteor","center":50,"colors":[{"colors":[255,65280]},{"colors":[16711680]}],"continuous":false,"delay":10,"delayMod":1.5,"direction":"BACKWARD","distance":45,"id":"TEST","section":"SECT","spacing":5}`,
		string(json))
}

func TestAnimationData_FromGoodJson(t *testing.T) {
	jsonStr := `DATA:{"animation":"Meteor","center":50,"colors":[{"colors":[255,65280]},{"colors":[16711680]}],"continuous":false,"delay":10,"delayMod":1.5,"direction":"BACKWARD","distance":45,"id":"TEST","section":"SECT","spacing":5}`

	data, _ := AnimationDataFromJson(jsonStr)

	assert.Equal(t, "Meteor", data.Animation)
	assert.Len(t, data.Colors, 2)
	assert.Len(t, data.Colors[0].Colors, 2)
	assert.Len(t, data.Colors[1].Colors, 1)
	assert.Equal(t, 0xFF, data.Colors[0].Colors[0])
	assert.Equal(t, 0xFF00, data.Colors[0].Colors[1])
	assert.Equal(t, 0xFF0000, data.Colors[1].Colors[0])
	assert.Equal(t, 50, data.Center)
	assert.Equal(t, NONCONTINUOUS, data.Continuous)
	assert.Equal(t, 10, data.Delay)
	assert.Equal(t, 1.5, data.DelayMod)
	assert.Equal(t, BACKWARD, data.Direction)
	assert.Equal(t, 45, data.Distance)
	assert.Equal(t, "TEST", data.Id)
	assert.Equal(t, "SECT", data.Section)
	assert.Equal(t, 5, data.Spacing)
}

func TestAnimationData_FromBadJson(t *testing.T) {
	jsonStr := "{}"

	data, _ := AnimationDataFromJson(jsonStr)

	assert.Equal(t, "Color", data.Animation)
	assert.Len(t, data.Colors, 0)
	assert.Equal(t, -1, data.Center)
	assert.Equal(t, DEFAULT, data.Continuous)
	assert.Equal(t, -1, data.Delay)
	assert.Equal(t, 1.0, data.DelayMod)
	assert.Equal(t, FORWARD, data.Direction)
	assert.Equal(t, -1, data.Distance)
	assert.Equal(t, "", data.Id)
	assert.Equal(t, "", data.Section)
	assert.Equal(t, -1, data.Spacing)
}

func TestAnimationData_FromJson_Err(t *testing.T) {
	jsonStr := `{"animation":5}`

	_, err := AnimationDataFromJson(jsonStr)
	assert.NotNil(t, err)
}

func TestAnimationData_ContinuousFromJson(t *testing.T) {
	// Tests for other continuous values

	jsonStr := `{"continuous":null}`
	data, _ := AnimationDataFromJson(jsonStr)
	assert.Equal(t, DEFAULT, data.Continuous)

	jsonStr = `{"continuous":true}`
	data, _ = AnimationDataFromJson(jsonStr)
	assert.Equal(t, CONTINUOUS, data.Continuous)

	jsonStr = `{"continuous":false}`
	data, _ = AnimationDataFromJson(jsonStr)
	assert.Equal(t, NONCONTINUOUS, data.Continuous)

	jsonStr = `{"continuous":-1}`
	data, _ = AnimationDataFromJson(jsonStr)
	assert.Equal(t, DEFAULT, data.Continuous)
}
