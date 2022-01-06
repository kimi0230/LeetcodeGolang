# [CountSemiprimes](https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/count_semiprimes/)
Count the semiprime(半質數:兩個質數的乘積所得的自然數我們稱之為半質數) numbers in the given range [a..b]

A prime is a positive integer X that has exactly two distinct divisors: 1 and X. The first few prime integers are 2, 3, 5, 7, 11 and 13.

A semiprime is a natural number that is the product of two (not necessarily distinct) prime numbers. The first few semiprimes are 4, 6, 9, 10, 14, 15, 21, 22, 25, 26.

You are given two non-empty arrays P and Q, each consisting of M integers. These arrays represent queries about the number of semiprimes within specified ranges.

Query K requires you to find the number of semiprimes within the range (P[K], Q[K]), where 1 ≤ P[K] ≤ Q[K] ≤ N.

For example, consider an integer N = 26 and arrays P, Q such that:

    P[0] = 1    Q[0] = 26
    P[1] = 4    Q[1] = 10
    P[2] = 16   Q[2] = 20
The number of semiprimes within each of these ranges is as follows:

(1, 26) is 10,
(4, 10) is 4,
(16, 20) is 0.
Write a function:

func Solution(N int, P []int, Q []int) []int

that, given an integer N and two non-empty arrays P and Q consisting of M integers, returns an array consisting of M elements specifying the consecutive answers to all the queries.

For example, given an integer N = 26 and arrays P, Q such that:

    P[0] = 1    Q[0] = 26
    P[1] = 4    Q[1] = 10
    P[2] = 16   Q[2] = 20
the function should return the values [10, 4, 0], as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..50,000];
M is an integer within the range [1..30,000];
each element of arrays P, Q is an integer within the range [1..N];
P[i] ≤ Q[i].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.


## 題目大意
計算[a,b]區間得半質數的個數

## 解題思路
先把半質數列表找出來. 
(半質數:兩個質數的乘積所得的自然數我們稱之為半質數ㄝ開始的幾個半質數是4, 6, 9, 10, 14, 15, 21, 22, 25, 26, ... （OEIS中的數列A001358）它們包含1及自己在內合共有3或4個因數)
計算當前array和前面一共有幾個半質數
遍歷P arry 算出 Q跟P 之前的半質數總和差, 並寫入結果


## 來源
* https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/count_semiprimes/
* https://github.com/Luidy/codility-golang/blob/master/Lesson11/02_countSemiprimes.go
* https://github.com/Anfany/Codility-Lessons-By-Python3/blob/master/L11_Sieve%20of%20Eratosthenes/11.2%20CountSemiprimes.md

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0011.Sieve-of-Eratosthenes/CountSemiprimes/CountSemiprimes.go


```go
package countsemiprimes

func Solution(N int, P []int, Q []int) []int {
	semiPrime := []int{}

	// 半質數:兩個質數的乘積所得的自然數我們稱之為半質數.
	// 4, 6, 9, 10, 14, 15, 21,22,25,26,33,34,35,38,39,46,49,51,55,57,58,62,65,69,74,77,82,85,86,87,91,93,94,95,106, ...
	// 它們包含1及自己在內合共有3或4個因數
	for i := 1; i <= N; i++ {
		factorCount := 0
		sign := 0
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				factorCount++
				f := i / j
				if f != j {
					if f == j*j {
						// 3個相同: ex i=27, j=3, f=9
						sign = 1
						semiPrime = append(semiPrime, 0)
						break
					} else {
						factorCount++
					}
				}
			}
			if factorCount > 4 {
				sign = 1
				semiPrime = append(semiPrime, 0)
				break
			}
		}
		if sign != 1 {
			if factorCount >= 3 {
				semiPrime = append(semiPrime, i)
			} else {
				semiPrime = append(semiPrime, 0)
			}
		}
	}
	// idx		 0 1 2 3 4 5 6 7 8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25
	// semiPrime 0 0 0 4 0 6 0 0 9 10  0  0  0 14 15  0  0  0  0  0 21 22  0  0 25 26

	// fmt.Println("semiPrime", semiPrime)

	// 當前array和前面一共有幾個半質數
	indexMap := make(map[int]int)
	// 如果是半質數添加到 map
	semiMap := make(map[int]struct{})
	count := 0

	for i := 0; i < len(semiPrime); i++ {
		if semiPrime[i] != 0 {
			count++
			indexMap[semiPrime[i]] = count
			semiMap[semiPrime[i]] = struct{}{}
		} else {
			indexMap[i+1] = count
		}
	}
	// 	indexMap :  map[1:0 2:0 3:0 4:1 5:1 6:2 7:2 8:2 9:3 10:4 11:4 12:4 13:4 14:5 15:6 16:6 17:6 18:6 19:6 20:6 21:7 22:8 23:8 24:8 25:9 26:10]
	// semiMap :  map[4:0 6:0 9:0 10:0 14:0 15:0 21:0 22:0 25:0 26:0]
	// fmt.Println("indexMap : ", indexMap)
	// fmt.Println("semiMap : ", semiMap)

	result := []int{}
	for i := 0; i < len(P); i++ {
		if _, ok := semiMap[P[i]]; ok {
			result = append(result, indexMap[Q[i]]-indexMap[P[i]]+1)
		} else {
			result = append(result, indexMap[Q[i]]-indexMap[P[i]])
		}
	}
	return result
}

// TODO:
func Solution2(N int, P []int, Q []int) []int {
	prime := make([]int, N+1)
	i := 2
	for i*i <= N {
		if prime[i] == 0 {
			k := i * i
			for k <= N {
				if prime[k] == 0 {
					prime[k] = i
				}
				k = k + i
			}
		}
		i++
	}
	// fmt.Println("prime: ", prime)

	// 當前一共有幾個半質數
	semiprime := make([]int, N+1)
	for i := 1; i < len(prime); i++ {
		p := prime[i]
		if p == 0 {
			semiprime[i] = semiprime[i-1]
			continue
		}
		if prime[i/p] == 0 {
			semiprime[i] = semiprime[i-1] + 1
		} else {
			semiprime[i] = semiprime[i-1]
		}
	}
	// fmt.Println("semiprime: ", semiprime)

	result := make([]int, len(P))
	for i, p := range P {
		result[i] = semiprime[Q[i]] - semiprime[p-1]
	}
	return result
}

```