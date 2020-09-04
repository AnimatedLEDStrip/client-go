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

func TestAnimationInfoFromJson(t *testing.T) {
	jsonStr := `AINF:{"name":"Alternate","abbr":"ALT","description":"A description","signatureFile":"alternate.png","repetitive":true,"minimumColors":2,"unlimitedColors":false,"center":"NOTUSED","delay":"USED","direction":"NOTUSED","distance":"NOTUSED","spacing":"NOTUSED","delayDefault":1000,"distanceDefault":-1,"spacingDefault":3}`

	info, _ := AnimationInfoFromJson(jsonStr)

	if info.Name != "Alternate" {
		log.Print("Failed info.Name check")
		t.Fail()
	} else if info.Abbr != "ALT" {
		log.Print("Failed info.Abbr check")
		t.Fail()
	} else if info.Description != "A description" {
		log.Print("Failed info.Description check")
		t.Fail()
	} else if info.SignatureFile != "alternate.png" {
		log.Print("Failed info.SignatureFile check")
		t.Fail()
	} else if info.Repetitive != true {
		log.Print("Failed info.Repetitive check")
		t.Fail()
	} else if info.MinimumColors != 2 {
		log.Print("Failed info.MinimumColors check")
		t.Fail()
	} else if info.UnlimitedColors == true {
		log.Print("Failed info.UnlimitedColors check")
		t.Fail()
	} else if info.Center != NOTUSED {
		log.Print("Failed info.Center check")
		t.Fail()
	} else if info.Delay != USED {
		log.Print("Failed info.Delay check")
		t.Fail()
	} else if info.Direction != NOTUSED {
		log.Print("Failed info.Direction check")
		t.Fail()
	} else if info.Distance != NOTUSED {
		log.Print("Failed info.Distance check")
		t.Fail()
	} else if info.Spacing != NOTUSED {
		log.Print("Failed info.Spacing check")
		t.Fail()
	} else if info.DelayDefault != 1000 {
		log.Print("Failed info.DelayDefault check")
		t.Fail()
	} else if info.DistanceDefault != -1 {
		log.Print("Failed info.DistanceDefault check")
		t.Fail()
	} else if info.SpacingDefault != 3 {
		log.Print("Failed info.SpacingDefault check")
		t.Fail()
	}
}
