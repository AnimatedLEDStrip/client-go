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

func TestStripInfoFromJson(t *testing.T) {
	jsonStr := `SINF:{"numLEDs":240,"pin":12,"imageDebugging":false,"rendersBeforeSave":1000,"threadCount":100}`

	info, _ := StripInfoFromJson(jsonStr)

	if info.NumLEDs != 240 {
		log.Print("Failed info.NumLEDs check")
		t.Fail()
	} else if info.Pin != 12 {
		log.Print("Failed info.Pin check")
		t.Fail()
	} else if info.ImageDebugging != false {
		log.Print("Failed info.ImageDebugging check")
		t.Fail()
	} else if info.RendersBeforeSave != 1000 {
		log.Print("Failed info.RendersBeforeSave check")
		t.Fail()
	} else if info.ThreadCount != 100 {
		log.Print("Failed info.ThreadCount check")
		t.Fail()
	}
}
