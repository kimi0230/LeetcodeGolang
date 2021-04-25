# [PassingCars](https://app.codility.com/programmers/lessons/5-prefix_sums/passing_cars/)

A non-empty array A consisting of N integers is given. The consecutive(連續) elements of array A represent consecutive cars on a road.

Array A contains only 0s and/or 1s:

0 represents a car traveling east,
1 represents a car traveling west.
The goal is to count passing cars. We say that a pair of cars (P, Q), where 0 ≤ P < Q < N, is passing when P is traveling to the east and Q is traveling to the west.

For example, consider array A such that:

  A[0] = 0 // no.0 car trave to east
  A[1] = 1 // no.1 car trave to west
  A[2] = 0 // no.2 car trave to east
  A[3] = 1 // no.3 car trave to west
  A[4] = 1 // no.4 car trave to west
We have five pairs of passing cars: (0, 1), (0, 3), (0, 4), (2, 3), (2, 4).

Write a function:

func Solution(A []int) int

that, given a non-empty array A of N integers, returns the number of pairs of passing cars.

The function should return −1 if the number of pairs of passing cars exceeds 1,000,000,000.

For example, given:

  A[0] = 0
  A[1] = 1
  A[2] = 0
  A[3] = 1
  A[4] = 1
the function should return 5, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer that can have one of the following values: 0, 1.

Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
每台車的車號為A[]的index, 0 代表往東, 1代表往西, 向東的車號需要小於向西的車號. 找出會車的幾種可能性

## 解題思路
每一個向東走的車, 都會跟向西的配對. 當遇到向西時組合+1.
所以車號0可以跟所有大於0的向西車配對. 車號2跟所有大於0的向西車配對
1號車前面只有車號0這選擇. 車號3跟4有車號0跟2這兩個選擇. 所以是1+2*2=5

## 來源
https://app.codility.com/programmers/lessons/5-prefix_sums/passing_cars/