# [MinPerimeterRectangle](https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/min_perimeter_rectangle/)
Find the minimal perimeter of any rectangle whose area equals N.

An integer N is given, representing the area of some rectangle.

The area of a rectangle whose sides are of length A and B is A * B, and the perimeter is 2 * (A + B).

The goal is to find the minimal perimeter of any rectangle whose area equals N. The sides of this rectangle should be only integers.

For example, given integer N = 30, rectangles of area 30 are:

(1, 30), with a perimeter of 62,
(2, 15), with a perimeter of 34,
(3, 10), with a perimeter of 26,
(5, 6), with a perimeter of 22.
Write a function:

func Solution(N int) int

that, given an integer N, returns the minimal perimeter of any rectangle whose area is exactly equal to N.

For example, given an integer N = 30, the function should return 22, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..1,000,000,000].

Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.


## 題目大意
給整數的面積N, 找出面積為N的最小周長


## 解題思路
從不大於N的平方根的數開始遍歷,只要找到N的因子
因為越往後所得的周長越大.邊長接近平方根的矩形的周長是最小的

## 來源
* https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/min_perimeter_rectangle/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0010.Prime-And-Composite-Numbers/MinPerimeterRectangle/MinPerimeterRectangle.go


```go
package minperimeterrectangle

import (
	"math"
)

// O(sqrt(N))
func Solution(N int) int {
	if N <= 0 {
		return 0
	}

	min := math.MaxInt32
	for i := 1; i*i <= N; i++ {
		if N%i == 0 {
			perimeter := 2 * (i + N/i)
			min = int(math.Min(float64(min), float64(perimeter)))
		}

	}

	if min == math.MaxInt32 {
		return 0
	}
	return min
}

/*
O(N)
Task Score 60%
Correctness 100%
Performance 20%
*/
func Solution1(N int) int {
	if N <= 0 {
		return 0
	}

	min := math.MaxInt32
	for i := 1; i <= N; i++ {
		if N%i == 0 && i*i <= N {
			perimeter := 2 * (i + N/i)
			min = int(math.Min(float64(min), float64(perimeter)))
		}
	}
	if min == math.MaxInt32 {
		return 0
	}
	return min
}

// O(sqrt(N))
func Solution2(N int) int {
	if N <= 0 {
		return 0
	}
	pairs := make(map[int]int)
	i := 1
	for i*i <= N {
		if N%i == 0 {
			pairs[i] = N / i
		}
		i++
	}

	min := math.MaxInt32
	for i, v := range pairs {
		perimeter := 2 * (i + v)
		min = int(math.Min(float64(min), float64(perimeter)))

	}

	if min == math.MaxInt32 {
		return 0
	}
	return min
}
```