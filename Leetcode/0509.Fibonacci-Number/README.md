---
title: 0509.Fibonacci Number
tags: Easy, Dynamic Programming
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [509. Fibonacci Number](https://leetcode.com/problems/fibonacci-number/)

## 題目

The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

```
F(0) = 0, F(1) = 1
F(n) = F(n - 1) + F(n - 2), for n > 1.
```
Given n, calculate F(n).


Example 1:
```
Input: n = 2
Output: 1
Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
```

Example 2:
```
Input: n = 3
Output: 2
Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
```

Example 3:
```
Input: n = 4
Output: 3
Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
```

Constraints:
* 0 <= n <= 30

## 題目大意
斐波那契數列, 通常用 F(n) 表示

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), 其中 N > 1.
給定 N，計算 F(N)。

提示：0 ≤ N ≤ 30
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377 ,610, 987……

## 解題思路
遇到遞迴最好畫出遞迴樹

```
            f(20)
           /     \ 
        f(19)   f(18)
       ...           ...
      /  \           /  \   
    f(1) f(2)       f(1) f(2) 
```

這一題解法很多，大的分類是四種，遞歸，記憶化搜索(dp)，矩陣快速冪，通項公式。其中記憶化搜索可以寫 3 種方法，自底向上的，自頂向下的，優化空間複雜度版的。通項公式方法實質是求 a^b 這個還可以用快速冪優化時間複雜度到 O(log n) 。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0500~0599/0509.Fibonacci-Number/
* https://leetcode.com/problems/fibonacci-number/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0509.Fibonacci-Number/Fibonacci-Number.go

```go
// --- Directions
// Print out the n-th entry in the fibonacci series.
// The fibonacci series is an ordering of numbers where
// each number is the sum of the preceeding two.
// For example, the sequence
//  [0, 1, 1, 2, 3, 5, 8, 13, 21, 34]
// forms the first ten entries of the fibonacci series.
// Example:
//   fib(4) === 3

package fibonaccinumber

import "math"

// Fib : iterative 迴圈  O(n) . 空間複雜  O(n). 自底向上的記憶化搜索
func FibIterative(n int) int {
	var result = []int{0, 1}

	for i := 2; i <= n; i++ {
		a := result[i-1]
		b := result[i-2]
		result = append(result, a+b)
	}
	return result[n]
}

// Fibrecur : recursive 遞迴 O(2^n) . 空間複雜  O(n)
func Fibrecur(n int) int {
	if n < 2 {
		return n
	}

	return Fibrecur(n-1) + Fibrecur(n-2)
}

// the memo table
var memo = map[int]int{}

// FibDP : memoization 備忘錄的遞迴
// 比較快!
func FibDP(n int) int {
	// 如果有cached 直接回傳
	if v, vok := memo[n]; vok == true {
		// cached
		return v
	}

	if _, vok := memo[0]; vok == false {
		// 沒cache
		memo[0] = 0
	}

	if _, vok := memo[1]; vok == false {
		// 沒cache
		memo[1] = 1
	}

	for i := 2; i <= n; i++ {
		if _, vok := memo[i]; vok == false {
			// 沒cache
			a := memo[i-1]
			b := memo[i-2]
			memo[i] = a + b
		}
		// fmt.Println(memo)
	}
	return memo[n]
	// return FibDP(n-1) + FibDP(n-2)
}

// 狀態壓縮
func FibDPStateCompression(n int) int {
	if n == 0 {
		return 0
	}
	if n == 2 || n == 1 {
		return 1
	}

	prev := 1
	curr := 1

	for i := 3; i <= n; i++ {
		sum := prev + curr
		prev = curr
		curr = sum
	}
	return curr
}

var cache = map[int]int{}

func memoize(fn func(int) int) func(int) int {
	return func(args int) int {
		if _, vok := cache[args]; vok == true {
			return cache[args]
		}

		result := fn(args) // 跑 Fibrecur(args)
		cache[args] = result
		return result
	}
}

// Fibmem : memoization 拉出遞迴function
func Fibmem(n int) int {
	result := memoize(Fibrecur)
	return result(n)
}

// 解法五 矩阵快速幂 时间复杂度 O(log n)，空间复杂度 O(log n)
// | 1 1 | ^ n   = | F(n+1) F(n)   |
// | 1 0 |		   | F(n)	F(n-1) |
func Fib5(N int) int {
	if N <= 1 {
		return N
	}
	var A = [2][2]int{
		{1, 1},
		{1, 0},
	}
	A = matrixPower(A, N-1)
	return A[0][0]
}

func matrixPower(A [2][2]int, N int) [2][2]int {
	if N <= 1 {
		return A
	}
	A = matrixPower(A, N/2)
	A = multiply(A, A)

	var B = [2][2]int{
		{1, 1},
		{1, 0},
	}
	if N%2 != 0 {
		A = multiply(A, B)
	}

	return A
}

func multiply(A [2][2]int, B [2][2]int) [2][2]int {
	x := A[0][0]*B[0][0] + A[0][1]*B[1][0]
	y := A[0][0]*B[0][1] + A[0][1]*B[1][1]
	z := A[1][0]*B[0][0] + A[1][1]*B[1][0]
	w := A[1][0]*B[0][1] + A[1][1]*B[1][1]
	A[0][0] = x
	A[0][1] = y
	A[1][0] = z
	A[1][1] = w
	return A
}

// 解法六 公式法 f(n)=(1/√5)*{[(1+√5)/2]^n -[(1-√5)/2]^n}，用 时间复杂度在 O(log n) 和 O(n) 之间，空间复杂度 O(1)
// 经过实际测试，会发现 pow() 系统函数比快速幂慢，说明 pow() 比 O(log n) 慢
// 斐波那契数列是一个自然数的数列，通项公式却是用无理数来表达的。而且当 n 趋向于无穷大时，前一项与后一项的比值越来越逼近黄金分割 0.618（或者说后一项与前一项的比值小数部分越来越逼近 0.618）。
// 斐波那契数列用计算机计算的时候可以直接用四舍五入函数 Round 来计算。
func Fib6(N int) int {
	var goldenRatio float64 = float64((1 + math.Sqrt(5)) / 2)
	return int(math.Round(math.Pow(goldenRatio, float64(N)) / math.Sqrt(5)))
}

// 解法七 協程版，但是时间特别慢，不推荐，放在这里只是告诉大家，写 LeetCode 算法题的时候，启动 goroutine 特别慢
func Fib7(N int) int {
	return <-fibb(N)
}

func fibb(n int) <-chan int {
	result := make(chan int)
	go func() {
		defer close(result)

		if n <= 1 {
			result <- n
			return
		}
		result <- <-fibb(n-1) + <-fibb(n-2)
	}()
	return result
}
```