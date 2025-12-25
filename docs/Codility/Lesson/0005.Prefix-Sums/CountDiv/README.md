# [CountDiv](https://app.codility.com/programmers/lessons/5-prefix_sums/count_div/)
Compute number of integers divisible by k in range [a..b].

Write a function:

func Solution(A int, B int, K int) int

that, given three integers A, B and K, returns the number of integers within the range [A..B] that are divisible by K, i.e.:

{ i : A ≤ i ≤ B, i mod K = 0 }

For example, for A = 6, B = 11 and K = 2, your function should return 3, because there are three numbers divisible by 2 within the range [6..11], namely 6, 8 and 10.

Write an efficient algorithm for the following assumptions:

A and B are integers within the range [0..2,000,000,000];
K is an integer within the range [1..2,000,000,000];
A ≤ B.
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
A~B之間的數字 mod K 後為0 的數字有幾個

## 解題思路
B/K 找出最大的商, A/K 最小的商. 相減取得在此中間之商的數量. 如果A%K==0 需要在+1

## 來源
* https://app.codility.com/programmers/lessons/5-prefix_sums/count_div/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0005.Prefix-Sums/CountDiv/CountDiv.go


```go
package CountDiv

import "math"

// 時間: O(n)
func SolutionBurst(A int, B int, K int) int {
	result := 0
	for i := A; i <= B; i++ {
		if i%K == 0 {
			result++
		}
	}

	return result
}

// 時間:O(1) 空間: O(1)
func Solution(A int, B int, K int) int {
	result := 0
	if A%2 == 0 {
		result = 1
	}
	return int(math.Floor(float64(B/K))) - int(math.Floor(float64(A/K))) + result
}

```