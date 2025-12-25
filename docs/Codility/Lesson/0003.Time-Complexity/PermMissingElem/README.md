# [PermMissingElem](https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/)

An array A consisting of N different integers is given. The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.

Your goal is to find that missing element.

Write a function:

func Solution(A []int) int

that, given an array A, returns the value of the missing element.

For example, given array A such that:

  A[0] = 2
  A[1] = 3
  A[2] = 1
  A[3] = 5
the function should return 4, as it is the missing element.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
the elements of A are all distinct;
each element of array A is an integer within the range [1..(N + 1)].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意

## 解題思路

## 來源
* https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0003.Time-Complexity/PermMissingElem/PermMissingElem.go


```go
package permmissingelem

func Solution(A []int) int {
	if len(A) < 1 {
		return 1
	}

	n := len(A) + 1
	predictSume := (n + 1) * n / 2

	var sum int
	for _, v := range A {
		sum += v
	}
	return predictSume - sum
}
```