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

type animationInfo struct {
	Name            string     `json:"name"`
	Abbr            string     `json:"abbr"`
	Description     string     `json:"description"`
	SignatureFile   string     `json:"signatureFile"`
	Repetitive      bool       `json:"repetitive"`
	MinimumColors   int        `json:"minimumColors"`
	UnlimitedColors bool       `json:"unlimitedColors"`
	Center          ParamUsage `json:"center"`
	Delay           ParamUsage `json:"delay"`
	Direction       ParamUsage `json:"direction"`
	Distance        ParamUsage `json:"distance"`
	Spacing         ParamUsage `json:"spacing"`
	DelayDefault    int        `json:"delayDefault"`
	DistanceDefault int        `json:"distanceDefault"`
	SpacingDefault  int        `json:"spacingDefault"`
}

func AnimationInfoFromJson(data string) (*animationInfo, error) {
	dataStr := strings.TrimPrefix(data, "AINF:")
	animInfo := animationInfo{
		Name:            "",
		Abbr:            "",
		Description:     "",
		SignatureFile:   "",
		Repetitive:      false,
		MinimumColors:   -1,
		UnlimitedColors: false,
		Center:          NOTUSED,
		Delay:           NOTUSED,
		Direction:       NOTUSED,
		Distance:        NOTUSED,
		Spacing:         NOTUSED,
		DelayDefault:    -1,
		DistanceDefault: -1,
		SpacingDefault:  -1,
	}

	var infoFilter struct {
		Name            string `json:"name,omitempty"`
		Abbr            string `json:"abbr,omitempty"`
		Description     string `json:"description,omitempty"`
		SignatureFile   string `json:"signatureFile,omitempty"`
		Repetitive      bool   `json:"repetitive,omitempty"`
		MinimumColors   int    `json:"minimumColors,omitempty"`
		UnlimitedColors bool   `json:"unlimitedColors,omitempty"`
		DelayDefault    int    `json:"delayDefault,omitempty"`
		DistanceDefault int    `json:"distanceDefault,omitempty"`
		SpacingDefault  int    `json:"spacingDefault,omitempty"`
	}
	err := json.Unmarshal([]byte(dataStr), &infoFilter)
	if err != nil {
		return nil, err
	}
	jsonBytes, _ := json.Marshal(&infoFilter)
	_ = json.Unmarshal(jsonBytes, &animInfo)

	var getUsages interface{}
	_ = json.Unmarshal([]byte(dataStr), &getUsages)
	usg := getUsages.(map[string]interface{})

	// No need to specify a default for
	// center, delay, direction, distance or spacing
	// because ParamUsageFromString returns NOTUSED
	// for an empty string
	center, _ := usg["center"].(string)
	animInfo.Center = ParamUsageFromString(center)
	delay, _ := usg["delay"].(string)
	animInfo.Delay = ParamUsageFromString(delay)
	direction := usg["direction"].(string)
	animInfo.Direction = ParamUsageFromString(direction)
	distance := usg["distance"].(string)
	animInfo.Distance = ParamUsageFromString(distance)
	spacing := usg["spacing"].(string)
	animInfo.Spacing = ParamUsageFromString(spacing)

	return &animInfo, nil
}
