---
title: 0354. Russian Doll Envelope
tags: Hard, Dynamic Programming, Binary Search
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [354. Russian Doll Envelope](https://leetcode.com/problems/russian-doll-envelopes/)

## 題目
You are given a 2D array of integers envelopes where envelopes[i] = [wi, hi] represents the width and the height of an envelope.

One envelope can fit into another if and only if both the width and height of one envelope are greater than the other envelope's width and height.

Return the maximum number of envelopes you can Russian doll (i.e., put one inside the other).

Note: You cannot rotate an envelope.

 

Example 1:
```
Input: envelopes = [[5,4],[6,4],[6,7],[2,3]]
Output: 3
Explanation: The maximum number of envelopes you can Russian doll is 3 ([2,3] => [5,4] => [6,7]).
```

Example 2:
```
Input: envelopes = [[1,1],[1,1],[1,1]]
Output: 1
```

Constraints:

* 1 <= envelopes.length <= 105
* envelopes[i].length == 2
* 1 <= wi, hi <= 105

## 題目大意
給一些信封, 每個信封用寬度和高度的整數對形式(w,h)表示, 當一個信封A的寬度和高度都比另一個信封B大的時候, 則B就可以放進A裡.
計算出最多有多少信封能組成一組

## 解題思路
先對寬度w進行升序排序, 如果遇到w相同的情況, 則按照高度進行降序排序.
之後把所有的h最為一個array, 對此array做 最長遞增子序列(Longest Increasing Subsequence, LIS)的長度就是答案

此題是二維的LIS問題, 一維LIS參考 https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0300.Longest-Increasing-Subsequence/main.go
三維 Leetcode No.1691

## 來源
* https://leetcode.com/problems/russian-doll-envelopes/
* https://leetcode-cn.com/problems/russian-doll-envelopes/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0354.Russian-Doll-Envelopes/main.go

```go
package russiandollenvelopes

import (
	longestincreasingsubsequence "LeetcodeGolang/Leetcode/0300.Longest-Increasing-Subsequence"
	"sort"
)

type sortEnvelopes [][]int

func (s sortEnvelopes) Len() int {
	return len(s)
}

func (s sortEnvelopes) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		// 遇到w相同的情況, 則按照高度進行降序排序
		return s[i][1] > s[j][1]
	}
	// 對寬度w進行升序排序
	return s[i][0] < s[j][0]
}

func (s sortEnvelopes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func MaxEnvelopes(envelopes [][]int) int {
	// 先對寬度w進行升序排序, 如果遇到w相同的情況, 則按照高度進行降序排序.
	sort.Sort(sortEnvelopes(envelopes))
	// fmt.Println(envelopes)

	// 所有的h最為一個array, 對此array做 最長遞增子序列(Longest Increasing Subsequence, LIS)的長度就是答案
	height := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		height[i] = envelopes[i][1]
	}

	//  DP + 二分搜尋:patience sorting. O(nlogn)
	return longestincreasingsubsequence.LengthOfLISPatience(height)
}

func MaxEnvelopes2(envelopes [][]int) int {
	// 先對寬度w進行升序排序, 如果遇到w相同的情況, 則按照高度進行降序排序.
	sort.Sort(sortEnvelopes(envelopes))
	// fmt.Println(envelopes)

	// 最長遞增子序列(Longest Increasing Subsequence, LIS)的長度就是答案
	dp := []int{}
	for _, e := range envelopes {
		left, right := 0, len(dp)
		for left < right {
			mid := left + (right-left)/2
			if dp[mid] > e[1] {
				// 現在的牌比堆小, 所小範圍
				right = mid
			} else if dp[mid] < e[1] {
				// 現在的牌比堆大
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 沒有找到堆, 建立一個新的堆
		if left == len(dp) {
			dp = append(dp, e[1])
		}

		// 再把這張牌放在堆的頂端
		dp[left] = e[1]
	}

	return len(dp)
}

/*
func Sort(data Interface) {
	n := data.Len()
	quickSort(data, 0, n, maxDepth(n))
}
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	//
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/
```

```sh
go test -benchmem -run=none LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkMaxEnvelopes-8          9067520               167.0 ns/op            88 B/op          3 allocs/op
BenchmarkMaxEnvelopes2-8         6726646               214.9 ns/op            80 B/op          4 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0354.Russian-Doll-Envelopes     6.237s
```