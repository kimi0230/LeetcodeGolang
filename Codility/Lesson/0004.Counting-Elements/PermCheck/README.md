# [PermCheck](https://app.codility.com/programmers/lessons/4-counting_elements/perm_check/)
Check whether array A is a permutation.

A non-empty array A consisting of N integers is given.

A permutation(排列) is a sequence containing each element from 1 to N once, and only once.

For example, array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
is a permutation, but array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
is not a permutation, because value 2 is missing.

The goal is to check whether array A is a permutation.

Write a function:

func Solution(A []int) int

that, given an array A, returns 1 if array A is a permutation and 0 if it is not.

For example, given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
    A[3] = 2
the function should return 1.

Given array A such that:

    A[0] = 4
    A[1] = 1
    A[2] = 3
the function should return 0.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [1..1,000,000,000].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
如果是連續排列的array 回傳 1 反之回傳1

## 解題思路
類似lesson 4的MissingInteger. 先將現有的直寫入到map. 除了檢查是否有重複數字出現外,順便將總和算起來
最後檢查總時對不對

## 來源
https://app.codility.com/programmers/lessons/4-counting_elements/perm_check/