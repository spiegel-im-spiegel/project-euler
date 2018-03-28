package problem10

import (
	"math"
)

/**
 * Summation of primes
 *
 * https://projecteuler.net/problem=10
 * The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
 * Find the sum of all the primes below two million.
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%2010
 * 10以下の素数の和は 2 + 3 + 5 + 7 = 17 である.
 * 200万以下の全ての素数の和を求めよ.
 *
 * Contents of Project Euler are licenced under a Creative Commons Licence: Attribution-NonCommercial-ShareAlike 2.0 UK: England & Wales.
 * http://creativecommons.org/licenses/by-nc-sa/2.0/uk/
 */

//Answer0 returns answer to this problem
func Answer0(limit int64) int64 {
	cancel := make(chan struct{}, 1)
	pch := genPrime0(cancel)
	defer func() {
		cancel <- struct{}{}
		<-pch
	}()
	sum := int64(0)
	for p := range pch {
		if p > limit {
			break
		}
		sum += p
	}
	return sum
}
func genPrime0(cancel <-chan struct{}) <-chan int64 {
	ch := make(chan int64)
	go func() {
		defer close(ch)
		primes := []int64{}
		ch <- 2
		for n := int64(3); ; n += 2 {
			select {
			case <-cancel:
				return
			default:
				if n < 9 {
					primes = append(primes, n)
					ch <- n
				} else if n == 11 || n == 13 {
					primes = append(primes, n)
					ch <- n
				} else if n > 13 {
					pflag := true
					rn := int64(math.Sqrt(float64(n)))
					for _, p := range primes {
						if p > rn {
							break
						} else if n%p == 0 {
							pflag = false
							break
						}
					}
					if pflag {
						primes = append(primes, n)
						ch <- n
					}
				}
			}
		}
	}()
	return ch
}

//Answer1 returns answer to this problem (refactoring version)
func Answer1(limit int64) int64 {
	sum := int64(0)
	pch := genPrime1(limit)
	for p := range pch {
		sum += p
	}
	return sum
}

//genPrime1 is the Sieve of Eratosthenes
func genPrime1(limit int64) <-chan int64 {
	sievebound := (limit - 1) / 2
	sieve := make([]bool, sievebound+1)
	sieve[0] = true
	for i := int64(1); i <= sievebound; i++ {
		sieve[i] = false
	}
	crosslimit := (int64(math.Sqrt(float64(limit))) - 1) / 2

	for i := int64(1); i <= crosslimit; i++ {
		if !sieve[i] { // 2*i+1 is prime, and mark multiples
			for j := 2 * i * (i + 1); j <= sievebound; j += 2*i + 1 {
				sieve[j] = true
			}
		}
	}

	ch := make(chan int64)
	go func() {
		defer close(ch)
		ch <- 2 // 2 is prime
		for i, sv := range sieve {
			if !sv {
				ch <- 2*int64(i) + 1
			}
		}
	}()
	return ch
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
