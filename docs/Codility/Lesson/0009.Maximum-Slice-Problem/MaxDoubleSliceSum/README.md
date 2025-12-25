# [MaxDoubleSliceSum](https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_double_slice_sum/)
Find the maximal sum of any double slice.

A non-empty array A consisting of N integers is given.

A triplet (X, Y, Z), such that 0 ≤ X < Y < Z < N, is called a double slice.

The sum of double slice (X, Y, Z) is the total of A[X + 1] + A[X + 2] + ... + A[Y − 1] + A[Y + 1] + A[Y + 2] + ... + A[Z − 1].

For example, array A such that:

    A[0] = 3
    A[1] = 2
    A[2] = 6
    A[3] = -1
    A[4] = 4
    A[5] = 5
    A[6] = -1
    A[7] = 2
contains the following example double slices:

double slice (0, 3, 6), sum is 2 + 6 + 4 + 5 = 17,
double slice (0, 3, 7), sum is 2 + 6 + 4 + 5 − 1 = 16,
double slice (3, 4, 5), sum is 0.
The goal is to find the maximal sum of any double slice.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the maximal sum of any double slice.

For example, given:

    A[0] = 3
    A[1] = 2
    A[2] = 6
    A[3] = -1
    A[4] = 4
    A[5] = 5
    A[6] = -1
    A[7] = 2
the function should return 17, because no double slice of array A has a sum of greater than 17.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [3..100,000];
each element of array A is an integer within the range [−10,000..10,000].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
A[X+1]到A[Y-1] + A[Y+1]到A[Z-1] 最大的和

## 解題思路
正向尋過array, 獲得到達每個index可以得到的最大值序列, 然后反向尋過array獲得到達每個index可以得到的最大值序列，
反向的的最大值序列需要倒轉.然後間隔一個位置，
最後尋遍array起兩者相加最大值




## 來源
* https://app.codility.com/programmers/lessons/9-maximum_slice_problem/max_double_slice_sum/
* https://github.com/Anfany/Codility-Lessons-By-Python3/blob/master/L9_Maximum%20Slice%20Problem/9.3%20%20MaxDoubleSliceSum.md

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0009.Maximum-Slice-Problem/MaxDoubleSliceSum/MaxDoubleSliceSum.go


```go
package MaxDoubleSliceSum

import (
	"math"
)

func Solution(A []int) int {
	if len(A) < 4 {
		return 0
	}
	N := len(A) - 2
	forwardSum := make([]int, N)
	reverseSum := make([]int, N)

	//  0 ≤ X < Y < Z < N,
	//  A[X + 1] + A[X + 2] + ... + A[Y − 1] + A[Y + 1] + A[Y + 2] + ... + A[Z − 1].
	// 			A : [ 3,  2, 6, -1,  4,  5, -1, 2]
	// forwardSum : [ 0,  2, 8,  7, 11, 16]
	// reverseSum : [14,  8, 9,  5,  0,  0]
	for i := 0; i < N-1; i++ {
		forwardVal := A[i+1]
		reverseVal := A[N-i]

		forwardSum[i+1] = int(math.Max(0, float64(forwardVal)+float64(forwardSum[i])))
		reverseSum[N-i-2] = int(math.Max(0, float64(reverseVal)+float64(reverseSum[N-i-1])))
	}

	combineMax := math.MinInt64
	for i := 0; i < N; i++ {
		combineMax = int(math.Max(float64(combineMax), float64(forwardSum[i])+float64(reverseSum[i])))
	}

	return combineMax
}
```