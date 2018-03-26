package problem4

import "testing"

func TestAnswer(t *testing.T) {
	testCases := []struct {
		min, max, answer int
	}{
		{min: 10, max: 99, answer: 9009},
		{min: 100, max: 999, answer: 906609},
	}

	for i, tc := range testCases {
		a := Answer0(tc.min, tc.max)
		if a != tc.answer {
			t.Errorf("Answer0(%v, %v) = %v, want %v.", tc.min, tc.max, a, tc.answer)
		}
		if i == 1 {
			a := Answer1()
			if a != tc.answer {
				t.Errorf("Answer1(%v, %v) = %v, want %v.", tc.min, tc.max, a, tc.answer)
			}
		}
	}
}

func BenchmarkAnswer0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Answer0(100, 999)
	}
}

func BenchmarkAnswer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Answer1()
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
