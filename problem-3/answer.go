package problem3

import (
	"math"
)

/**
 * Largest prime factor
 *
 * https://projecteuler.net/problem=3
 * The prime factors of 13195 are 5, 7, 13 and 29.
 * What is the largest prime factor of the number 600851475143 ?
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%203
 * 13195 の素因数は 5, 7, 13, 29 である.
 * 600851475143 の素因数のうち最大のものを求めよ.
 *
 * Contents of Project Euler are licenced under a Creative Commons Licence: Attribution-NonCommercial-ShareAlike 2.0 UK: England & Wales.
 * http://creativecommons.org/licenses/by-nc-sa/2.0/uk/
 */

//Answer0 returns answer to this problem
func Answer0(n int64) int64 {
	cancel := make(chan struct{}, 1)
	pch := genPrime(cancel)
	defer func() {
		cancel <- struct{}{}
		<-pch
	}()
	var max int64
	for p := range pch {
		if n < p {
			break
		}
		if n%p == 0 {
			//fmt.Println(p)
			max = p
			for {
				if n%p != 0 {
					break
				}
				n /= p
			}
		}
	}
	return max
}
func genPrime(cancel <-chan struct{}) <-chan int64 {
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
func Answer1(n int64) int64 {
	cancel := make(chan struct{}, 1)
	fch := genFactor(cancel)
	defer func() {
		cancel <- struct{}{}
		<-fch
	}()
	lastFactor := int64(1)
	maxFactor := int64(math.Sqrt(float64(n)))
	for factor := range fch {
		if n <= 1 || factor > maxFactor {
			break
		}
		if n%factor == 0 {
			//fmt.Println(factor)
			n /= factor
			for n%factor == 0 {
				n /= factor
			}
			lastFactor = factor
			maxFactor = int64(math.Sqrt(float64(n)))
		}
	}
	if n == 1 {
		return lastFactor
	}
	return n
}
func genFactor(cancel <-chan struct{}) <-chan int64 {
	ch := make(chan int64)
	go func() {
		defer close(ch)
		ch <- 2
		for n := int64(3); ; n += 2 {
			select {
			case <-cancel:
				return
			default:
				ch <- n
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
