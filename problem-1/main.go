package main

import "fmt"

/**
 * Multiples of 3 and 5
 *
 * https://projecteuler.net/problem=1
 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 * Find the sum of all the multiples of 3 or 5 below 1000.
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%201
 * 10未満の自然数のうち, 3 もしくは 5 の倍数になっているものは 3, 5, 6, 9 の4つがあり, これらの合計は 23 になる.
 * 同じようにして, 1000 未満の 3 か 5 の倍数になっている数字の合計を求めよ.
 */

func answer0(max int) int {
	sum := 0
	for n := 1; n <= max; n++ {
		if (n%3) == 0 || (n%5) == 0 {
			sum += n
		}
	}
	return sum
}

func sumDivisibleBy(max, n int) int {
	m := max / n
	return n * (m * (m + 1)) / 2
}
func answer1(max int) int {
	return sumDivisibleBy(max, 3) + sumDivisibleBy(max, 5) - sumDivisibleBy(max, 3*5)
}

func main() {
	fmt.Println(answer0(999))
	fmt.Println(answer1(999))
}
