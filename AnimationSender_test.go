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
	"net"
	"testing"
)

func createConnection(t *testing.T, readyCh chan bool) {
	ln, err := net.Listen("tcp", ":2501")
	assert.Nil(t, err)
	readyCh <- true
	conn, err2 := ln.Accept()
	assert.Nil(t, err2)
	err3 := conn.Close()
	assert.Nil(t, err3)
}

func TestAnimationSender_Start(t *testing.T) {
	calledCh1 := make(chan bool)
	callback := func(ip string, port int) {
		calledCh1 <- true
		assert.Equal(t, ip, "0.0.0.0")
		assert.Equal(t, port, 2501)
	}
	calledCh2 := make(chan bool)
	callback2 := func(ip string, port int) {
		calledCh2 <- true
		assert.Equal(t, ip, "0.0.0.0")
		assert.Equal(t, port, 2501)
	}

	readyCh := make(chan bool)

	go createConnection(t, readyCh)

	ready := <-readyCh

	assert.True(t, ready)

	sender := AnimationSender{
		Address: "0.0.0.0",
		Port:    2501,
	}

	sender.SetOnConnectCallback(callback)
	sender.SetOnDisconnectCallback(callback2)

	sender.Start()

	called1 := <-calledCh1
	called2 := <-calledCh2
	assert.True(t, called1)
	assert.True(t, called2)
}

func TestAnimationSender_Start_UnableToConnect(t *testing.T) {
	calledCh := make(chan bool)
	callback := func(ip string, port int) {
		calledCh <- true
		assert.Equal(t, ip, "0.0.0.0")
		assert.Equal(t, port, 2502)
	}

	sender := AnimationSender{
		Address: "0.0.0.0",
		Port:    2502,
	}

	sender.SetOnUnableToConnectCallback(callback)

	sender.Start()

	called := <-calledCh
	assert.True(t, called)
}

func TestAnimationSender_Start_Already_Started(t *testing.T) {
	sender := AnimationSender{
		Address: "0.0.0.0",
		Port:    5,
		Started: true,
	}

	sender.Start()

	assert.Nil(t, sender.RunningAnimations)
	assert.Nil(t, sender.Sections)
	assert.Nil(t, sender.SupportedAnimations)
}

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

func TestAnimationSender_SetOnReceiveCallback(t *testing.T) {
	called := false
	callback := func(data string) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:   NewRunningAnimationMap(),
		Sections:            map[string]*section{},
		SupportedAnimations: map[string]*animationInfo{},
	}

	sender.SetOnReceiveCallback(callback)

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

func TestAnimationSender_SetOnNewAnimationDataCallback(t *testing.T) {
	called := false
	callback := func(data *animationData) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:   NewRunningAnimationMap(),
		Sections:            map[string]*section{},
		SupportedAnimations: map[string]*animationInfo{},
	}

	sender.SetOnNewAnimationDataCallback(callback)

	jsonStr := `DATA:{"id":"test"};;;`
	sender.processData([]byte(jsonStr))

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

func TestAnimationSender_SetOnNewAnimationInfoCallback(t *testing.T) {
	t.Skip()
	called := false
	callback := func(info *animationInfo) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:   NewRunningAnimationMap(),
		Sections:            map[string]*section{},
		SupportedAnimations: map[string]*animationInfo{},
	}

	sender.SetOnNewAnimationInfoCallback(callback)

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

func TestAnimationSender_SetOnNewEndAnimationCallback(t *testing.T) {
	called := false
	callback := func(end *endAnimation) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:   NewRunningAnimationMap(),
		Sections:            map[string]*section{},
		SupportedAnimations: map[string]*animationInfo{},
	}

	sender.SetOnNewEndAnimationCallback(callback)

	jsonStr := `END :{"id":"test"};;;`
	sender.processData([]byte(jsonStr))

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

func TestAnimationSender_SetOnNewMessageCallback(t *testing.T) {
	called := false
	callback := func(msg *message) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:   NewRunningAnimationMap(),
		Sections:            map[string]*section{},
		SupportedAnimations: map[string]*animationInfo{},
	}

	sender.SetOnNewMessageCallback(callback)

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

func TestAnimationSender_SetOnNewSectionCallback(t *testing.T) {
	called := false
	callback := func(sect *section) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:   NewRunningAnimationMap(),
		Sections:            map[string]*section{},
		SupportedAnimations: map[string]*animationInfo{},
	}

	sender.SetOnNewSectionCallback(callback)

	jsonStr := "SECT:{};;;"
	sender.processData([]byte(jsonStr))

	assert.True(t, called)
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

func TestAnimationSender_SetOnNewStripInfoCallback(t *testing.T) {
	called := false
	callback := func(info *stripInfo) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:   NewRunningAnimationMap(),
		Sections:            map[string]*section{},
		SupportedAnimations: map[string]*animationInfo{},
	}

	sender.SetOnNewStripInfoCallback(callback)

	jsonStr := "SINF:{};;;"
	sender.processData([]byte(jsonStr))

	assert.True(t, called)
}
