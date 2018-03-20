package main

import (
	"fmt"
	"math"
	"time"
)

/**
 * 10001st prime
 *
 * https://projecteuler.net/problem=7
 * By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
 * What is the 10 001st prime number?
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%207
 * 素数を小さい方から6つ並べると 2, 3, 5, 7, 11, 13 であり, 6番目の素数は 13 である.
 * 10 001 番目の素数を求めよ.
 *
 * Contents of Project Euler are licenced under a Creative Commons Licence: Attribution-NonCommercial-ShareAlike 2.0 UK: England & Wales.
 * http://creativecommons.org/licenses/by-nc-sa/2.0/uk/
 */

func GenPrime0(cancel <-chan struct{}) <-chan int64 {
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

func GenPrime1(cancel <-chan struct{}) <-chan int64 {
	ch := make(chan int64)
	go func() {
		defer close(ch)
		ch <- 2
		for n := int64(3); ; n += 2 {
			select {
			case <-cancel:
				return
			default:
				if n < 4 {
					ch <- n
				} else if n < 9 {
					ch <- n
				} else if n%3 != 0 {
					r := int64(math.Sqrt(float64(n)))
					flag := true
					for f := int64(5); f <= r; f += 6 {
						if n%f == 0 || n%(f+2) == 0 {
							flag = false
							break
						}
					}
					if flag {
						ch <- n
					}
				}
			}
		}
	}()
	return ch
}

func answer(order int64, genPrime func(<-chan struct{}) <-chan int64) (int64, time.Duration) {
	cancel := make(chan struct{}, 1)
	pch := genPrime(cancel)
	defer func() {
		cancel <- struct{}{}
		<-pch
	}()
	i := int64(1)
	var last int64
	start := time.Now() // Start
	for p := range pch {
		if i == order {
			last = p
			break
		}
		i++
	}
	goal := time.Now()
	return last, goal.Sub(start)
}

func main() {
	fmt.Println(answer(10001, GenPrime0))
	fmt.Println(answer(10001, GenPrime1))
	fmt.Println(answer(1000001, GenPrime0))
	fmt.Println(answer(1000001, GenPrime1))
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
