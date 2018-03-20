package main

import (
	"fmt"
	"strconv"
)

/**
 * Largest palindrome product
 *
 * https://projecteuler.net/problem=4
 * A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
 * Find the largest palindrome made from the product of two 3-digit numbers.
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%204
 * 左右どちらから読んでも同じ値になる数を回文数という. 2桁の数の積で表される回文数のうち, 最大のものは 9009 = 91 × 99 である.
 * では, 3桁の数の積で表される回文数の最大値を求めよ.
 *
 * Contents of Project Euler are licenced under a Creative Commons Licence: Attribution-NonCommercial-ShareAlike 2.0 UK: England & Wales.
 * http://creativecommons.org/licenses/by-nc-sa/2.0/uk/
 */

func reverse0(n int) int {
	r := []rune(strconv.Itoa(n))
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	rn, _ := strconv.Atoi(string(r))
	return rn
}

func answer0() int {
	max := 0
	for n := 100; n <= 999; n++ {
		for m := n; m <= 999; m++ {
			nm := n * m
			if nm == reverse0(nm) && nm > max {
				max = nm
			}
		}
	}
	return max
}

func reverse1(n int) int {
	rn := 0
	for n > 0 {
		rn = 10*rn + n%10
		n /= 10
	}
	return rn
}
func answer1() int {
	// P = 100000x + 10000y + 1000z + 100z + 10y + x
	//   = 100001x + 10010y + 1100z
	//   = 11 * (9091x + 910y + 100z)
	max := 0
	for n := 999; n >= 100; n-- {
		m := 990 //The largest number less than or equal 999 and divisible by 11
		dm := 11
		if n%11 == 0 {
			m = 999
			dm = 1
		}
		for ; n <= m && n*m > max; m -= dm {
			nm := n * m
			if nm == reverse1(nm) {
				max = nm
			}
		}
	}
	return max
}

func main() {
	fmt.Println(answer0())
	fmt.Println(answer1())
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
