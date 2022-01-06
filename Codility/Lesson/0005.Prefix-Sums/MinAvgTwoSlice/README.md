# [MinAvgTwoSlice](https://app.codility.com/programmers/lessons/5-prefix_sums/min_avg_two_slice/)
Find the minimal average of any slice containing at least two elements.

A non-empty array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P < Q < N, is called a slice of array A (notice that the slice contains at least two elements). The average of a slice (P, Q) is the sum of A[P] + A[P + 1] + ... + A[Q] divided by the length of the slice. To be precise, the average equals (A[P] + A[P + 1] + ... + A[Q]) / (Q − P + 1).

For example, array A such that:

    A[0] = 4
    A[1] = 2
    A[2] = 2
    A[3] = 5
    A[4] = 1
    A[5] = 5
    A[6] = 8
contains the following example slices:

slice (1, 2), whose average is (2 + 2) / 2 = 2;
slice (3, 4), whose average is (5 + 1) / 2 = 3;
slice (1, 4), whose average is (2 + 2 + 5 + 1) / 4 = 2.5.
The goal is to find the starting position of a slice whose average is minimal.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the starting position of the slice with the minimal average. If there is more than one slice with a minimal average, you should return the smallest starting position of such a slice.

For example, given array A such that:

    A[0] = 4
    A[1] = 2
    A[2] = 2
    A[3] = 5
    A[4] = 1
    A[5] = 5
    A[6] = 8
the function should return 1, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [2..100,000];
each element of array A is an integer within the range [−10,000..10,000].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
找出array兩index中平均值最小 ,並回傳start的 index

## 解題思路
最小avg的slice(n,m)，m-n+1一定是2或者3，也就是這個最小avg的slice由2個或者3個元素組成
因為題目中說明 `0<=n<m<N` 因此可以得出2個或者3個元素是最小的組合，比如length=3的數組，你無法一次分出2個slice，length=2的數組也一樣。為什麼要這麼去想呢？因為你要“比較”出最小的avg，怎麼才能"比較"？那就是必須一次至少有2個slice才能相互比較。那麼當N>=4時，我們就能一次最少分出2個slice

## 來源
* https://app.codility.com/programmers/lessons/5-prefix_sums/min_avg_two_slice/
* https://blog.csdn.net/dear0607/article/details/42581149

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0005.Prefix-Sums/MinAvgTwoSlice/MinAvgTwoSlice.go


```go
package MinAvgTwoSlice

import "math"

func Solution(A []int) int {

	min := math.MaxFloat64
	minIndex := -1
	for i := 0; i < len(A)-1; i++ {
		// 2個數平均
		if i+1 < len(A) {
			tmp := (float64(A[i]) + float64(A[i+1])) / float64(2)
			if tmp < min {
				min = tmp
				minIndex = i
			}
		}
		// 3個數平均
		if i+2 < len(A) {
			tmp := (float64(A[i]) + float64(A[i+1]) + float64(A[i+2])) / float64(3)
			if tmp < min {
				min = tmp
				minIndex = i
			}
		}
	}

	return minIndex
}
```