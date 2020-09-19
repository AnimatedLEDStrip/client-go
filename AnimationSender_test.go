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

func TestAnimationSender_processData_partialData(t *testing.T) {
	called := false
	callback := func(data string) {
		called = true
		assert.Equal(t, data, "DATA:{}")
	}

	sender := AnimationSender{
		RunningAnimations:   NewRunningAnimationMap(),
		Sections:            map[string]*section{},
		SupportedAnimations: map[string]*animationInfo{},
		onReceiveCallback:   callback,
	}

	jsonStr := "DAT"
	sender.processData([]byte(jsonStr))

	jsonStr = "A:{};;;"
	sender.processData([]byte(jsonStr))

	assert.True(t, called)
}

func TestAnimationSender_processData_onReceiveCallback(t *testing.T) {
	called := false
	callback := func(data string) {
		called = true
		assert.Equal(t, data, "DATA:{}")
	}

	sender := AnimationSender{
		RunningAnimations:   NewRunningAnimationMap(),
		Sections:            map[string]*section{},
		SupportedAnimations: map[string]*animationInfo{},
		onReceiveCallback:   callback,
	}

	jsonStr := `DATA:{};;;`
	sender.processData([]byte(jsonStr))

	assert.True(t, called)
}

func TestAnimationSender_processData_AnimationData(t *testing.T) {
	called := false
	callback := func(data *animationData) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:          NewRunningAnimationMap(),
		Sections:                   map[string]*section{},
		SupportedAnimations:        map[string]*animationInfo{},
		onNewAnimationDataCallback: callback,
	}

	jsonStr := `DATA:{"id":"test"};;;`
	sender.processData([]byte(jsonStr))

	_, ok := sender.RunningAnimations.Load("test")

	assert.True(t, ok)
	assert.True(t, called)
}

func TestAnimationSender_processData_AnimationInfo(t *testing.T) {
	t.Skip()
	called := false
	callback := func(info *animationInfo) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:          NewRunningAnimationMap(),
		Sections:                   map[string]*section{},
		SupportedAnimations:        map[string]*animationInfo{},
		onNewAnimationInfoCallback: callback,
	}

	jsonStr := "AINF:{};;;"
	sender.processData([]byte(jsonStr))

	assert.True(t, called)
}

func TestAnimationSender_processData_EndAnimation(t *testing.T) {
	called := false
	callback := func(end *endAnimation) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:         NewRunningAnimationMap(),
		Sections:                  map[string]*section{},
		SupportedAnimations:       map[string]*animationInfo{},
		onNewEndAnimationCallback: callback,
	}

	jsonStr := `DATA:{"id":"test"};;;`
	sender.processData([]byte(jsonStr))

	_, ok := sender.RunningAnimations.Load("test")

	assert.True(t, ok)

	jsonStr = `END :{"id":"test"};;;`
	sender.processData([]byte(jsonStr))

	_, ok = sender.RunningAnimations.Load("test")

	assert.False(t, ok)
	assert.True(t, called)
}

func TestAnimationSender_processData_Message(t *testing.T) {
	called := false
	callback := func(msg *message) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:    NewRunningAnimationMap(),
		Sections:             map[string]*section{},
		SupportedAnimations:  map[string]*animationInfo{},
		onNewMessageCallback: callback,
	}

	jsonStr := "MSG :{};;;"
	sender.processData([]byte(jsonStr))

	assert.True(t, called)
}

func TestAnimationSender_processData_Section(t *testing.T) {
	called := false
	callback := func(sect *section) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:    NewRunningAnimationMap(),
		Sections:             map[string]*section{},
		SupportedAnimations:  map[string]*animationInfo{},
		onNewSectionCallback: callback,
	}

	jsonStr := "SECT:{};;;"
	sender.processData([]byte(jsonStr))

	assert.True(t, called)
	assert.Len(t, sender.Sections, 1)
}

func TestAnimationSender_processData_StripInfo(t *testing.T) {
	called := false
	callback := func(info *stripInfo) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:      NewRunningAnimationMap(),
		Sections:               map[string]*section{},
		SupportedAnimations:    map[string]*animationInfo{},
		onNewStripInfoCallback: callback,
	}

	jsonStr := "SINF:{};;;"
	sender.processData([]byte(jsonStr))

	assert.True(t, called)
	assert.NotNil(t, sender.StripInfo)
}
