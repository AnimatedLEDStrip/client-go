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
	"encoding/json"
	"strings"
)

type animationData struct {
	Animation  string            `json:"animation"`
	Center     int               `json:"center"`
	Colors     []*ColorContainer `json:"colors"`
	Continuous Continuous        `json:"continuous"`
	Delay      int               `json:"delay"`
	DelayMod   float64           `json:"delayMod"`
	Direction  Direction         `json:"direction"`
	Distance   int               `json:"distance"`
	Id         string            `json:"id"`
	Section    string            `json:"section"`
	Spacing    int               `json:"spacing"`
}

func AnimationData() *animationData {
	return &animationData{
		Animation:  "Color",
		Center:     -1,
		Continuous: DEFAULT,
		Delay:      -1,
		DelayMod:   1.0,
		Direction:  FORWARD,
		Distance:   -1,
		Id:         "",
		Section:    "",
		Spacing:    -1,
	}
}

func (d *animationData) Json() []byte {
	str, _ := json.Marshal(d)
	return append([]byte("DATA:"), str...)
}

func AnimationDataFromJson(data string) (*animationData, error) {
	dataStr := strings.TrimPrefix(data, "DATA:")
	animData := AnimationData()

	var dataFilter struct {
		Animation string            `json:"animation,omitempty"`
		Center    int               `json:"center,omitempty"`
		Colors    []*ColorContainer `json:"colors,omitempty"`
		Delay     int               `json:"delay,omitempty"`
		DelayMod  float64           `json:"delayMod,omitempty"`
		Distance  int               `json:"distance,omitempty"`
		Id        string            `json:"id,omitempty"`
		Section   string            `json:"section,omitempty"`
		Spacing   int               `json:"spacing,omitempty"`
	}
	err := json.Unmarshal([]byte(dataStr), &dataFilter)
	if err != nil {
		return nil, err
	}
	jsonBytes, _ := json.Marshal(&dataFilter)
	_ = json.Unmarshal(jsonBytes, &animData)

	var temp interface{}
	_ = json.Unmarshal([]byte(dataStr), &temp)
	remainingData := temp.(map[string]interface{})

	continuous, _ := remainingData["continuous"]
	switch t := continuous.(type) {
	case nil:
		animData.Continuous = DEFAULT
	case bool:
		if t {
			animData.Continuous = CONTINUOUS
		} else {
			animData.Continuous = NONCONTINUOUS
		}
	default:
		animData.Continuous = DEFAULT
	}

	direction, _ := remainingData["direction"].(string)
	animData.Direction = DirectionFromString(direction)

	return animData, nil
}

func (d *animationData) SetAnimation(anim string) *animationData {
	d.Animation = anim
	return d
}

func (d *animationData) AddColor(color *ColorContainer) *animationData {
	d.Colors = append(d.Colors, color)
	return d
}

func (d *animationData) SetCenter(pixel int) *animationData {
	d.Center = pixel
	return d
}

func (d *animationData) SetContinuous(c Continuous) *animationData {
	d.Continuous = c
	return d
}

func (d *animationData) SetDelay(time int) *animationData {
	d.Delay = time
	return d
}

func (d *animationData) SetDelayMod(multiplier float64) *animationData {
	d.DelayMod = multiplier
	return d
}

func (d *animationData) SetDirection(dir Direction) *animationData {
	d.Direction = dir
	return d
}

func (d *animationData) SetDistance(pixels int) *animationData {
	d.Distance = pixels
	return d
}

func (d *animationData) SetID(i string) *animationData {
	d.Id = i
	return d
}

func (d *animationData) SetSection(sect string) *animationData {
	d.Section = sect
	return d
}

func (d *animationData) SetSpacing(pixels int) *animationData {
	d.Spacing = pixels
	return d
}
