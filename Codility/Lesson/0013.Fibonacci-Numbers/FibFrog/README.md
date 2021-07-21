# [FibFrog](https://app.codility.com/programmers/lessons/13-fibonacci_numbers/fib_frog/)

The Fibonacci sequence is defined using the following recursive formula:

    F(0) = 0
    F(1) = 1
    F(M) = F(M - 1) + F(M - 2) if M >= 2
A small frog wants to get to the other side of a river. The frog is initially located at one bank of the river (position −1) and wants to get to the other bank (position N). The frog can jump over any distance F(K), where F(K) is the K-th Fibonacci number. Luckily, there are many leaves on the river, and the frog can jump between the leaves, but only in the direction of the bank at position N.

The leaves on the river are represented in an array A consisting of N integers. Consecutive(連續的) elements of array A represent consecutive positions from 0 to N − 1 on the river. Array A contains only 0s and/or 1s:

0 represents a position without a leaf;
1 represents a position containing a leaf.
The goal is to count the minimum number of jumps in which the frog can get to the other side of the river (from position −1 to position N). The frog can jump between positions −1 and N (the banks of the river) and every position containing a leaf.

For example, consider array A such that:

    A[0] = 0
    A[1] = 0
    A[2] = 0
    A[3] = 1
    A[4] = 1
    A[5] = 0
    A[6] = 1
    A[7] = 0
    A[8] = 0
    A[9] = 0
    A[10] = 0
The frog can make three jumps of length F(5) = 5, F(3) = 2 and F(5) = 5.

Write a function:

func Solution(A []int) int

that, given an array A consisting of N integers, returns the minimum number of jumps by which the frog can get to the other side of the river. If the frog cannot reach the other side of the river, the function should return −1.

For example, given:

    A[0] = 0
    A[1] = 0
    A[2] = 0
    A[3] = 1
    A[4] = 1
    A[5] = 0
    A[6] = 1
    A[7] = 0
    A[8] = 0
    A[9] = 0
    A[10] = 0
the function should return 3, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..100,000];
each element of array A is an integer that can have one of the following values: 0, 1.
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.


## 題目大意
一只小青蛙想到對岸。它開始位於河的另一邊 位置-1, 想要到對面的河岸 位置N . 青蛙可以跳任意距離 F(K). 其中F(K)是第K個斐波那契數.
且河上有許多樹葉 A[0] = 0 代表位置 0 有樹葉, 1 代表沒樹葉
青蛙可以在樹葉之間跳, 但只能朝河岸 N 的方向跳
找出最小跳的次數
## 解題思路
廣度優先搜尋 (Breadth-First Search, BFS) 問題.
對於河上有樹葉的位置index, 則遍歷比index小的斐波那契數f, 
只要 index - f 這個位置可以達到, 這index的位置就可以經過一次跳躍長度為f

典型的BFS搜索問題。對於河面上有樹葉的位置index，則就要遍歷比index不大的所有斐波那契數f，只要index-f這個位置可以達到，那麼index這個位置就可以經過一次跳躍長度爲f達到。因爲是遍歷，所以就決定了最後得到的是最小次數。

## 來源
* https://app.codility.com/programmers/lessons/13-fibonacci_numbers/fib_frog/