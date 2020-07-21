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

type stripInfo struct {
	NumLEDs           int
	Pin               int
	ImageDebugging    bool
	RendersBeforeSave int
	ThreadCount       int
}

func StripInfo() *stripInfo {
	return &stripInfo{
		NumLEDs:           0,
		Pin:               -1,
		ImageDebugging:    false,
		RendersBeforeSave: -1,
		ThreadCount:       100,
	}
}

func StripInfoFromJson(data string) *stripInfo {
	info := StripInfo()

	dataStr := strings.TrimPrefix(data, "SINF:")
	var infoJson interface{}
	_ = json.Unmarshal([]byte(dataStr), &infoJson)
	i := infoJson.(map[string]interface{})

	num, ok := i["numLEDs"].(float64)
	if !ok {
		num = 0
	}

	info.NumLEDs = int(num)

	pin, ok := i["pin"].(float64)
	if !ok {
		pin = -1
	}

	info.Pin = int(pin)

	imgDebug, ok := i["imageDebugging"].(bool)
	if !ok {
		imgDebug = false
	}

	info.ImageDebugging = imgDebug

	renders, ok := i["rendersBeforeSave"].(float64)
	if !ok {
		renders = -1
	}

	info.RendersBeforeSave = int(renders)

	threads, ok := i["threadCount"].(float64)
	if !ok {
		threads = 100
	}

	info.ThreadCount = int(threads)

	return info
}
