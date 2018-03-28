package problem9

import "testing"

func TestAnswer0(t *testing.T) {
	testCases := []struct {
		total, answer int64
	}{
		{total: 1000, answer: 31875000},
	}

	for _, tc := range testCases {
		a := Answer0(tc.total)
		if a != tc.answer {
			t.Errorf("Answer0(%v) = %v, want %v.", tc.total, a, tc.answer)
		}
		a = Answer1(tc.total)
		if a != tc.answer {
			t.Errorf("Answer1(%v) = %v, want %v.", tc.total, a, tc.answer)
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
