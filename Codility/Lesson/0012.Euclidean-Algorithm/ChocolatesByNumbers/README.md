# [ChocolatesByNumbers](https://app.codility.com/programmers/lessons/12-euclidean_algorithm/chocolates_by_numbers/)
There are N chocolates in a circle. Count the number of chocolates you will eat.

Two positive integers N and M are given. Integer N represents the number of chocolates arranged in a circle, numbered from 0 to N − 1.

You start to eat the chocolates. After eating a chocolate you leave only a wrapper.

You begin with eating chocolate number 0. Then you omit(忽略) the next M − 1 chocolates or wrappers on the circle, and eat the following one.

More precisely(恰恰), if you ate chocolate number X, then you will next eat the chocolate with number (X + M) modulo N (remainder of division).

You stop eating when you encounter an empty wrapper.

For example, given integers N = 10 and M = 4. You will eat the following chocolates: 0, 4, 8, 2, 6.

The goal is to count the number of chocolates that you will eat, following the above rules.

Write a function:

func Solution(N int, M int) int

that, given two positive integers N and M, returns the number of chocolates that you will eat.

For example, given integers N = 10 and M = 4. the function should return 5, as explained above.

Write an efficient algorithm for the following assumptions:

N and M are integers within the range [1..1,000,000,000].

Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
N塊巧克力,如果吃的是X號 下一個是吃 (X + M) modulo N 號
總共可以吃幾顆.

## 解題思路
方法ㄧ: 從0號開始吃, 下一個號碼+M-1號. 迴圈去跑
方法二: 可以吃到的巧克力的數量就是總的巧克力顆數 N 除以 N 和 M 的最大公因數. 計算 N和M的最大公因數P, N除以P得到商即為答案
## 來源
https://app.codility.com/programmers/lessons/12-euclidean_algorithm/chocolates_by_numbers/