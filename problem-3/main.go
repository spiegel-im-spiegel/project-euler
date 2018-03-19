package main

import (
	"fmt"
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
 */

func GenPrime() <-chan int64 {
	ch := make(chan int64)
	go func() {
		primes := []int64{2}
		primesf := []float64{2.0}
		ch <- 2
		for n := int64(3); ; n += 2 {
			pflag := true
			f := float64(n)
			for i, p := range primes {
				if primesf[i] > math.Sqrt(f) {
					break
				} else if (n % p) == 0 {
					pflag = false
					break
				}
			}
			if pflag {
				primes = append(primes, n)
				primesf = append(primesf, f)
				ch <- n
			}
		}
	}()
	return ch
}
func answer0(n int64) int64 {
	pch := GenPrime()
	var max int64 = 0
	for p := range pch {
		if n < p {
			break
		}
		if n%p == 0 {
			fmt.Println(p)
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

func GenFactor() <-chan int64 {
	ch := make(chan int64)
	go func() {
		ch <- 2
		for n := int64(3); ; n += 2 {
			ch <- n
		}
	}()
	return ch
}
func answer1(n int64) int64 {
	fch := GenFactor()
	lastFactor := int64(1)
	maxFactor := int64(math.Sqrt(float64(n)))
	for factor := range fch {
		if n <= 1 || factor > maxFactor {
			break
		}
		if n%factor == 0 {
			fmt.Println(factor)
			n /= factor
			for {
				if n%factor != 0 {
					break
				}
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

func main() {
	fmt.Println("Largest prime:", answer0(600851475143))
	fmt.Println("Largest prime:", answer1(600851475143))
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
