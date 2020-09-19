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

func TestStripInfo_FromGoodJson(t *testing.T) {
	jsonStr := `SINF:{"numLEDs":240,"pin":12,"imageDebugging":true,"rendersBeforeSave":1000,"threadCount":100}`

	info, _ := StripInfoFromJson(jsonStr)

	assert.Equal(t, info.NumLEDs, 240)
	assert.Equal(t, info.Pin, 12)
	assert.True(t, info.ImageDebugging)
	assert.Equal(t, info.RendersBeforeSave, 1000)
	assert.Equal(t, info.ThreadCount, 100)
}

func TestStripInfo_FromBadJson(t *testing.T) {
	jsonStr := `SINF:{}`

	info, _ := StripInfoFromJson(jsonStr)

	assert.Equal(t, info.NumLEDs, 0)
	assert.Equal(t, info.Pin, -1)
	assert.False(t, info.ImageDebugging)
	assert.Equal(t, info.RendersBeforeSave, -1)
	assert.Equal(t, info.ThreadCount, -1)
}

func TestStripInfo_FromJson_Err(t *testing.T) {
	jsonStr := `SINF:{"numLEDs":false}`

	_, err := StripInfoFromJson(jsonStr)

	assert.NotNil(t, err)
}
