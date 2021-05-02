# [Dominator](https://app.codility.com/programmers/lessons/8-leader/dominator/)
Find an index of an array such that its value occurs at more than half of indices in the array.

An array A consisting of N integers is given. The dominator of array A is the value that occurs in more than half of the elements of A.

For example, consider array A such that

 A[0] = 3    A[1] = 4    A[2] =  3
 A[3] = 2    A[4] = 3    A[5] = -1
 A[6] = 3    A[7] = 3
The dominator of A is 3 because it occurs in 5 out of 8 elements of A (namely in those with indices 0, 2, 4, 6 and 7) and 5 is more than a half of 8.

Write a function

func Solution(A []int) int

that, given an array A consisting of N integers, returns index of any element of array A in which the dominator of A occurs. The function should return −1 if array A does not have a dominator.

For example, given array A such that

 A[0] = 3    A[1] = 4    A[2] =  3
 A[3] = 2    A[4] = 3    A[5] = -1
 A[6] = 3    A[7] = 3
the function may return 0, 2, 4, 6 or 7, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer within the range [−2,147,483,648..2,147,483,647].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
返回Array中的支配數. A的支配數是3，因為它出現在A的8個元素中的5個元素中(index為0、2、4、6和7). 而5是8的一半以上
可以返回 0,2,4,6,7中的任一數

## 解題思路
用map紀錄每筆數出現次數. 取最大次數看是否有超過一半以上.
是的話返回此數任一個index, 反之返回-1

## 來源
https://app.codility.com/programmers/lessons/8-leader/dominator/