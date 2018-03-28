package problem8

import "testing"

func TestAnswer0(t *testing.T) {
	testCases := []struct {
		digit  int
		answer int64
	}{
		{digit: 4, answer: 5832},
		{digit: 13, answer: 23514624000},
	}

	for _, tc := range testCases {
		a := Answer0(tc.digit)
		if a != tc.answer {
			t.Errorf("Answer0(%v) = %v, want %v.", tc.digit, a, tc.answer)
		}
		a = Answer1(tc.digit)
		if a != tc.answer {
			t.Errorf("Answer1(%v) = %v, want %v.", tc.digit, a, tc.answer)
		}
	}
}

func BenchmarkAnswer0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Answer0(3)
	}
}

func BenchmarkAnswer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Answer1(3)
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
