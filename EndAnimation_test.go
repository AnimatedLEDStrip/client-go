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

func TestEndAnimation(t *testing.T) {
	anim := EndAnimation()

	if anim.Id != "" {
		t.Fail()
	}
}

func TestEndAnimation_SetId(t *testing.T) {
	anim := EndAnimation()

	anim.SetId("12345")

	if anim.Id != "12345" {
		t.Fail()
	}
}

func TestEndAnimation_Json(t *testing.T) {
	anim := EndAnimation()

	anim.SetId("54321")

	json := anim.Json()
	if string(json) != `END :{"id":"54321"}` {
		t.Fail()
	}
}

func TestEndAnimation_FromGoodJson(t *testing.T) {
	jsonStr := `END :{"id":"13579"}`

	anim, _ := EndAnimationFromJson(jsonStr)

	if anim.Id != "13579" {
		t.Fail()
	}
}

func TestEndAnimation_FromBadJson(t *testing.T) {
	jsonStr := `END :{}`

	anim, _ := EndAnimationFromJson(jsonStr)

	if anim.Id != "" {
		t.Fail()
	}
}

func TestEndAnimation_FromJson_Err(t *testing.T) {
	jsonStr := `END :{"id":false}`

	_, err := EndAnimationFromJson(jsonStr)

	if err == nil {
		t.Fail()
	}
}
