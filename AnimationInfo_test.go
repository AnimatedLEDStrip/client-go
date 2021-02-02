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

//func TestAnimationInfo_FromGoodJson(t *testing.T) {
//	jsonStr := `AINF:{"name":"Alternate","abbr":"ALT","description":"A description","signatureFile":"alternate.png","repetitive":true,"minimumColors":2,"unlimitedColors":true,"center":"NOTUSED","delay":"USED","direction":"NOTUSED","distance":"NOTUSED","spacing":"NOTUSED","delayDefault":1000,"distanceDefault":20,"spacingDefault":3}`
//
//	info, _ := AnimationInfoFromJson(jsonStr)
//
//	assert.Equal(t, "Alternate", info.Name)
//	assert.Equal(t, "ALT", info.Abbr)
//	assert.Equal(t, "A description", info.Description)
//	assert.Equal(t, "alternate.png", info.SignatureFile)
//	assert.True(t, info.Repetitive)
//	assert.Equal(t, 2, info.MinimumColors)
//	assert.True(t, info.UnlimitedColors)
//	assert.Equal(t, NOTUSED, info.Center)
//	assert.Equal(t, USED, info.Delay)
//	assert.Equal(t, NOTUSED, info.Direction)
//	assert.Equal(t, NOTUSED, info.Distance)
//	assert.Equal(t, NOTUSED, info.Spacing)
//	assert.Equal(t, 1000, info.DelayDefault)
//	assert.Equal(t, 20, info.DistanceDefault)
//	assert.Equal(t, 3, info.SpacingDefault)
//}
//
//func TestAnimationInfo_FromBadJson(t *testing.T) {
//	jsonStr := `AINF:{}`
//
//	info, _ := AnimationInfoFromJson(jsonStr)
//
//	assert.Equal(t, "", info.Name)
//	assert.Equal(t, "", info.Abbr)
//	assert.Equal(t, "", info.Description)
//	assert.Equal(t, "", info.SignatureFile)
//	assert.False(t, info.Repetitive)
//	assert.Equal(t, -1, info.MinimumColors)
//	assert.False(t, info.UnlimitedColors)
//	assert.Equal(t, NOTUSED, info.Center)
//	assert.Equal(t, NOTUSED, info.Delay)
//	assert.Equal(t, NOTUSED, info.Direction)
//	assert.Equal(t, NOTUSED, info.Distance)
//	assert.Equal(t, NOTUSED, info.Spacing)
//	assert.Equal(t, -1, info.DelayDefault)
//	assert.Equal(t, -1, info.DistanceDefault)
//	assert.Equal(t, -1, info.SpacingDefault)
//}
//
//func TestAnimationInfo_FromJson_Err(t *testing.T) {
//	jsonStr := `AINF:{"name":3}`
//
//	_, err := AnimationInfoFromJson(jsonStr)
//	assert.NotNil(t, err)
//}
