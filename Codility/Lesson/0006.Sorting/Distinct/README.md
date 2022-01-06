# [Distinct](https://app.codility.com/programmers/lessons/6-sorting/distinct/)
Compute number of distinct values in an array.

Write a function

func Solution(A []int) int

that, given an array A consisting of N integers, returns the number of distinct values in array A.

For example, given array A consisting of six elements such that:

 A[0] = 2    A[1] = 1    A[2] = 1
 A[3] = 2    A[4] = 3    A[5] = 1
the function should return 3, because there are 3 distinct values appearing in array A, namely 1, 2 and 3.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
返回不重複的整數數量

## 解題思路
1. 方法ㄧ: 先排序, 在檢查當前跟前一個是不是同一個整數. 不是的會結果+1
2. 方法二: 建立一個set. 最後返回set的長度

## 來源
https://app.codility.com/programmers/lessons/6-sorting/distinct/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0006.Sorting/Distinct/Distinct.go


```go
package Distinct

import "sort"

func Solution(A []int) int {
	sort.Ints(A)
	if len(A) == 0 {
		return 0
	}
	result := 1
	for i := 1; i < len(A); i++ {
		if A[i] != A[i-1] {
			result++
		}
	}
	return result
}

func SolutionSet(A []int) int {
	set := make(map[int]struct{})
	var void struct{}

	for i := 0; i < len(A); i++ {
		if _, ok := set[A[i]]; !ok {
			set[A[i]] = void
		}
	}
	return len(set)
}
```