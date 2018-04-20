package problem14

/**
 * Longest Collatz sequence
 *
 * https://projecteuler.net/problem=14
 * The following iterative sequence is defined for the set of positive integers:
 *   n → n/2 (n is even)
 *   n → 3n + 1 (n is odd)
 * Using the rule above and starting with 13, we generate the following sequence:
 *   13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
 * It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
 * Which starting number, under one million, produces the longest chain?
 * NOTE: Once the chain starts the terms are allowed to go above one million.
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%2014
 * 正の整数に以下の式で繰り返し生成する数列を定義する.
 *   n → n/2 (n is even)
 *   n → 3n + 1 (n is odd)
 * 13からはじめるとこの数列は以下のようになる.
 *   13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
 * 13から1まで10個の項になる. この数列はどのような数字からはじめても最終的には 1 になると考えられているが, まだそのことは証明されていない(コラッツ問題)
 * さて, 100万未満の数字の中でどの数字からはじめれば最長の数列を生成するか.
 * 注意: 数列の途中で100万以上になってもよい
 *
 * Contents of Project Euler are licenced under a Creative Commons Licence: Attribution-NonCommercial-ShareAlike 2.0 UK: England & Wales.
 * http://creativecommons.org/licenses/by-nc-sa/2.0/uk/
 */

func nextNumber(n int64) int64 {
	if n&0x01 == 0 { //even
		n >>= 1
	} else { //odd
		n = 3*n + 1
	}
	return n
}

//Answer0 returns answer to this problem
func Answer0(max int64) int64 {
	maxCount := int64(0)
	cn := int64(0)
	for i := int64(3); i < max; i++ {
		ct := int64(0)
		for n := i; n > 1; n = nextNumber(n) {
			//fmt.Printf("%d -> ", n)
			ct++
		}
		//fmt.Println(1)
		if maxCount < ct {
			maxCount = ct
			cn = i
		}
	}
	return cn
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
