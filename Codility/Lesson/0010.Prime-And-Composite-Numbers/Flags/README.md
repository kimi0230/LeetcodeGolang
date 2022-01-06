# [Flags](https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/)
Find the maximum number of flags that can be set on mountain peaks.

A non-empty array A consisting of N integers is given.

A peak is an array element which is larger than its neighbours. More precisely, it is an index P such that 0 < P < N − 1 and A[P − 1] < A[P] > A[P + 1].

For example, the following array A:

    A[0] = 1
    A[1] = 5 //peaks
    A[2] = 3
    A[3] = 4 //peaks
    A[4] = 3
    A[5] = 4 //peaks
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6 //peaks
    A[11] = 2
has exactly four peaks: elements 1, 3, 5 and 10.

You are going on a trip to a range of mountains whose relative heights are represented by array A, as shown in a figure below. You have to choose how many flags you should take with you. The goal is to set the maximum number of flags on the peaks, according to certain rules.

![](https://codility-frontend-prod.s3.amazonaws.com/media/task_static/flags/static/images/auto/6f5e8faa3000c0a74157e6e0bc759b8a.png)

Flags can only be set on peaks. What's more, if you take K flags, then the distance between any two flags should be greater than or equal to K. The distance between indices P and Q is the absolute value |P − Q|.

For example, given the mountain range represented by array A, above, with N = 12, if you take:

two flags, you can set them on peaks 1 and 5;
three flags, you can set them on peaks 1, 5 and 10;
four flags, you can set only three flags, on peaks 1, 5 and 10.
You can therefore set a maximum of three flags in this case.

Write a function:

func Solution(A []int) int

that, given a non-empty array A of N integers, returns the maximum number of flags that can be set on the peaks of the array.

For example, the following array A:

    A[0] = 1
    A[1] = 5
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2
the function should return 3, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..400,000];
each element of array A is an integer within the range [0..1,000,000,000].

Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
計算可以在山峰上設置的最大旗標數量
旗標只能在山峰上設置, 如果帶了K個旗標, 則任意兩個旗標的索引距離不能小於K

## 解題思路
先找出peak的索引位置並存入array中. 根據peak array的第一個最後一個可以判斷理論上最多的旗標數為K: K*(K-1)<=P[-1] - P[0]=dis, 所以K的最大值為 int(sqrt(P[-1] - P[0]) +1).
然後從K的最大值開始找, 尋遍peak array, 只要滿足距離大於等於就將旗標存入旗標array, 只要旗標array數不小於K值，就返回

## 來源
* https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/
* https://github.com/Anfany/Codility-Lessons-By-Python3/blob/master/L10_Prime%20and%20composite%20numbers/10.4%20Flags.md

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0010.Prime-And-Composite-Numbers/Flags/Flags.go


```go
package Flags

import (
	"math"
)

func Solution(A []int) int {
	var peaksFlag []int

	for i := 1; i < len(A)-1; i++ {
		if A[i] > A[i-1] && A[i] > A[i+1] {
			peaksFlag = append(peaksFlag, i)
		}
	}

	if len(peaksFlag) == 0 {
		return 0
	}
	if len(peaksFlag) == 1 {
		return 1
	}

	maxFlag := int(math.Pow(float64(peaksFlag[len(peaksFlag)-1]-peaksFlag[0]), 0.5) + 1)

	for i := maxFlag; i > 1; i-- {
		addressFlag := []int{peaksFlag[0]}
		for _, val := range peaksFlag[1:] {
			if val-addressFlag[len(addressFlag)-1] >= i {
				addressFlag = append(addressFlag, val)
				if len(addressFlag) >= i {
					return i
				}
			}
		}
	}

	return 1
}
```