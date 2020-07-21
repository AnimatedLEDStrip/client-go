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

func TestAnimationInfoFromJson(t *testing.T) {
	jsonStr := `AINF:{"name":"Alternate","abbr":"ALT","description":"A description","signatureFile":"alternate.png","repetitive":true,"minimumColors":2,"unlimitedColors":false,"center":"NOTUSED","delay":"USED","direction":"NOTUSED","distance":"NOTUSED","spacing":"NOTUSED","delayDefault":1000,"distanceDefault":-1,"spacingDefault":3}`

	info := AnimationInfoFromJson(jsonStr)

	if info.Name != "Alternate" {
		t.Fail()
	} else if info.Abbr != "ALT" {
		t.Fail()
	} else if info.Description != "A description" {
		t.Fail()
	} else if info.SignatureFile != "alternate.png" {
		t.Fail()
	} else if info.Repetitive != true {
		t.Fail()
	} else if info.MinimumColors != 2 {
		t.Fail()
	} else if info.UnlimitedColors == true {
		t.Fail()
	} else if info.Center != NOTUSED {
		t.Fail()
	} else if info.Delay != USED {
		t.Fail()
	} else if info.Direction != NOTUSED {
		t.Fail()
	} else if info.Distance != NOTUSED {
		t.Fail()
	} else if info.Spacing != NOTUSED {
		t.Fail()
	} else if info.DelayDefault != 1000 {
		t.Fail()
	} else if info.DistanceDefault != -1 {
		t.Fail()
	} else if info.SpacingDefault != 3 {
		t.Fail()
	}
}
