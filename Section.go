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

type section struct {
	Name          string `json:"name"`
	StartPixel    int    `json:"startPixel"`
	EndPixel      int    `json:"endPixel"`
	PhysicalStart int    `json:"physicalStart"`
	NumLEDs       int    `json:"numLEDs"`
}

func (s *section) MarshalJSON() ([]byte, error) {
	var tmp struct {
		Name       string `json:"name"`
		StartPixel int    `json:"startPixel"`
		EndPixel   int    `json:"endPixel"`
	}
	tmp.Name = s.Name
	tmp.StartPixel = s.StartPixel
	tmp.EndPixel = s.EndPixel
	return json.Marshal(&tmp)
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

func (s *section) Json() ([]byte, error) {
	str, err := json.Marshal(s)
	if err != nil {
		return nil, err
	} else {
		return append([]byte("SECT:"), str...), nil
	}
}

func SectionFromJson(data string) (*section, error) {
	dataStr := strings.TrimPrefix(data, "SECT:")
	sect := Section()
	err := json.Unmarshal([]byte(dataStr), sect)
	if err != nil {
		return nil, err
	} else {
		return sect, nil
	}
}

func (s *section) SetName(name string) *section {
	s.Name = name
	return s
}

func (s *section) SetStartPixel(pixel int) *section {
	s.StartPixel = pixel
	s.NumLEDs = s.EndPixel - s.StartPixel
	return s
}

func (s *section) SetEndPixel(pixel int) *section {
	s.EndPixel = pixel
	s.NumLEDs = s.EndPixel - s.StartPixel
	return s
}
