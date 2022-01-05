# [AbsDistinct](https://app.codility.com/programmers/lessons/15-caterpillar_method/abs_distinct/)

A non-empty array A consisting of N numbers is given. The array is sorted in non-decreasing order. The absolute distinct count of this array is the number of distinct absolute values among the elements of the array.

For example, consider array A such that:

  A[0] = -5
  A[1] = -3
  A[2] = -1
  A[3] =  0
  A[4] =  3
  A[5] =  6
The absolute distinct count of this array is 5, because there are 5 distinct absolute values among the elements of this array, namely 0, 1, 3, 5 and 6.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N numbers, returns absolute distinct count of array A.

For example, given array A such that:

  A[0] = -5
  A[1] = -3
  A[2] = -1
  A[3] =  0
  A[4] =  3
  A[5] =  6
the function should return 5, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−2,147,483,648..2,147,483,647];
array A is sorted in non-decreasing order.
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.


"有序數組中絕對值不同的數的個數"指的是，一個已經排好序的整數數組中絕對值不相同的數字的個數，

例如：

A[0] = -5    A[1] = -3    A[2] = -1
A[3] =  0    A[4] =  3    A[5] =  6
絕對值不同的數的個數為 5， 因為其中有 5 個不同的絕對值： 0, 1, 3, 5, 6

編寫一個函數：

func Solution(A []int) int

請返回給定有序數組中絕對值不同的數的個數。

例如，給出數組A：

A[0] = -5    A[1] = -3    A[2] = -1
A[3] =  0    A[4] =  3    A[5] =  6
函數應返回5。

假定:

N 是 [1..100,000] 內的 整數;
數組 A 每個元素是取值範圍 [−2,147,483,648..2,147,483,647] 內的 整數 ;
數組 A 是 非-遞增 序列.

## 題目大意

## 解題思路


## 來源
* https://app.codility.com/programmers/lessons/15-caterpillar_method/abs_distinct/