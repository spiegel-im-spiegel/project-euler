package problem2

import (
	"reflect"
	"runtime"
	"strings"
	"testing"
)

type Answer func(int64) int64

func TestAnswer0(t *testing.T) {
	testCases := []struct {
		max, answer int64
		fnc         Answer
	}{
		{max: 4000000, answer: 4613732, fnc: Answer0},
		{max: 4000000, answer: 4613732, fnc: Answer1},
	}

	for _, tc := range testCases {
		a := tc.fnc(tc.max)
		if a != tc.answer {
			fname := runtime.FuncForPC(reflect.ValueOf(tc.fnc).Pointer()).Name()
			fname = fname[strings.LastIndex(fname, ".")+1:]
			t.Errorf("%v(%v) = %v, want %v.", fname, tc.max, a, tc.answer)
		}
	}
}

func BenchmarkAnswer0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Answer0(1000)
	}
}

func BenchmarkAnswer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Answer1(1000)
	}
}

/* Copyright 2018 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
