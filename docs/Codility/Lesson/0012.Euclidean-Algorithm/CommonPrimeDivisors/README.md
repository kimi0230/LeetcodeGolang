# [CommonPrimeDivisors](https://app.codility.com/programmers/lessons/12-euclidean_algorithm/common_prime_divisors/)
Check whether two numbers have the same prime divisors.

A prime is a positive integer X that has exactly two distinct divisors: 1 and X. The first few prime integers are 2, 3, 5, 7, 11 and 13.

A prime D is called a prime divisor(質因數) of a positive integer P if there exists a positive integer K such that D * K = P. For example, 2 and 5 are prime divisors of 20.

You are given two positive integers N and M. The goal is to check whether the sets of prime divisors of integers N and M are exactly the same.

For example, given:

N = 15 and M = 75, the prime divisors are the same: {3, 5};
N = 10 and M = 30, the prime divisors aren't the same: {2, 5} is not equal to {2, 3, 5};
N = 9 and M = 5, the prime divisors aren't the same: {3} is not equal to {5}.
Write a function:

func Solution(A []int, B []int) int

that, given two non-empty arrays A and B of Z integers, returns the number of positions K for which the prime divisors of A[K] and B[K] are exactly the same.

For example, given:

    A[0] = 15   B[0] = 75
    A[1] = 10   B[1] = 30
    A[2] = 3    B[2] = 5
the function should return 1, because only one pair (15, 75) has the same set of prime divisors.

Write an efficient algorithm for the following assumptions:

Z is an integer within the range [1..6,000];
each element of arrays A, B is an integer within the range [1..2,147,483,647].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
## 題目大意
判斷兩個數是否有相同的質因數

## 解題思路
先判斷兩數的最大公因數, 再判斷兩個數是含有最大公因數沒有的因子
15 , 75 的最大公因數為 3*5
15= 3*5
75= 3*5*5

## 來源
* https://app.codility.com/programmers/lessons/12-euclidean_algorithm/common_prime_divisors/
* https://github.com/Anfany/Codility-Lessons-By-Python3/blob/master/L12_Euclidean%20algorithm/12.2%20CommonPrimeDivisors.md
* https://github.com/Luidy/codility-golang/blob/master/Lesson12/02_commonPrimeDivisors.go

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0012.Euclidean-Algorithm/CommonPrimeDivisors/CommonPrimeDivisors.go


```go
package commonprimedivisors

/*
func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
*/
func gcd(N int, M int) int {
	if N%M == 0 {
		return M
	} else {
		return gcd(M, N%M)
	}
}

func Solution(A []int, B []int) int {
	result := 0
	for i := 0; i < len(A); i++ {
		if A[i] == B[i] {
			result++
			continue
		}
		// 先判斷兩數的最大公因數,
		abGcd := gcd(A[i], B[i])

		// 再判斷兩個數是含有最大公因數沒有的因子
		a := A[i] / abGcd
		aGcd := gcd(a, abGcd)
		for aGcd != 1 {
			// 還有其他因子
			a = a / aGcd
			aGcd = gcd(aGcd, a)
		}

		b := B[i] / abGcd
		bGcd := gcd(b, abGcd)
		for bGcd != 1 {
			b = b / bGcd
			bGcd = gcd(bGcd, b)
		}
		if a == b {
			result++
		}
	}
	return result
}
```