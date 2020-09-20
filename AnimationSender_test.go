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
	"bytes"
	"github.com/stretchr/testify/assert"
	"go.uber.org/atomic"
	"log"
	"net"
	"os"
	"testing"
	"time"
)

func TestAnimationSender_Start_End(t *testing.T) {
	calledCh1 := make(chan bool)
	callback := func(ip string, port int) {
		calledCh1 <- true
		assert.Equal(t, ip, "0.0.0.0")
		assert.Equal(t, port, 1101)
	}
	calledCh2 := make(chan bool)
	callback2 := func(ip string, port int) {
		calledCh2 <- true
		assert.Equal(t, ip, "0.0.0.0")
		assert.Equal(t, port, 1101)
	}

	ln, err := net.Listen("tcp", ":1101")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
		return
	}

	go func() {
		_, err2 := ln.Accept()
		assert.Nil(t, err2)
	}()

	sender := AnimationSender{
		Address: "0.0.0.0",
		Port:    1101,
	}

	sender.SetOnConnectCallback(callback)
	sender.SetOnDisconnectCallback(callback2)

	sender.Start()
	called1 := <-calledCh1
	assert.True(t, called1)
	assert.True(t, sender.IsStarted())
	assert.True(t, sender.IsConnected())
	d, _ := time.ParseDuration("5ms")
	time.Sleep(d)
	sender.End()

	called2 := <-calledCh2
	assert.True(t, called2)
	assert.False(t, sender.IsStarted())
	assert.False(t, sender.IsConnected())
	err = ln.Close()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestAnimationSender_Start_UnableToConnect(t *testing.T) {
	calledCh := make(chan bool)
	callback := func(ip string, port int) {
		calledCh <- true
		assert.Equal(t, ip, "0.0.0.0")
		assert.Equal(t, port, 1102)
	}

	sender := AnimationSender{
		Address: "0.0.0.0",
		Port:    1102,
	}

	sender.SetOnUnableToConnectCallback(callback)

	sender.Start()

	called := <-calledCh
	assert.True(t, called)
}

func TestAnimationSender_Start_Already_Started(t *testing.T) {
	sender := AnimationSender{
		Address: "0.0.0.0",
		Port:    1103,
		started: *atomic.NewBool(true),
	}

	sender.Start()

	assert.Nil(t, sender.RunningAnimations)
	assert.Nil(t, sender.Sections)
	assert.Nil(t, sender.SupportedAnimations)
}

func TestAnimationSender_End_Not_Connected(t *testing.T) {
	sender := AnimationSender{
		Address: "0.0.0.0",
		Port:    1104,
		started: *atomic.NewBool(true),
	}

	sender.End()

	assert.True(t, sender.IsStarted())
}

func TestAnimationSender_processData_partialData(t *testing.T) {
	called := false
	callback := func(data string) {
		called = true
		assert.Equal(t, data, "DATA:{}")
	}

	sender := AnimationSender{
		RunningAnimations: NewRunningAnimationMap(),
		onReceiveCallback: callback,
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
		RunningAnimations: NewRunningAnimationMap(),
		onReceiveCallback: callback,
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
		RunningAnimations: NewRunningAnimationMap(),
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
		onNewAnimationDataCallback: callback,
	}

	jsonStr := `DATA:{"id":"test"};;;`
	sender.processData([]byte(jsonStr))

	_, ok := sender.RunningAnimations.Load("test")

	assert.True(t, ok)
	assert.True(t, called)
}

func TestAnimationSender_processData_AnimationData_err(t *testing.T) {
	called := false
	callback := func(data *animationData) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:          NewRunningAnimationMap(),
		onNewAnimationDataCallback: callback,
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	jsonStr := `DATA:{"id":"test";;;`
	sender.processData([]byte(jsonStr))

	assert.False(t, called)
	assert.Equal(t, buf.String()[20:], "unexpected end of JSON input\n")
}

func TestAnimationSender_SetOnNewAnimationDataCallback(t *testing.T) {
	called := false
	callback := func(data *animationData) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations: NewRunningAnimationMap(),
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

func TestAnimationSender_processData_Command_err(t *testing.T) {
	sender := AnimationSender{}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	jsonStr := `CMD :{};;;`
	sender.processData([]byte(jsonStr))

	assert.Equal(t, buf.String()[20:], "WARNING: Receiving Command is not supported by client\n")
}

func TestAnimationSender_processData_EndAnimation(t *testing.T) {
	called := false
	callback := func(end *endAnimation) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:         NewRunningAnimationMap(),
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

func TestAnimationSender_processData_EndAnimation_err(t *testing.T) {
	called := false
	callback := func(end *endAnimation) {
		called = true
	}

	sender := AnimationSender{
		onNewEndAnimationCallback: callback,
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	jsonStr := `END :{"id":"test";;;`
	sender.processData([]byte(jsonStr))

	assert.False(t, called)
	assert.Equal(t, buf.String()[20:], "unexpected end of JSON input\n")
}

func TestAnimationSender_SetOnNewEndAnimationCallback(t *testing.T) {
	called := false
	callback := func(end *endAnimation) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations: NewRunningAnimationMap(),
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
		onNewMessageCallback: callback,
	}

	jsonStr := "MSG :{};;;"
	sender.processData([]byte(jsonStr))

	assert.True(t, called)
}

func TestAnimationSender_processData_Message_err(t *testing.T) {
	called := false
	callback := func(msg *message) {
		called = true
	}

	sender := AnimationSender{
		onNewMessageCallback: callback,
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	jsonStr := "MSG :{;;;"
	sender.processData([]byte(jsonStr))

	assert.False(t, called)
	assert.Equal(t, buf.String()[20:], "unexpected end of JSON input\n")
}

func TestAnimationSender_SetOnNewMessageCallback(t *testing.T) {
	called := false
	callback := func(msg *message) {
		called = true
	}

	sender := AnimationSender{}
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
		Sections:             map[string]*section{},
		onNewSectionCallback: callback,
	}

	jsonStr := "SECT:{};;;"
	sender.processData([]byte(jsonStr))

	assert.True(t, called)
	assert.Len(t, sender.Sections, 1)
}

func TestAnimationSender_processData_Section_err(t *testing.T) {
	called := false
	callback := func(sect *section) {
		called = true
	}

	sender := AnimationSender{
		Sections:             map[string]*section{},
		onNewSectionCallback: callback,
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	jsonStr := "SECT:{;;;"
	sender.processData([]byte(jsonStr))

	assert.False(t, called)
	assert.Len(t, sender.Sections, 0)
	assert.Equal(t, buf.String()[20:], "unexpected end of JSON input\n")
}

func TestAnimationSender_SetOnNewSectionCallback(t *testing.T) {
	called := false
	callback := func(sect *section) {
		called = true
	}

	sender := AnimationSender{
		Sections: map[string]*section{},
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
		onNewStripInfoCallback: callback,
	}

	jsonStr := "SINF:{};;;"
	sender.processData([]byte(jsonStr))

	assert.True(t, called)
	assert.NotNil(t, sender.StripInfo)
}

func TestAnimationSender_processData_StripInfo_err(t *testing.T) {
	called := false
	callback := func(info *stripInfo) {
		called = true
	}

	sender := AnimationSender{
		onNewStripInfoCallback: callback,
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	jsonStr := "SINF:{;;;"
	sender.processData([]byte(jsonStr))

	assert.False(t, called)
	assert.Nil(t, sender.StripInfo)
	assert.Equal(t, buf.String()[20:], "unexpected end of JSON input\n")
}

func TestAnimationSender_SetOnNewStripInfoCallback(t *testing.T) {
	called := false
	callback := func(info *stripInfo) {
		called = true
	}

	sender := AnimationSender{}
	sender.SetOnNewStripInfoCallback(callback)

	jsonStr := "SINF:{};;;"
	sender.processData([]byte(jsonStr))

	assert.True(t, called)
}

func TestAnimationSender_processData_type_err(t *testing.T) {
	sender := AnimationSender{}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	jsonStr := `TADA:{};;;`
	sender.processData([]byte(jsonStr))

	assert.Equal(t, buf.String()[20:], "WARNING: Unrecognized data type: TADA\n")
}
