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

func TestSection(t *testing.T) {
	sect := Section()

	if sect.Name != "" {
		log.Print("Failed sect.Name check")
		t.Fail()
	} else if sect.StartPixel != -1 {
		log.Print("Failed sect.StartPixel check")
		t.Fail()
	} else if sect.EndPixel != -1 {
		log.Print("Failed sect.EndPixel check")
		t.Fail()
	} else if sect.PhysicalStart != -1 {
		log.Print("Failed sect.PhysicalStart check")
		t.Fail()
	} else if sect.NumLEDs != 0 {
		log.Print("Failed sect.NumLEDs check")
		t.Fail()
	}
}

func TestSection_SetName(t *testing.T) {
	sect := Section()
	sect.SetName("Test")

	if sect.Name != "Test" {
		t.Fail()
	}
}

func TestSection_SetStartPixel(t *testing.T) {
	sect := Section()
	sect.SetStartPixel(5)

	if sect.StartPixel != 5 {
		t.Fail()
	}
}

func TestSection_SetEndPixel(t *testing.T) {
	sect := Section()
	sect.SetEndPixel(20)

	if sect.EndPixel != 20 {
		t.Fail()
	}
}

func TestSection_Json(t *testing.T) {
	sect := Section()
	sect.SetName("Section")
	sect.SetStartPixel(30)
	sect.SetEndPixel(40)

	json := sect.Json()
	if string(json) != `SECT:{"name":"Section","startPixel":30,"endPixel":40}` {
		t.Fail()
	}
}

func TestSection_FromGoodJson(t *testing.T) {
	jsonStr := `SECT:{"physicalStart":0,"numLEDs":240,"name":"section","startPixel":0,"endPixel":239}`

	sect := SectionFromJson(jsonStr)

	if sect.Name != "section" {
		log.Print("Failed sect.Name check")
		t.Fail()
	} else if sect.StartPixel != 0 {
		log.Print("Failed sect.StartPixel check")
		t.Fail()
	} else if sect.EndPixel != 239 {
		log.Print("Failed sect.EndPixel check")
		t.Fail()
	} else if sect.PhysicalStart != 0 {
		log.Print("Failed sect.PhysicalStart check")
		t.Fail()
	} else if sect.NumLEDs != 240 {
		log.Print("Failed sect.NumLEDs check")
		t.Fail()
	}
}

func TestSection_FromBadJson(t *testing.T) {
	jsonStr := "{}"

	sect := SectionFromJson(jsonStr)

	if sect.Name != "" {
		log.Print("Failed sect.Name check")
		t.Fail()
	} else if sect.StartPixel != -1 {
		log.Print("Failed sect.StartPixel check")
		t.Fail()
	} else if sect.EndPixel != -1 {
		log.Print("Failed sect.EndPixel check")
		t.Fail()
	} else if sect.PhysicalStart != -1 {
		log.Print("Failed sect.PhysicalStart check")
		t.Fail()
	} else if sect.NumLEDs != 0 {
		log.Print("Failed sect.NumLEDs check")
		t.Fail()
	}
}
