# [MaxProductOfThree](https://app.codility.com/programmers/lessons/6-sorting/max_product_of_three/)
Maximize A[P] * A[Q] * A[R] for any triplet (P, Q, R).

A non-empty array A consisting of N integers is given. The product of triplet (P, Q, R) equates to A[P] * A[Q] * A[R] (0 ≤ P < Q < R < N).

For example, array A such that:

  A[0] = -3
  A[1] = 1
  A[2] = 2
  A[3] = -2
  A[4] = 5
  A[5] = 6
contains the following example triplets:

(0, 1, 2), product is −3 * 1 * 2 = −6
(1, 2, 4), product is 1 * 2 * 5 = 10
(2, 4, 5), product is 2 * 5 * 6 = 60
Your goal is to find the maximal product of any triplet.

Write a function:

func Solution(A []int) int

that, given a non-empty array A, returns the value of the maximal product of any triplet.

For example, given array A such that:

  A[0] = -3
  A[1] = 1
  A[2] = 2
  A[3] = -2
  A[4] = 5
  A[5] = 6
the function should return 60, as the product of triplet (2, 4, 5) is maximal.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [3..100,000];
each element of array A is an integer within the range [−1,000..1,000].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
給一整數陣列A，找出陣列中任意三數乘積的最大值

## 解題思路
先排序.然後比較 `前兩個元素*最後一個元素的乘積` 和 `最後三個元素的乘積` 取最大值

## 來源
https://app.codility.com/programmers/lessons/6-sorting/max_product_of_three/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0006.Sorting/MaxProductOfThree/MaxProductOfThree.go


```go
package MaxProductOfThree

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Solution(A []int) int {
	sort.Ints(A)
	aLen := len(A)
	return max(A[0]*A[1]*A[aLen-1], A[aLen-1]*A[aLen-2]*A[aLen-3])
}
```