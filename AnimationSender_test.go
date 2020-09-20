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

func TestAnimationSender(t *testing.T) {
	calledCh1 := make(chan bool)
	callback1 := func(ip string, port int) {
		calledCh1 <- true
		assert.Equal(t, ip, "0.0.0.0")
		assert.Equal(t, port, 1101)
	}
	calledCh2 := make(chan bool)
	callback2 := func(data *animationData) {
		calledCh2 <- true
		actualData := animationData{
			Animation:  "Meteor",
			Center:     50,
			Colors:     nil,
			Continuous: NONCONTINUOUS,
			Delay:      10,
			DelayMod:   1.5,
			Direction:  BACKWARD,
			Distance:   45,
			Id:         "TEST",
			Section:    "SECT",
			Spacing:    5,
		}
		assert.Equal(t, &actualData, data)
	}
	calledCh3 := make(chan bool)
	callback3 := func(ip string, port int) {
		calledCh3 <- true
		assert.Equal(t, ip, "0.0.0.0")
		assert.Equal(t, 1101, port)
	}

	ln, err := net.Listen("tcp", ":1101")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
		return
	}

	var conn net.Conn

	readyCh := make(chan bool)

	go func() {
		conn, _ = ln.Accept()
		readyCh <- true
	}()

	sender := AnimationSender{
		Address: "0.0.0.0",
		Port:    1101,
	}

	sender.SetOnConnectCallback(callback1)
	sender.SetOnNewAnimationDataCallback(callback2)
	sender.SetOnDisconnectCallback(callback3)

	sender.Start()
	called1 := <-calledCh1
	assert.True(t, called1)
	assert.True(t, sender.IsStarted())
	assert.True(t, sender.IsConnected())

	_ = <-readyCh

	_, _ = conn.Write([]byte(`DATA:{"animation":"Meteor","center":50,"continuous":false,"delay":10,"delayMod":1.5,"direction":"BACKWARD","distance":45,"id":"TEST","section":"SECT","spacing":5};;;`))

	called2 := <-calledCh2
	assert.True(t, called2)

	d, _ := time.ParseDuration("5ms")
	time.Sleep(d)
	sender.End()

	called3 := <-calledCh3
	assert.True(t, called3)
	assert.False(t, sender.IsStarted())
	assert.False(t, sender.IsConnected())

	err = ln.Close()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestAnimationSender_Start_End(t *testing.T) {
	calledCh1 := make(chan bool)
	callback1 := func(ip string, port int) {
		calledCh1 <- true
		assert.Equal(t, ip, "0.0.0.0")
		assert.Equal(t, port, 1102)
	}
	calledCh2 := make(chan bool)
	callback2 := func(ip string, port int) {
		calledCh2 <- true
		assert.Equal(t, ip, "0.0.0.0")
		assert.Equal(t, port, 1102)
	}

	ln, err := net.Listen("tcp", ":1102")
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
		Port:    1102,
	}

	sender.SetOnConnectCallback(callback1)
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
		assert.Equal(t, port, 1103)
	}

	sender := AnimationSender{
		Address: "0.0.0.0",
		Port:    1103,
	}

	sender.SetOnUnableToConnectCallback(callback)

	sender.Start()

	called := <-calledCh
	assert.True(t, called)
}

func TestAnimationSender_Start_Already_Started(t *testing.T) {
	sender := AnimationSender{
		Address: "0.0.0.0",
		Port:    1104,
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
		Port:    1105,
		started: *atomic.NewBool(true),
	}

	sender.End()

	assert.True(t, sender.IsStarted())
}

func TestAnimationSender_receiveData(t *testing.T) {
	called := false
	callback := func(data *animationData) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:          NewRunningAnimationMap(),
		onNewAnimationDataCallback: callback,
	}

	sender.processData([]byte(`DATA:{"id":"test"};;;`))

	_, ok := sender.RunningAnimations.Load("test")

	assert.True(t, ok)
	assert.True(t, called)
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

	sender.processData([]byte("DAT"))
	sender.processData([]byte("A:{};;;"))

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

	sender.processData([]byte(`DATA:{};;;`))

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

	sender.processData([]byte(`DATA:{};;;`))

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

	sender.processData([]byte(`DATA:{"id":"test"};;;`))

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

	sender.processData([]byte(`DATA:{"id":"test";;;`))

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

	sender.processData([]byte(`DATA:{"id":"test"};;;`))

	assert.True(t, called)
}

func TestAnimationSender_processData_AnimationInfo(t *testing.T) {
	called := false
	callback := func(info *animationInfo) {
		called = true
	}

	sender := AnimationSender{
		SupportedAnimations:        map[string]*animationInfo{},
		onNewAnimationInfoCallback: callback,
	}

	sender.processData([]byte("AINF:{};;;"))

	assert.True(t, called)
}

func TestAnimationSender_processData_AnimationInfo_err(t *testing.T) {
	called := false
	callback := func(info *animationInfo) {
		called = true
	}

	sender := AnimationSender{
		SupportedAnimations:        map[string]*animationInfo{},
		onNewAnimationInfoCallback: callback,
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	sender.processData([]byte("AINF:{;;;"))

	assert.False(t, called)
	assert.Equal(t, buf.String()[20:], "unexpected end of JSON input\n")
}

func TestAnimationSender_SetOnNewAnimationInfoCallback(t *testing.T) {
	called := false
	callback := func(info *animationInfo) {
		called = true
	}

	sender := AnimationSender{
		SupportedAnimations: map[string]*animationInfo{},
	}

	sender.SetOnNewAnimationInfoCallback(callback)

	sender.processData([]byte("AINF:{};;;"))

	assert.True(t, called)
}

func TestAnimationSender_processData_Command_err(t *testing.T) {
	sender := AnimationSender{}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	sender.processData([]byte("CMD :{};;;"))

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

	sender.processData([]byte(`DATA:{"id":"test"};;;`))

	_, ok := sender.RunningAnimations.Load("test")
	assert.True(t, ok)

	sender.processData([]byte(`END :{"id":"test"};;;`))

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

	sender.processData([]byte(`END :{"id":"test";;;`))

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

	sender.processData([]byte(`END :{"id":"test"};;;`))

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

	sender.processData([]byte("MSG :{};;;"))

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

	sender.processData([]byte("MSG :{;;;"))

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

	sender.processData([]byte("MSG :{};;;"))

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

	sender.processData([]byte("SECT:{};;;"))

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

	sender.processData([]byte("SECT:{;;;"))

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

	sender.processData([]byte("SECT:{};;;"))

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

	sender.processData([]byte("SINF:{};;;"))

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

	sender.processData([]byte("SINF:{;;;"))

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

	sender.processData([]byte("SINF:{};;;"))

	assert.True(t, called)
}

func TestAnimationSender_processData_type_err(t *testing.T) {
	sender := AnimationSender{}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	sender.processData([]byte("TADA:{};;;"))

	assert.Equal(t, buf.String()[20:], "WARNING: Unrecognized data type: TADA\n")
}
