/*
 *  Modified code from
 *  https://gist.github.com/deckarep/45527da11a10ccb0ca127ca7a82027e2#file-regular_map-go
 *
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

import "sync"

type RunningAnimationMap struct {
	sync.RWMutex
	internal map[string]*animationData
}

func NewRunningAnimationMap() *RunningAnimationMap {
	return &RunningAnimationMap{
		internal: make(map[string]*animationData),
	}
}

func (rm *RunningAnimationMap) Load(key string) (value *animationData, ok bool) {
	rm.RLock()
	result, ok := rm.internal[key]
	rm.RUnlock()
	return result, ok
}

func (rm *RunningAnimationMap) Delete(key string) {
	rm.Lock()
	delete(rm.internal, key)
	rm.Unlock()
}

func (rm *RunningAnimationMap) Store(key string, value *animationData) {
	rm.Lock()
	rm.internal[key] = value
	rm.Unlock()
}

func (rm *RunningAnimationMap) Keys() []string {
	rm.RLock()
	keys := make([]string, 0, len(rm.internal))
	for key := range rm.internal {
		keys = append(keys, key)
	}
	rm.RUnlock()
	return keys
}
