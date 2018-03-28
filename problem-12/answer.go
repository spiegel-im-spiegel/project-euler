package problem12

import "math"

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
	if limit == 0 {
		return 1
	}
	cancel := make(chan struct{}, 1)
	ch := genTriangularNumber(cancel)
	defer func() {
		cancel <- struct{}{}
		<-ch
	}()
	var tn int64 = 0
	for n := range ch {
		//fmt.Println(n)
		rn := int64(math.Sqrt(float64(n)))
		ct := int64(0)
		for i := int64(1); i <= rn; i++ {
			if n%i == 0 {
				ct += 2
			}
		}
		if rn*rn == n {
			ct--
		}
		if ct > limit {
			tn = n
			break
		}
	}
	return tn
}
func genTriangularNumber(cancel <-chan struct{}) <-chan int64 {
	ch := make(chan int64)
	go func() {
		defer close(ch)
		tn := int64(0)
		for n := int64(1); ; n++ {
			select {
			case <-cancel:
				return
			default:
				tn += n
				ch <- tn
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
