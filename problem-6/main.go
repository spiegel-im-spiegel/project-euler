package main

import (
	"fmt"
)

/**
 * Sum square difference
 *
 * https://projecteuler.net/problem=6
 * The sum of the squares of the first ten natural numbers is,
 *   1^2 + 2^2 + ... + 10^2 = 385
 * The square of the sum of the first ten natural numbers is,
 *   (1 + 2 + ... + 10)^2 = 55^2 = 3025
 * Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.
 * Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%206
 * 最初の10個の自然数について, その二乗の和は,
 *   1^2 + 2^2 + ... + 10^2 = 385
 * 最初の10個の自然数について, その和の二乗は,
 *   (1 + 2 + ... + 10)^2 = 55^2 = 3025
 * これらの数の差は 3025 - 385 = 2640 となる.
 * 同様にして, 最初の100個の自然数について二乗の和と和の二乗の差を求めよ.
 */

func answer0(max int64) int64 {
	sum1 := int64(0)
	sum2 := int64(0)
	for n := int64(1); n <= max; n++ {
		sum1 += n * n
		sum2 += n
	}
	sum2 *= sum2
	return sum2 - sum1
}

func answer1(max int64) int64 {
	sum1 := (2*max + 1) * (max + 1) * max / 6 // 1^2 + 2^2 + ... + max^2
	sum2 := max * (max + 1) / 2               // (1 + 2 + ... + max)^2 = (max(max+1)/2)^2
	sum2 *= sum2
	return sum2 - sum1
}

func main() {
	fmt.Println(answer0(100))
	fmt.Println(answer1(100))
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
