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
	NumLEDs           int  `json:"numLEDs"`
	Pin               int  `json:"pin"`
	ImageDebugging    bool `json:"imageDebugging"`
	RendersBeforeSave int  `json:"rendersBeforeSave"`
	ThreadCount       int  `json:"threadCount"`
}

func StripInfoFromJson(data string) (*stripInfo, error) {
	dataStr := strings.TrimPrefix(data, "SINF:")
	info := stripInfo{
		NumLEDs:           0,
		Pin:               -1,
		ImageDebugging:    false,
		RendersBeforeSave: -1,
		ThreadCount:       100,
	}
	err := json.Unmarshal([]byte(dataStr), &info)
	if err != nil {
		return nil, err
	} else {
		return &info, nil
	}
}
