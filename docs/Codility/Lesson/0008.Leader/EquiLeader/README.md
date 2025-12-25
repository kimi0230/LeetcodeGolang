# [EquiLeader](https://app.codility.com/programmers/lessons/8-leader/equi_leader/)
Find the index S such that the leaders of the sequences A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N - 1] are the same.


A non-empty array A consisting of N integers is given.

The leader of this array is the value that occurs in more than half of the elements of A.

An equi leader is an index S such that 0 ≤ S < N − 1 and two sequences A[0], A[1], ..., A[S] and A[S + 1], A[S + 2], ..., A[N − 1] have leaders of the same value.

For example, given array A such that:

    A[0] = 4
    A[1] = 3
    A[2] = 4
    A[3] = 4
    A[4] = 4
    A[5] = 2
we can find two equi leaders:

* 0, because sequences: (4) and (3, 4, 4, 4, 2) have the same leader, whose value is 4.
* 2, because sequences: (4, 3, 4) and (4, 4, 2) have the same leader, whose value is 4.
The goal is to count the number of equi leaders.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the number of equi leaders.

For example, given:

    A[0] = 4
    A[1] = 3
    A[2] = 4
    A[3] = 4
    A[4] = 4
    A[5] = 2
the function should return 2, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
選定一個下標, 將一個數組分爲左右兩個子數組, 使得兩個子數組都有相同的leader,
則稱此下標爲EquiLeader, 要求返回給定數組中EquiLeader的個數n.
事實上, 若一個數同時是左子數組的leader, 它必然也是整個數組的leader.

## 解題思路
需要先找出序列中的leader, 記錄其出現的次數.
然後再遍歷整個數組，枚舉分割點，記錄下左子數組leader出現的次數s,
看s與n-s是否能使得leader在左右子數組中仍爲leader.

## 來源
* https://app.codility.com/programmers/lessons/8-leader/equi_leader/
* https://github.com/Anfany/Codility-Lessons-By-Python3/blob/master/L8_Leader/8.1%20EquiLeader.md

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0008.Leader/EquiLeader/EquiLeader.go


```go
package EquiLeader

func Solution(A []int) int {
	leaderDict := make(map[int]int)
	for i := 0; i < len(A); i++ {
		if _, ok := leaderDict[A[i]]; ok {
			leaderDict[A[i]]++
		} else {
			leaderDict[A[i]] = 1
		}
	}

	leader := 0
	times := 0
	for k, v := range leaderDict {
		if v > times {
			times = v
			leader = k
		}
	}

	equiCount := 0
	count := 0 // 超頻數已出現的次數

	for index, v := range A {
		if v == leader {
			count++
		}
		if count > (index+1)/2 && (times-count) > (len(A)-(index+1))/2 {
			equiCount++
		}

	}

	return equiCount
}
```