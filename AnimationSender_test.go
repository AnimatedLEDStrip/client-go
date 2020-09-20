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
	"sync"
	"testing"
	"time"
)

func TestAnimationSender(t *testing.T) {
	calledCh1 := make(chan bool)
	callback1 := func(ip string, port int) {
		calledCh1 <- true
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

	_ = <-readyCh

	_, _ = conn.Write([]byte(`DATA:{"animation":"Meteor","center":50,"continuous":false,"delay":10,"delayMod":1.5,"direction":"BACKWARD","distance":45,"id":"TEST","section":"SECT","spacing":5};;;`))

	called2 := <-calledCh2
	assert.True(t, called2)

	d, _ := time.ParseDuration("5ms")
	time.Sleep(d)

	sender.End()

	called3 := <-calledCh3
	assert.True(t, called3)

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
		assert.Equal(t, "0.0.0.0", ip)
		assert.Equal(t, 1102, port)
	}
	calledCh2 := make(chan bool)
	callback2 := func(ip string, port int) {
		calledCh2 <- true
		assert.Equal(t, "0.0.0.0", ip)
		assert.Equal(t, 1102, port)
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
		assert.Equal(t, "0.0.0.0", ip)
		assert.Equal(t, 1103, port)
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
		RunningAnimations:          new(sync.Map),
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
		assert.Equal(t, "DATA:{}", data)
	}

	sender := AnimationSender{
		RunningAnimations: new(sync.Map),
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
		assert.Equal(t, "DATA:{}", data)
	}

	sender := AnimationSender{
		RunningAnimations: new(sync.Map),
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
		RunningAnimations: new(sync.Map),
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
		RunningAnimations:          new(sync.Map),
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
		RunningAnimations:          new(sync.Map),
		onNewAnimationDataCallback: callback,
	}

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	sender.processData([]byte(`DATA:{"id":"test";;;`))

	assert.False(t, called)
	assert.Equal(t, "unexpected end of JSON input\n", buf.String()[20:])
}

func TestAnimationSender_SetOnNewAnimationDataCallback(t *testing.T) {
	called := false
	callback := func(data *animationData) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations: new(sync.Map),
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
	assert.Equal(t, "unexpected end of JSON input\n", buf.String()[20:])
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

	assert.Equal(t, "WARNING: Receiving Command is not supported by client\n", buf.String()[20:])
}

func TestAnimationSender_processData_EndAnimation(t *testing.T) {
	called := false
	callback := func(end *endAnimation) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations:         new(sync.Map),
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
	assert.Equal(t, "unexpected end of JSON input\n", buf.String()[20:])
}

func TestAnimationSender_SetOnNewEndAnimationCallback(t *testing.T) {
	called := false
	callback := func(end *endAnimation) {
		called = true
	}

	sender := AnimationSender{
		RunningAnimations: new(sync.Map),
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
	assert.Equal(t, "unexpected end of JSON input\n", buf.String()[20:])
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
	assert.Equal(t, "unexpected end of JSON input\n", buf.String()[20:])
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
	assert.Equal(t, "unexpected end of JSON input\n", buf.String()[20:])
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

	assert.Equal(t, "WARNING: Unrecognized data type: TADA\n", buf.String()[20:])
}

func TestAnimationSender_SendAnimationData(t *testing.T) {
	calledCh1 := make(chan bool)
	callback1 := func(ip string, port int) {
		calledCh1 <- true
	}
	calledCh2 := make(chan bool)
	callback2 := func(ip string, port int) {
		calledCh2 <- true
	}

	ln, err := net.Listen("tcp", ":1106")
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
		Port:    1106,
	}

	sender.SetOnConnectCallback(callback1)
	sender.SetOnDisconnectCallback(callback2)

	sender.Start()
	called1 := <-calledCh1
	assert.True(t, called1)

	_ = <-readyCh

	buff := make([]byte, 16384)
	var length int

	go func() {
		length, _ = conn.Read(buff)
		assert.Greater(t, length, 0)
		readyCh <- true
	}()

	data := AnimationData()
	data.SetAnimation("Meteor")
	data.SetCenter(50)
	data.SetContinuous(NONCONTINUOUS)
	data.SetDelay(10)
	data.SetDelayMod(1.5)
	data.SetDirection(BACKWARD)
	data.SetDistance(45)
	data.SetID("TEST")
	data.SetSection("SECT")
	data.SetSpacing(5)

	cc := ColorContainer{}
	cc.AddColor(0xFF).AddColor(0xFF00)
	cc2 := ColorContainer{}
	cc2.AddColor(0xFF0000)
	data.AddColor(&cc)
	data.AddColor(&cc2)

	sender.SendAnimationData(data)

	_ = <-readyCh

	assert.Equal(t,
		`DATA:{"animation":"Meteor","center":50,"colors":[{"colors":[255,65280]},{"colors":[16711680]}],"continuous":false,"delay":10,"delayMod":1.5,"direction":"BACKWARD","distance":45,"id":"TEST","section":"SECT","spacing":5};;;`,
		string(buff)[:length])

	d, _ := time.ParseDuration("5ms")
	time.Sleep(d)

	sender.End()

	called2 := <-calledCh2
	assert.True(t, called2)

	err = ln.Close()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestAnimationSender_SendCommand(t *testing.T) {
	calledCh1 := make(chan bool)
	callback1 := func(ip string, port int) {
		calledCh1 <- true
	}
	calledCh2 := make(chan bool)
	callback2 := func(ip string, port int) {
		calledCh2 <- true
	}

	ln, err := net.Listen("tcp", ":1107")
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
		Port:    1107,
	}

	sender.SetOnConnectCallback(callback1)
	sender.SetOnDisconnectCallback(callback2)

	sender.Start()
	called1 := <-calledCh1
	assert.True(t, called1)

	_ = <-readyCh

	buff := make([]byte, 16384)
	var length int

	go func() {
		length, _ = conn.Read(buff)
		assert.Greater(t, length, 0)
		readyCh <- true
	}()

	cmd := Command{Cmd: "A command"}

	sender.SendCommand(&cmd)

	_ = <-readyCh

	assert.Equal(t, `CMD :{"command":"A command"};;;`, string(buff)[:length])

	d, _ := time.ParseDuration("5ms")
	time.Sleep(d)

	sender.End()

	called2 := <-calledCh2
	assert.True(t, called2)

	err = ln.Close()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestAnimationSender_SendEndAnimation(t *testing.T) {
	calledCh1 := make(chan bool)
	callback1 := func(ip string, port int) {
		calledCh1 <- true
	}
	calledCh2 := make(chan bool)
	callback2 := func(ip string, port int) {
		calledCh2 <- true
	}

	ln, err := net.Listen("tcp", ":1108")
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
		Port:    1108,
	}

	sender.SetOnConnectCallback(callback1)
	sender.SetOnDisconnectCallback(callback2)

	sender.Start()
	called1 := <-calledCh1
	assert.True(t, called1)

	_ = <-readyCh

	buff := make([]byte, 16384)
	var length int

	go func() {
		length, _ = conn.Read(buff)
		assert.Greater(t, length, 0)
		readyCh <- true
	}()

	anim := EndAnimation()
	anim.SetId("54321")

	sender.SendEndAnimation(anim)

	_ = <-readyCh

	assert.Equal(t, `END :{"id":"54321"};;;`, string(buff)[:length])

	d, _ := time.ParseDuration("5ms")
	time.Sleep(d)

	sender.End()

	called2 := <-calledCh2
	assert.True(t, called2)

	err = ln.Close()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestAnimationSender_SendSection(t *testing.T) {
	calledCh1 := make(chan bool)
	callback1 := func(ip string, port int) {
		calledCh1 <- true
	}
	calledCh2 := make(chan bool)
	callback2 := func(ip string, port int) {
		calledCh2 <- true
	}

	ln, err := net.Listen("tcp", ":1109")
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
		Port:    1109,
	}

	sender.SetOnConnectCallback(callback1)
	sender.SetOnDisconnectCallback(callback2)

	sender.Start()
	called1 := <-calledCh1
	assert.True(t, called1)

	_ = <-readyCh

	buff := make([]byte, 16384)
	var length int

	go func() {
		length, _ = conn.Read(buff)
		assert.Greater(t, length, 0)
		readyCh <- true
	}()

	sect := Section()
	sect.SetName("Section")
	sect.SetStartPixel(30)
	sect.SetEndPixel(40)

	sender.SendSection(sect)

	_ = <-readyCh

	assert.Equal(t, `SECT:{"name":"Section","startPixel":30,"endPixel":40};;;`, string(buff)[:length])

	d, _ := time.ParseDuration("5ms")
	time.Sleep(d)

	sender.End()

	called2 := <-calledCh2
	assert.True(t, called2)

	err = ln.Close()
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	}
}

func TestAnimationSender_send_not_connected(t *testing.T) {
	sender := AnimationSender{
		Address: "0.0.0.0",
		Port:    1110,
	}

	sect := Section()
	sect.SetName("Section")
	sect.SetStartPixel(30)
	sect.SetEndPixel(40)

	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	sender.SendSection(sect)
	assert.Equal(t, "WARNING: Not connected\n", buf.String()[20:])
}
