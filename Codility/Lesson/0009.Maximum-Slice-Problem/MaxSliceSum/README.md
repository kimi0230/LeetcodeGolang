# [MaxSliceSum](https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_slice_sum/)
Find a maximum sum of a compact subsequence of array elements.

A non-empty array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P ≤ Q < N, is called a slice of array A. The sum of a slice (P, Q) is the total of A[P] + A[P+1] + ... + A[Q].

Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers, returns the maximum sum of any slice of A.

For example, given array A such that:

A[0] = 3  A[1] = 2  A[2] = -6
A[3] = 4  A[4] = 0
the function should return 5 because:

(3, 4) is a slice of A that has sum 4,
(2, 2) is a slice of A that has sum −6,
(0, 1) is a slice of A that has sum 5,
no other slice of A has sum greater than (0, 1).
Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..1,000,000];
each element of array A is an integer within the range [−1,000,000..1,000,000];
the result will be an integer within the range [−2,147,483,648..2,147,483,647].

Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
找出連續子序列最大的和
## 解題思路
長度如果為1, 回傳第一筆
當下的值跟當下的值加上先前的和, 取最大值. 再將剛剛算出的最大值跟紀錄中的最大值比較,取最大值

## 來源
https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_slice_sum/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0009.Maximum-Slice-Problem/MaxSliceSum/MaxSliceSum.go


```go
package MaxSliceSum

import (
	"math"
)

func Solution(A []int) int {
	if len(A) == 1 {
		return A[0]
	}
	result := math.MinInt64
	sum := math.MinInt64
	for i := 0; i < len(A); i++ {
		sum = int(math.Max(float64(A[i]), float64(A[i])+float64(sum)))
		result = int(math.Max(float64(sum), float64(result)))
	}
	return result
}
```