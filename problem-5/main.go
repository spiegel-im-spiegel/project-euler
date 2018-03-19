package main

import (
	"fmt"
	"math"
)

/**
 * Smallest multiple
 *
 * https://projecteuler.net/problem=5
 * 2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
 * What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%205
 * では, 1 から 20 までの整数全てで割り切れる数字の中で最小の正の数はいくらになるか.
 * では, 3桁の数の積で表される回文数の最大値を求めよ.
 */

func answer0(k int64) int64 {
	max := int64(1)
	for n := k; n > 1; n-- {
		if max%n != 0 {
			max *= n
		}
	}
	min := max
	for n := max; n > k; n-- {
		flag := true
		for m := k; m > 1; m-- {
			if n%m != 0 {
				flag = false
				break
			}
		}
		if flag {
			min = n
		}
	}
	return min
}

func primeList(k int64) []int64 {
	primes := []int64{2}
	for n := int64(3); n <= k; n += 2 {
		pflag := true
		maxPrime := int64(math.Sqrt(float64(n)))
		for _, p := range primes {
			if p > maxPrime {
				break
			} else if n%p == 0 {
				pflag = false
				break
			}
		}
		if pflag {
			primes = append(primes, n)
		}
	}
	return primes
}
func answer1(k int64) int64 {
	rk := int64(math.Sqrt(float64(k)))
	lk := math.Log(float64(k))
	flag := true
	plist := primeList(k)
	n := int64(1)
	for _, p := range plist {
		a := int64(1)
		if flag {
			if p <= rk {
				a = int64(math.Floor(lk / math.Log(float64(p))))
			} else {
				flag = false
			}
		}
		n *= int64(math.Pow(float64(p), float64(a)))
	}
	return n
}
func main() {
	fmt.Println(answer0(20))
	fmt.Println(answer1(20))
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
