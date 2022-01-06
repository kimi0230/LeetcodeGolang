# [MissingInteger](https://app.codility.com/programmers/lessons/4-counting_elements/missing_integer/)
Find the smallest positive integer that does not occur in a given sequence.

This is a demo task.

Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.


## 題目大意
找出該array沒出現的最小整數

## 解題思路
先講出現的數字記錄起來, 再依序從1開始往後找出最小的整數且沒出現過

## 來源
https://app.codility.com/programmers/lessons/4-counting_elements/missing_integer/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0004.Counting-Elements/MissingInteger/MissingInteger.go


```go
package MissingInteger

func Solution(A []int) int {
	smallNum := 1
	intMap := make(map[int]bool)

	// 將出現的數字塞入map
	for _, v := range A {
		if v > 0 && !intMap[v] {
			intMap[v] = true
		}
	}

	for i := 1; i <= len(intMap); i++ {
		if !intMap[i] {
			// 此正整數沒在map找到
			return i
		}
		smallNum = i + 1
	}

	return smallNum
}
```