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

func TestAnimationInfo_FromGoodJson(t *testing.T) {
	jsonStr := `AINF:{"name":"Alternate","abbr":"ALT","description":"A description","signatureFile":"alternate.png","repetitive":true,"minimumColors":2,"unlimitedColors":true,"center":"NOTUSED","delay":"USED","direction":"NOTUSED","distance":"NOTUSED","spacing":"NOTUSED","delayDefault":1000,"distanceDefault":20,"spacingDefault":3}`

	info, _ := AnimationInfoFromJson(jsonStr)

	assert.Equal(t, info.Name, "Alternate")
	assert.Equal(t, info.Abbr, "ALT")
	assert.Equal(t, info.Description, "A description")
	assert.Equal(t, info.SignatureFile, "alternate.png")
	assert.True(t, info.Repetitive)
	assert.Equal(t, info.MinimumColors, 2)
	assert.True(t, info.UnlimitedColors)
	assert.Equal(t, info.Center, NOTUSED)
	assert.Equal(t, info.Delay, USED)
	assert.Equal(t, info.Direction, NOTUSED)
	assert.Equal(t, info.Distance, NOTUSED)
	assert.Equal(t, info.Spacing, NOTUSED)
	assert.Equal(t, info.DelayDefault, 1000)
	assert.Equal(t, info.DistanceDefault, 20)
	assert.Equal(t, info.SpacingDefault, 3)
}

func TestAnimationInfo_FromBadJson(t *testing.T) {
	jsonStr := `AINF:{}`

	info, _ := AnimationInfoFromJson(jsonStr)

	assert.Equal(t, info.Name, "")
	assert.Equal(t, info.Abbr, "")
	assert.Equal(t, info.Description, "")
	assert.Equal(t, info.SignatureFile, "")
	assert.False(t, info.Repetitive)
	assert.Equal(t, info.MinimumColors, -1)
	assert.False(t, info.UnlimitedColors)
	assert.Equal(t, info.Center, NOTUSED)
	assert.Equal(t, info.Delay, NOTUSED)
	assert.Equal(t, info.Direction, NOTUSED)
	assert.Equal(t, info.Distance, NOTUSED)
	assert.Equal(t, info.Spacing, NOTUSED)
	assert.Equal(t, info.DelayDefault, -1)
	assert.Equal(t, info.DistanceDefault, -1)
	assert.Equal(t, info.SpacingDefault, -1)
}

func TestAnimationInfo_FromJson_Err(t *testing.T) {
	jsonStr := `AINF:{"name":3}`

	_, err := AnimationInfoFromJson(jsonStr)
	assert.NotNil(t, err)
}
