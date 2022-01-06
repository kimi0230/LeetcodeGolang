# [CountNonDivisible](https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/count_non_divisible/)
Calculate the number of elements of an array that are not divisors(因數) of each element.

You are given an array A consisting of N integers.

For each number A[i] such that 0 ≤ i < N, we want to count the number of elements of the array that are not the divisors of A[i]. We say that these elements are non-divisors.

For example, consider integer N = 5 and array A such that:

    A[0] = 3
    A[1] = 1
    A[2] = 2
    A[3] = 3
    A[4] = 6
For the following elements:

A[0] = 3, the non-divisors are: 2, 6,
A[1] = 1, the non-divisors are: 3, 2, 3, 6,
A[2] = 2, the non-divisors are: 3, 3, 6,
A[3] = 3, the non-divisors are: 2, 6,
A[4] = 6, there aren't any non-divisors.
Write a function:

func Solution(A []int) []int

that, given an array A consisting of N integers, returns a sequence of integers representing the amount of non-divisors.

Result array should be returned as an array of integers.

For example, given:

    A[0] = 3
    A[1] = 1
    A[2] = 2
    A[3] = 3
    A[4] = 6
the function should return [2, 4, 3, 2, 0], as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..50,000];
each element of array A is an integer within the range [1..2 * N].

Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
算出array中所有元素的非因子數的個數

## 解題思路
先算數每個數字出現的次數存入map
遍歷A, 對於每個元素從1到sqrt(i)中找出因子,如果是因子,就去字典找出出現次數
最後用總長度減去因子數就可得出非因子數, 並將結果存入map, 空間換取時間

factor <= int(math.Pow(float64(val), 0.5)) 改成 factor*factor <= val 可提高效能

## 來源
* https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/count_non_divisible/
* https://github.com/Anfany/Codility-Lessons-By-Python3/blob/master/L11_Sieve%20of%20Eratosthenes/11.1%20CountNonDivisible.md

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0011.Sieve-of-Eratosthenes/CountNonDivisible/CountNonDivisible.go


```go
package countnondivisible

import (
	"math"
)

/*
Task Score 100%
Correctness 100%
Performance 100%
*/
func Solution(A []int) []int {
	// write your code in Go 1.4
	result := []int{}
	if len(A) < 0 {
		return result
	}

	elementDict := make(map[int]int)
	for _, i := range A {
		elementDict[i]++
	}

	// 用 map存起非因子數的個數, 空間換取時間
	nonDivisorsCountMap := make(map[int]int)
	for _, val := range A {
		if v, ok := nonDivisorsCountMap[val]; ok {
			result = append(result, v)
		} else {
			divisors := 0
			for factor := 1; factor*factor <= val; factor++ {
				if val%factor == 0 {
					// 檢查是否有在原先的 array 中, 並取得因子次數
					if v, ok := elementDict[factor]; ok {
						divisors += v
					}

					// 避免因子重複計算
					otherFactor := int(val / factor)
					if v, ok := elementDict[otherFactor]; ok && otherFactor != factor {
						divisors += v
					}
				}
			}
			// 推出非因子次數
			nonDivisors := len(A) - divisors
			result = append(result, nonDivisors)
			nonDivisorsCountMap[val] = nonDivisors
		}
	}
	return result
}

/*
Task Score 88%
Correctness 100%
Performance 75%
*/
func Solution2(A []int) []int {
	result := []int{}
	if len(A) < 0 {
		return result
	}

	elementDict := make(map[int]int)
	for _, i := range A {
		elementDict[i]++
	}

	// 用 map存起非因子數的個數, 空間換取時間
	nonDivisorsCountMap := make(map[int]int)
	for _, val := range A {
		if v, ok := nonDivisorsCountMap[val]; ok {
			result = append(result, v)
		} else {
			divisors := 0
			for factor := 1; factor <= int(math.Pow(float64(val), 0.5)); factor++ {
				if val%factor == 0 {
					// 檢查是否有在原先的 array 中, 並取得因子次數
					if v, ok := elementDict[factor]; ok {
						divisors += v
					}

					// 避免因子重複計算
					otherFactor := int(val / factor)
					if v, ok := elementDict[otherFactor]; ok && otherFactor != factor {
						divisors += v
					}
				}
			}
			// 推出非因子次數
			nonDivisors := len(A) - divisors
			result = append(result, nonDivisors)
			nonDivisorsCountMap[val] = nonDivisors
		}
	}
	return result
}

/*
Task Score 77%
Correctness 100%
Performance 50%
*/
func Solution1(A []int) []int {
	result := []int{}
	if len(A) < 0 {
		return result
	}

	elementDict := make(map[int]int)
	for _, i := range A {
		elementDict[i]++
	}

	for _, val := range A {
		divisors := 0
		for factor := 1; factor <= int(math.Pow(float64(val), 0.5)); factor++ {
			if val%factor == 0 {
				// 檢查是否有在原先的 array 中, 並取得因子次數
				if v, ok := elementDict[factor]; ok {
					divisors += v
				}

				// 避免因子重複計算
				otherFactor := int(val / factor)
				if v, ok := elementDict[otherFactor]; ok && otherFactor != factor {
					divisors += v
				}
			}
		}
		// 推出非因子次數
		nonDivisors := len(A) - divisors
		result = append(result, nonDivisors)
	}
	return result
}

```