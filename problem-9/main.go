package main

import (
	"fmt"
	"math"
)

/**
 * Special Pythagorean triplet
 *
 * https://projecteuler.net/problem=9
 * A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
 *   a^2 + b^2 = c^2
 * For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
 * There exists exactly one Pythagorean triplet for which a + b + c = 1000.
 * Find the product abc.
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%209
 * ピタゴラス数(ピタゴラスの定理を満たす自然数)とは a < b < c で以下の式を満たす数の組である.
 *   a^2 + b^2 = c^2
 * 例えば, 3^2 + 4^2 = 9 + 16 = 25 = 5^2 である.
 * a + b + c = 1000 となるピタゴラスの三つ組が一つだけ存在する.
 * これらの積 abc を計算しなさい.
 *
 * Contents of Project Euler are licenced under a Creative Commons Licence: Attribution-NonCommercial-ShareAlike 2.0 UK: England & Wales.
 * http://creativecommons.org/licenses/by-nc-sa/2.0/uk/
 */

func answer0(s int64) (a, b, c, m int64) {
	for a = 1; a < (s-3)/3; a++ {
		for b = a + 1; b < (s-a-1)/2; b++ {
			c = 1000 - a - b
			if a*a+b*b == c*c {
				m = a * b * c
				return
			}
		}
	}
	return
}

func gcd(a, b int64) int64 {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func answer1(s int64) (a, b, c, mm int64) {
	// a = (m^1 - n^2)d, b = 2mnd, c = (m^2 + n^2)d
	// a + b + c = 2m(m + n)d = 2mkd
	// k = m + n; m < k < 2m and gcd(m,k) = 1
	s2 := s / 2
	mmax := int64(math.Ceil(math.Sqrt(float64(s2)))) - 1
	for m := int64(2); m <= mmax; m++ {
		if s2%m == 0 {
			sm := s2 / m
			for sm&0x01 == 0 { // removing all factors 2
				sm >>= 1
			}
			var k int64
			if m&0x01 != 0 { //odd
				k = m + 2
			} else { //even
				k = m + 1
			}
			for ; k < 2*m && k <= sm; k += 2 {
				if sm%k == 0 && gcd(k, m) == 1 {
					d := s2 / (k * m)
					n := k - m
					a = d * (m*m - n*n)
					b = 2 * d * m * n
					c = d * (m*m + n*n)
					mm = a * b * c
					return
				}
			}
		}
	}
	return
}

func main() {
	fmt.Println(answer0(1000))
	fmt.Println(answer1(1000))
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
