package problem20

import (
	"math/big"
)

/**
 * Factorial digit sum
 *
 * https://projecteuler.net/problem=20
 * n! means n × (n − 1) × ... × 3 × 2 × 1
 * For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
 * and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.
 * Find the sum of the digits in the number 100!
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%2020
 * n × (n - 1) × ... × 3 × 2 × 1 を n! と表す.
 * 例えば, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800 となる.
 * この数の各桁の合計は 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27 である.
 * では, 100! の各位の数字の和を求めよ.
 *
 * Contents of Project Euler are licenced under a Creative Commons Licence: Attribution-NonCommercial-ShareAlike 2.0 UK: England & Wales.
 * http://creativecommons.org/licenses/by-nc-sa/2.0/uk/
 */

//Answer0 returns answer to this problem
func Answer0(f int64) int64 {
	fn := (&big.Int{}).MulRange(1, f)
	//fmt.Println(fn)
	ten := big.NewInt(10)
	dig := big.NewInt(0)
	sum := int64(0)
	for {
		fn, dig = (&big.Int{}).DivMod(fn, ten, dig)
		sum += dig.Int64()
		if fn.Sign() == 0 {
			break
		}
	}
	return sum
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
