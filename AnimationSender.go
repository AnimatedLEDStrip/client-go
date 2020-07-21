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
	"net"
	"strconv"
	"strings"
)

type AnimationSender struct {
	Ip                  string
	Port                int
	connection          *net.Conn
	RunningAnimations   *RunningAnimationMap
	Info                *stripInfo
	SupportedAnimations []*animationInfo
}

func (s *AnimationSender) addAnimation(data *animationData) {
	s.RunningAnimations.Store(data.Id, data)
}

func (s *AnimationSender) removeAnimation(data *animationData) {
	s.RunningAnimations.Delete(data.Id)
}

func (s *AnimationSender) connect() {
	conn, err := net.Dial("tcp", s.Ip+":"+strconv.Itoa(s.Port))
	if err != nil {
		print("Error connecting")
	} else {
		s.connection = &conn
	}
}

func (s *AnimationSender) receiverLoop() {
	for {
		buff := make([]byte, 16384)
		length, err := (*s.connection).Read(buff)
		if err != nil {
			_ = (*s.connection).Close()
			return
		}
		if length > 0 {
			tokens := strings.Split(string(buff), ";")
			for i := 0; i < len(tokens); i++ {
				token := tokens[i]
				if strings.HasPrefix(token, "DATA:") {
					anim := AnimationDataFromJson(token)
					s.RunningAnimations.Store(anim.Id, anim)
				} else if strings.HasPrefix(token, "AINF:") {
					info := AnimationInfoFromJson(token)
					s.SupportedAnimations = append(s.SupportedAnimations, info)
				} else if strings.HasPrefix(token, "END :") {
					anim := EndAnimationFromJson(token)
					s.RunningAnimations.Delete(anim.Id)
				} else if strings.HasPrefix(token, "SECT:") {
					// TODO
				} else if strings.HasPrefix(token, "SINF") {
					info := StripInfoFromJson(token)
					s.Info = info
				}
			}
		}
	}
}

func (s *AnimationSender) Start() {
	s.RunningAnimations = NewRunningAnimationMap()
	s.SupportedAnimations = []*animationInfo{}
	s.connect()
	go s.receiverLoop()
}

func (s *AnimationSender) End() {
	_ = (*s.connection).Close()
	s.RunningAnimations = nil
	s.SupportedAnimations = nil
}

func (s *AnimationSender) SendAnimation(data *animationData) {
	_, err := (*s.connection).Write([]byte(data.Json()))
	if err != nil {
		println("error sending")
	} else {
		println("sent")
	}
}
