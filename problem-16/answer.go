package problem12

import (
	"math/big"
)

/**
 * Power digit sum
 *
 * https://projecteuler.net/problem=16
 * 2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.
 * What is the sum of the digits of the number 2^1000?
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%2016
 * 2^15 = 32768 であり, 各位の数字の和は 3 + 2 + 7 + 6 + 8 = 26 となる.
 * 同様にして, 2^1000 の各位の数字の和を求めよ.
 *
 * Contents of Project Euler are licenced under a Creative Commons Licence: Attribution-NonCommercial-ShareAlike 2.0 UK: England & Wales.
 * http://creativecommons.org/licenses/by-nc-sa/2.0/uk/
 */

//Answer0 returns answer to this problem
func Answer0(p int64) int64 {
	pow := (&big.Int{}).Exp(big.NewInt(2), big.NewInt(p), big.NewInt(0))
	//fmt.Println(pow)
	ten := big.NewInt(10)
	dig := big.NewInt(0)
	sum := int64(0)
	for {
		pow, dig = (&big.Int{}).DivMod(pow, ten, dig)
		sum += dig.Int64()
		if pow.Sign() == 0 {
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
