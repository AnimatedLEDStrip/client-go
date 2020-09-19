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

func TestSection(t *testing.T) {
	sect := Section()

	assert.Equal(t, sect.Name, "")
	assert.Equal(t, sect.StartPixel, -1)
	assert.Equal(t, sect.EndPixel, -1)
	assert.Equal(t, sect.PhysicalStart, -1)
	assert.Equal(t, sect.NumLEDs, 0)
}

func TestSection_SetName(t *testing.T) {
	sect := Section()
	sect.SetName("Test")

	assert.Equal(t, sect.Name, "Test")
}

func TestSection_SetStartPixel(t *testing.T) {
	sect := Section()
	sect.SetStartPixel(5)

	assert.Equal(t, sect.StartPixel, 5)
}

func TestSection_SetEndPixel(t *testing.T) {
	sect := Section()
	sect.SetEndPixel(20)

	assert.Equal(t, sect.EndPixel, 20)
}

func TestSection_Json(t *testing.T) {
	sect := Section()
	sect.SetName("Section")
	sect.SetStartPixel(30)
	sect.SetEndPixel(40)

	json := sect.Json()
	assert.Equal(t, string(json), `SECT:{"name":"Section","startPixel":30,"endPixel":40}`)
}

func TestSection_FromGoodJson(t *testing.T) {
	jsonStr := `SECT:{"physicalStart":0,"numLEDs":240,"name":"section","startPixel":0,"endPixel":239}`

	sect, _ := SectionFromJson(jsonStr)

	assert.Equal(t, sect.Name, "section")
	assert.Equal(t, sect.StartPixel, 0)
	assert.Equal(t, sect.EndPixel, 239)
	assert.Equal(t, sect.PhysicalStart, 0)
	assert.Equal(t, sect.NumLEDs, 240)
}

func TestSection_FromBadJson(t *testing.T) {
	jsonStr := "SECT:{}"

	sect, _ := SectionFromJson(jsonStr)

	assert.Equal(t, sect.Name, "")
	assert.Equal(t, sect.StartPixel, -1)
	assert.Equal(t, sect.EndPixel, -1)
	assert.Equal(t, sect.PhysicalStart, -1)
	assert.Equal(t, sect.NumLEDs, 0)
}

func TestSection_FromJsonErr(t *testing.T) {
	jsonStr := `SECT:{"name":false}`

	_, err := SectionFromJson(jsonStr)

	assert.NotNil(t, err)
}
