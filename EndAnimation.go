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

type endAnimation struct {
	Id string `json:"id"`
}

func EndAnimation() *endAnimation {
	return &endAnimation{Id: ""}
}

func (e endAnimation) Json() []byte {
	str, _ := json.Marshal(e)
	return append([]byte("END :"), str...)
}

func EndAnimationFromJson(data string) (*endAnimation, error) {
	dataStr := strings.TrimPrefix(data, "END :")
	anim := EndAnimation()
	err := json.Unmarshal([]byte(dataStr), &anim)
	if err != nil {
		return nil, err
	} else {
		return anim, nil
	}
}

func (e *endAnimation) SetId(id string) *endAnimation {
	e.Id = id
	return e
}
