# [CountFactors](https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/count_factors/)
Count factors of given number n.

A positive integer D is a factor of a positive integer N if there exists an integer M such that N = D * M.

For example, 6 is a factor of 24, because M = 4 satisfies the above condition (24 = 6 * 4).

Write a function:

func Solution(N int) int

that, given a positive integer N, returns the number of its factors.

For example, given N = 24, the function should return 8, because 24 has 8 factors, namely 1, 2, 3, 4, 6, 8, 12, 24. There are no other factors of 24.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..2,147,483,647].

Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.
## 題目大意
找出該數的因子個數

## 解題思路
尋遍該數字平方根的整數, 每次可以獲得2個因子

## 來源
https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/count_factors/