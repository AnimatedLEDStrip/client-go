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
	"strconv"
	"strings"
)

type ColorContainer struct {
	Colors []int
}

func (c *ColorContainer) AddColor(color int) *ColorContainer {
	c.Colors = append(c.Colors, color)
	return c
}

func (c ColorContainer) Json() string {
	var stringParts []string
	stringParts = append(stringParts, `{"colors":[`)
	if len(c.Colors) != 0 {
		for _, c := range c.Colors {
			stringParts = append(stringParts, strconv.Itoa(c))
			stringParts = append(stringParts, ",")
		}
		stringParts = stringParts[:len(stringParts)-1]
	}
	stringParts = append(stringParts, "]}")
	return strings.Join(stringParts, "")
}

func ColorContainerFromJson(ccJson interface{}) *ColorContainer {
	cc := ColorContainer{}
	colors, _ := ccJson.(map[string]interface{})["colors"].([]interface{})
	for _, c := range colors {
		cc.AddColor(int(c.(float64)))
	}
	return &cc
}
