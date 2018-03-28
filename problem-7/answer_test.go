package problem7

import "testing"

func TestAnswer0(t *testing.T) {
	testCases := []struct {
		order, answer int64
	}{
		{order: 6, answer: 13},
		{order: 10001, answer: 104743},
	}

	for _, tc := range testCases {
		a := Answer(tc.order, genPrime0)
		if a != tc.answer {
			t.Errorf("Answer(%v, genPrime0) = %v, want %v.", tc.order, a, tc.answer)
		}
		a = Answer(tc.order, genPrime1)
		if a != tc.answer {
			t.Errorf("Answer(%v, genPrime1) = %v, want %v.", tc.order, a, tc.answer)
		}
	}
}

func BenchmarkAnswer0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Answer(1000001, genPrime0)
	}
}

func BenchmarkAnswer1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Answer(1000001, genPrime1)
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
