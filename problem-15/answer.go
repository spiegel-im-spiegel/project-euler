package problem12

import "math/big"

/**
 * Lattice paths
 *
 * https://projecteuler.net/problem=15
 * Starting in the top left corner of a 2×2 grid, and only being able to move to the right and down, there are exactly 6 routes to the bottom right corner.
 * How many such routes are there through a 20×20 grid?
 *
 * http://odz.sakura.ne.jp/projecteuler/index.php?cmd=read&page=Problem%2015
 * 2×2 のマス目の左上からスタートした場合, 引き返しなしで右下にいくルートは 6 つある.
 * では, 20×20 のマス目ではいくつのルートがあるか.
 *
 * Contents of Project Euler are licenced under a Creative Commons Licence: Attribution-NonCommercial-ShareAlike 2.0 UK: England & Wales.
 * http://creativecommons.org/licenses/by-nc-sa/2.0/uk/
 */

//Answer0 returns answer to this problem
func Answer0(grid int64) int64 {
	fg := factorial0(grid)
	fgfg := (&big.Int{}).Mul(fg, fg)
	fg2 := factorial0(2 * grid)
	return (&big.Int{}).Div(fg2, fgfg).Int64()
}
func factorial0(n int64) *big.Int {
	return (&big.Int{}).MulRange(1, n)
}

//Answer1 returns answer to this problem (refactoring version)
func Answer1(grid int64) int64 {
	return (&big.Int{}).Div(factorial1(2*grid, grid), factorial1(grid, 0)).Int64()
}
func factorial1(max, min int64) *big.Int {
	return (&big.Int{}).MulRange(min+1, max)
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
