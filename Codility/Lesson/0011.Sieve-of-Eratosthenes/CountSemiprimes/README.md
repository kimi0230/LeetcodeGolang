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
先把半質數列表計算出來. 
(半質數:兩個質數的乘積所得的自然數我們稱之為半質數ㄝ開始的幾個半質數是4, 6, 9, 10, 14, 15, 21, 22, 25, 26, ... （OEIS中的數列A001358）它們包含1及自己在內合共有3或4個因數)


## 來源
* https://app.codility.com/programmers/lessons/11-sieve_of_eratosthenes/count_semiprimes/
* https://github.com/Luidy/codility-golang/blob/master/Lesson11/02_countSemiprimes.go
* https://github.com/Anfany/Codility-Lessons-By-Python3/blob/master/L11_Sieve%20of%20Eratosthenes/11.2%20CountSemiprimes.md