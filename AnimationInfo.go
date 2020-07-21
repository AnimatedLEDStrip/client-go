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
	Name            string
	Abbr            string
	Description     string
	SignatureFile   string
	Repetitive      bool
	MinimumColors   int
	UnlimitedColors bool
	Center          ParamUsage
	Delay           ParamUsage
	Direction       ParamUsage
	Distance        ParamUsage
	Spacing         ParamUsage
	DelayDefault    int
	DistanceDefault int
	SpacingDefault  int
}

func AnimationInfo() *animationInfo {
	return &animationInfo{
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
}

func AnimationInfoFromJson(data string) *animationInfo {
	info := AnimationInfo()

	dataStr := strings.TrimPrefix(data, "AINF:")
	var infoJson interface{}
	_ = json.Unmarshal([]byte(dataStr), &infoJson)
	i := infoJson.(map[string]interface{})

	name, _ := i["name"].(string)
	info.Name = name
	// No need to specify a default here because
	// default for name is an empty string

	abbr, _ := i["abbr"].(string)
	info.Abbr = abbr
	// No need to specify a default here because
	// default for abbr is an empty string

	desc, _ := i["description"].(string)
	info.Description = desc
	// No need to specify a default here because
	// default for description is an empty string

	sig, _ := i["signatureFile"].(string)
	info.SignatureFile = sig
	// No need to specify a default here because
	// default for signatureFile is an empty string

	rep, ok := i["repetitive"].(bool)
	if !ok {
		rep = false
	}
	info.Repetitive = rep

	min, ok := i["minimumColors"].(float64)
	if !ok {
		min = -1
	}
	info.MinimumColors = int(min)

	unlimited, _ := i["unlimitedColors"].(bool)
	info.UnlimitedColors = unlimited

	center, _ := i["center"].(string)
	info.Center = ParamUsageFromString(center)
	// No need to specify a default here because
	// ParamUsageFromString returns NOTUSED for an
	// empty string

	delay, _ := i["delay"].(string)
	info.Delay = ParamUsageFromString(delay)
	// No need to specify a default here because
	// ParamUsageFromString returns NOTUSED for an
	// empty string

	direction := i["direction"].(string)
	info.Direction = ParamUsageFromString(direction)
	// No need to specify a default here because
	// ParamUsageFromString returns NOTUSED for an
	// empty string

	distance := i["distance"].(string)
	info.Distance = ParamUsageFromString(distance)
	// No need to specify a default here because
	// ParamUsageFromString returns NOTUSED for an
	// empty string

	spacing := i["spacing"].(string)
	info.Spacing = ParamUsageFromString(spacing)
	// No need to specify a default here because
	// ParamUsageFromString returns NOTUSED for an
	// empty string

	delayDefault, ok := i["delayDefault"].(float64)
	if !ok {
		delayDefault = -1
	}
	info.DelayDefault = int(delayDefault)

	distanceDefault, ok := i["distanceDefault"].(float64)
	if !ok {
		distanceDefault = -1
	}
	info.DistanceDefault = int(distanceDefault)

	spacingDefault, ok := i["spacingDefault"].(float64)
	if !ok {
		spacingDefault = -1
	}
	info.SpacingDefault = int(spacingDefault)

	return info
}
