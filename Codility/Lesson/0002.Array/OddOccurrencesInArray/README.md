# [OddOccurrencesInArray](https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/)

A non-empty array A consisting of N integers is given. The array contains an odd number of elements, and each element of the array can be paired with another element that has the same value, except for one element that is left unpaired.

For example, in array A such that:

  A[0] = 9  A[1] = 3  A[2] = 9
  A[3] = 3  A[4] = 9  A[5] = 7
  A[6] = 9
the elements at indexes 0 and 2 have value 9,
the elements at indexes 1 and 3 have value 3,
the elements at indexes 4 and 6 have value 9,
the element at index 5 has value 7 and is unpaired.
Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers fulfilling the above conditions, returns the value of the unpaired element.

For example, given array A such that:

  A[0] = 9  A[1] = 3  A[2] = 9
  A[3] = 3  A[4] = 9  A[5] = 7
  A[6] = 9
the function should return 7, as explained in the example above.

Write an efficient algorithm for the following assumptions:

N is an odd integer within the range [1..1,000,000];
each element of array A is an integer within the range [1..1,000,000,000];
all but one of the values in A occur an even number of times.
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
給定一個非空array A，包含有N個整數，找出只出現基數次的整式

## 解題思路
1. 方法一: 可以用Map紀錄每個整數出現的次數, 在檢查是否是偶數
2. 方法二: 所有的整數XOR起來, 若是兩個整數相同XOR得到0, 最後剩下基數次的數字

## 來源
* https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/


## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0002.Array/OddOccurrencesInArray/OddOccurrencesInArray.go


```go
package oddoccurrencesinarray

func Solution(A []int) int {
	intMap := make(map[int]int)
	for i := 0; i < len(A); i++ {
		intMap[A[i]] += 1
	}

	for k, v := range intMap {
		if v%2 != 0 {
			return k
		}
	}
	return -1
}

// 所有的整數XOR起來, 若是兩個整數相同XOR得到0, 最後剩下基數次的數字
// 前提只能有一個基數數字
func Solution2(A []int) int {
	result := 0
	for i := 0; i < len(A); i++ {
		result ^= A[i]
	}
	return result
}
```