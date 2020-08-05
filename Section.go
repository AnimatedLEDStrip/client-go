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
	"strconv"
	"strings"
)

type section struct {
	Name          string
	StartPixel    int
	EndPixel      int
	PhysicalStart int
	NumLEDs       int
}

func Section() *section {
	return &section{
		Name:          "",
		StartPixel:    -1,
		EndPixel:      -1,
		PhysicalStart: -1,
		NumLEDs:       0,
	}
}

func (s *section) SetName(name string) *section {
	s.Name = name
	return s
}

func (s *section) SetStartPixel(pixel int) *section {
	s.StartPixel = pixel
	return s
}

func (s *section) SetEndPixel(pixel int) *section {
	s.EndPixel = pixel
	return s
}

func (s *section) SetPhysicalStart(pixel int) *section {
	s.PhysicalStart = pixel
	return s
}

func (s *section) SetNumLEDs(leds int) *section {
	s.NumLEDs = leds
	return s
}

func (s *section) Json() string {
	var stringParts []string
	stringParts = append(stringParts, "SECT:{")
	stringParts = append(stringParts, `"name":"`)
	stringParts = append(stringParts, s.Name)
	stringParts = append(stringParts, `","startPixel":`)
	stringParts = append(stringParts, strconv.Itoa(s.StartPixel))
	stringParts = append(stringParts, `,"endPixel":`)
	stringParts = append(stringParts, strconv.Itoa(s.EndPixel))
	stringParts = append(stringParts, "}")

	return strings.Join(stringParts, "")
}

func SectionFromJson(data string) *section {
	sect := Section()

	dataStr := strings.TrimPrefix(data, "SECT:")
	var sectJson interface{}
	_ = json.Unmarshal([]byte(dataStr), &sectJson)
	s := sectJson.(map[string]interface{})

	name, _ := s["name"].(string)
	sect.Name = name
	// No need to specify a default here because
	// default for name is an empty string

	start, ok := s["startPixel"].(float64)
	if !ok {
		start = -1
	}
	sect.StartPixel = int(start)

	end, ok := s["endPixel"].(float64)
	if !ok {
		end = -1
	}
	sect.EndPixel = int(end)

	pStart, ok := s["physicalStart"].(float64)
	if !ok {
		pStart = -1
	}
	sect.PhysicalStart = int(pStart)

	num, ok := s["numLEDs"].(float64)
	if !ok {
		num = 0
	}
	sect.NumLEDs = int(num)

	return sect
}
