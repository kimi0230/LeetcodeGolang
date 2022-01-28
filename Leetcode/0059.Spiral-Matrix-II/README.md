---
title: 0059.Spiral Matrix II
tags: Medium, Array
author: Kimi Tsai <kimi0230@gmail.com>
description:
---

# [59. Spiral Matrix II](https://leetcode.com/problems/spiral-matrix-ii/)


## 題目

Given a positive integer *n*, generate a square matrix filled with elements from 1 to *n*2 in spiral order.

**Example:**


    Input: 3
    Output:
    [
     [ 1, 2, 3 ],
     [ 8, 9, 4 ],
     [ 7, 6, 5 ]
    ]


## 題目大意

給定一個正整數 n，生成一個包含 1 到 n^2 所有元素，且元素按順時針順序螺旋排列的正方形矩陣。


## 解題思路

- 給出一個數組 n，要求輸出一個 n * n 的二維數組，裡面元素是 1 - n*n，且數組排列順序是螺旋排列的
- 這一題是第 54 題的加強版，沒有需要注意的特殊情況，直接模擬即可。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0059.Spiral-Matrix-II/
* https://leetcode-cn.com/problems/spiral-matrix-ii/
* [數組：這個循環可以轉懵很多人！](https://mp.weixin.qq.com/s/KTPhaeqxbMK9CxHUUgFDmg)

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0059.Spiral-Matrix-II/Spiral-Matrix-II.go

```go
package spiralmatrixii

// GenerateMatrix : 按層模擬, 時間複雜 O(n^2) 空間複雜 O(1)
func GenerateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	left, right, top, botton := 0, n-1, 0, n-1
	num := 1
	target := n * n

	for num <= target {
		// 上層  left to right, 上層邊界++
		for i := left; i <= right; i++ {
			result[top][i] = num
			num++
		}
		top++

		// 右層 top to botton , 右層邊界--
		for i := top; i <= botton; i++ {
			result[i][right] = num
			num++
		}
		right--

		// 下層 right to left , 下層邊界--
		for i := right; i >= left; i-- {
			result[botton][i] = num
			num++
		}
		botton--

		// 左層  botton to top, 左層邊界++
		for i := botton; i >= top; i-- {
			result[i][left] = num
			num++
		}
		left++
	}

	return result
}

// 模擬 : O(n)
// https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0059.Spiral-Matrix-II/
func GenerateMatrix2(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	result, visit, round, x, y, spDir := make([][]int, n), make([][]int, n), 0, 0, 0, [][]int{
		{0, 1},  // 朝右
		{1, 0},  // 朝下
		{0, -1}, // 朝左
		{-1, 0}, // 朝上
	}
	for i := 0; i < n; i++ {
		visit[i] = make([]int, n)
		result[i] = make([]int, n)
	}
	visit[x][y] = 1
	result[x][y] = 1

	for i := 0; i < n*n; i++ {
		x += spDir[round%4][0]
		y += spDir[round%4][1]
	}

	return result
}
```