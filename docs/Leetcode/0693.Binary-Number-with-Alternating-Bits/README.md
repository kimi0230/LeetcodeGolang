---
title: 0693.Binary Number with Alternating Bits
tags:
  - Medium
  - Bit Manipulation
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [693. Binary Number with Alternating Bits](https://leetcode.com/problems/binary-number-with-alternating-bits/)

## 題目

Given a positive integer, check whether it has alternating bits: namely, if two adjacent bits will always have different values.

**Example 1:**

    Input: 5
    Output: True
    Explanation:
    The binary representation of 5 is: 101

**Example 2:**

    Input: 7
    Output: False
    Explanation:
    The binary representation of 7 is: 111.

**Example 3:**

    Input: 11
    Output: False
    Explanation:
    The binary representation of 11 is: 1011.

**Example 4:**

    Input: 10
    Output: True
    Explanation:
    The binary representation of 10 is: 1010.


## 題目大意

給定一個正整數，檢查他是否為交替位二進制數：換句話說，就是他的二進制數相鄰的兩個位數永不相等。

## 解題思路

- 判斷一個數的二進制位相鄰兩個數是不相等的，即 `0101` 交叉間隔的，如果是，輸出 true。這一題有多種做法，最簡單的方法就是直接模擬。比較巧妙的方法是通過位運算，合理構造特殊數據進行位運算到達目的。 `010101` 構造出 `101010` 兩者相互 `&` 位運算以後就為 0，因為都“插空”了。

## 提示：
1 <= n <= 2^31 - 1

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0600~0699/0693.Binary-Number-with-Alternating-Bits/
* https://leetcode-cn.com/problems/binary-number-with-alternating-bits/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0693.Binary-Number-with-Alternating-Bits/Binary-Number-with-Alternating-Bits.go

```go
package binarynumberwithalternatingbits

// 暴力解 O(n)
func hasAlternatingBits(n int) bool {
	for n > 0 {
		preBit := n & 1
		n = n / 2
		curBit := n & 1
		if curBit == preBit {
			return false
		}
	}
	return true
}

// 數學解
func hasAlternatingBits2(n int) bool {
	/*
		n=5
		n=				1 0 1
		n >> 1			0 1 0
		n^(n>>1)		1 1 1  (XOR 不同時得1)
		n               1 1 1
		n+1			  1 0 0 0
		n & (n+1)	    0 0 0

		n=7
		n=				1 1 1
		n >> 1			0 1 1
		n^(n>>1)		1 0 0  (XOR 不同時得1)
		n               1 0 0
		n+1			    1 0 1
		n & (n+1)	    1 0 0
	*/
	n = n ^ (n >> 1)
	result := n & (n + 1)
	return result == 0
}
```