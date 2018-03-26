package problem6

import "testing"

func TestAnswer0(t *testing.T) {
	testCases := []struct {
		max, answer int64
	}{
		{max: 10, answer: 2640},
		{max: 100, answer: 25164150},
	}

	for _, tc := range testCases {
		a := Answer0(tc.max)
		if a != tc.answer {
			t.Errorf("Answer0(%v) = %v, want %v.", tc.max, a, tc.answer)
		}
		a = Answer1(tc.max)
		if a != tc.answer {
			t.Errorf("Answer1(%v) = %v, want %v.", tc.max, a, tc.answer)
		}
	}
}

func BenchmarkAnswer0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Answer0(100)
	}
}

func BenchmarkAnswer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Answer1(100)
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
