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
	"log"
	"net"
	"strconv"
	"strings"
)

type AnimationSender struct {
	Address    string
	Port       int
	connection *net.Conn

	Started   bool
	Connected bool

	RunningAnimations   *RunningAnimationMap
	Sections            map[string]*section
	SupportedAnimations map[string]*animationInfo
	StripInfo           *stripInfo

	onConnectCallback         func(string, int)
	onDisconnectCallback      func(string, int)
	onUnableToConnectCallback func(string, int)

	onReceiveCallback          func(string)
	onNewAnimationDataCallback func(*animationData)
	onNewAnimationInfoCallback func(*animationInfo)
	onNewEndAnimationCallback  func(*endAnimation)
	onNewMessageCallback       func(*message)
	onNewSectionCallback       func(*section)
	onNewStripInfoCallback     func(*stripInfo)

	partialData string
}

func (s *AnimationSender) Start() {
	if s.Started {
		return
	}

	s.RunningAnimations = NewRunningAnimationMap()
	s.Sections = map[string]*section{}
	s.SupportedAnimations = map[string]*animationInfo{}
	s.StripInfo = nil

	s.Started = true

	conn, err := net.Dial("tcp", s.Address+":"+strconv.Itoa(s.Port))
	if err != nil {
		s.onUnableToConnectCallback(s.Address, s.Port)
		s.Started = false
		s.Connected = false
	} else {
		s.connection = &conn

		s.Connected = true
		s.onConnectCallback(s.Address, s.Port)

		go s.receiverLoop()
	}
}

func (s *AnimationSender) End() {
	if !s.Connected {
		return
	}
	s.Started = false
	s.Connected = false
	_ = (*s.connection).Close()
}

func (s *AnimationSender) receiverLoop() {
	for s.Connected {
		buff := make([]byte, 16384)
		length, err := (*s.connection).Read(buff)
		if err != nil {
			s.Started = false
			s.Connected = false
			_ = (*s.connection).Close()
			s.onDisconnectCallback(s.Address, s.Port)
			return
		}

		if length > 0 {
			input := string(buff)
			tokens := strings.Split(s.partialData+input, ";;;")
			s.partialData = ""
			if !strings.HasSuffix(input, ";;;") {
				s.partialData = tokens[len(tokens)-1]
				tokens = tokens[:len(tokens)-1]
			}
			for i := 0; i < len(tokens); i++ {
				token := tokens[i]
				if len(token) == 0 {
					continue
				}

				if strings.HasPrefix(token, "DATA:") {
					anim, err := AnimationDataFromJson(token)
					if err != nil {
						log.Print(err.Error())
					} else {
						s.onNewAnimationDataCallback(anim)
						s.RunningAnimations.Store(anim.Id, anim)
					}
				} else if strings.HasPrefix(token, "AINF:") {
					info, err := AnimationInfoFromJson(token)
					if err != nil {
						log.Print(err.Error())
					} else {
						s.SupportedAnimations[info.Name] = info
						s.onNewAnimationInfoCallback(info)
					}
				} else if strings.HasPrefix(token, "CMD :") {
					log.Print("WARNING: Receiving Command is not supported by client")
				} else if strings.HasPrefix(token, "END :") {
					anim, err := EndAnimationFromJson(token)
					if err != nil {
						log.Print(err.Error())
					} else {
						s.onNewEndAnimationCallback(anim)
						s.RunningAnimations.Delete(anim.Id)
					}
				} else if strings.HasPrefix(token, "MSG :") {
					msg, err := MessageFromJson(token)
					if err != nil {
						log.Print(err.Error())
					} else {
						s.onNewMessageCallback(msg)
					}
				} else if strings.HasPrefix(token, "SECT:") {
					sect, err := SectionFromJson(token)
					if err != nil {
						log.Print(err.Error())
					} else {
						s.onNewSectionCallback(sect)
						s.Sections[sect.Name] = sect
					}
				} else if strings.HasPrefix(token, "SINF") {
					info, err := StripInfoFromJson(token)
					if err != nil {
						log.Print(err.Error())
					} else {
						s.StripInfo = info
						s.onNewStripInfoCallback(info)
					}
				} else {
					log.Print("WARNING: Unrecognized data type: " + token[:4])
				}
			}
		}
	}
}

func (s *AnimationSender) send(jsonBytes []byte) {
	jsonBytes = append(jsonBytes, []byte(";;;")...)
	_, err := (*s.connection).Write(jsonBytes)
	if err != nil {
		log.Print("ERROR: error sending")
	}
}

func (s *AnimationSender) SendAnimationData(data *animationData) {
	s.send(data.Json())
}

func (s *AnimationSender) SendCommand(cmd *Command) {
	s.send(cmd.Json())
}

func (s *AnimationSender) SendEndAnimation(endAnim *endAnimation) {
	s.send(endAnim.Json())
}

func (s *AnimationSender) SendSection(sect *section) {
	s.send(sect.Json())
}

func (s *AnimationSender) SetOnConnectCallback(action func(string, int)) {
	s.onConnectCallback = action
}

func (s *AnimationSender) SetOnDisconnectCallback(action func(string, int)) {
	s.onDisconnectCallback = action
}

func (s *AnimationSender) SetOnUnableToConnectCallback(action func(string, int)) {
	s.onUnableToConnectCallback = action
}

func (s *AnimationSender) SetOnNewAnimationDataCallback(action func(*animationData)) {
	s.onNewAnimationDataCallback = action
}

func (s *AnimationSender) SetOnNewAnimationInfoCallback(action func(*animationInfo)) {
	s.onNewAnimationInfoCallback = action
}

func (s *AnimationSender) SetOnNewEndAnimationCallback(action func(*endAnimation)) {
	s.onNewEndAnimationCallback = action
}

func (s *AnimationSender) SetOnNewMessageCallback(action func(*message)) {
	s.onNewMessageCallback = action
}

func (s *AnimationSender) SetOnNewSectionCallback(action func(*section)) {
	s.onNewSectionCallback = action
}

func (s *AnimationSender) SetOnNewStripInfoCallback(action func(*stripInfo)) {
	s.onNewStripInfoCallback = action
}
