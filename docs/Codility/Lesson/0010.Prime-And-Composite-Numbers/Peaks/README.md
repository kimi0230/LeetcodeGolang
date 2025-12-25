# [Peaks](https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/peaks/)
Divide an array into the maximum number of same-sized blocks, each of which should contain an index P such that A[P - 1] < A[P] > A[P + 1].

A non-empty array A consisting of N integers is given.

A peak is an array element which is larger than its neighbors. More precisely, it is an index P such that 0 < P < N − 1,  A[P − 1] < A[P] and A[P] > A[P + 1].

For example, the following array A:

    A[0] = 1
    A[1] = 2
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2
has exactly three peaks: 3, 5, 10.

We want to divide this array into blocks containing the same number of elements. More precisely, we want to choose a number K that will yield the following blocks:

A[0], A[1], ..., A[K − 1],
A[K], A[K + 1], ..., A[2K − 1],
...
A[N − K], A[N − K + 1], ..., A[N − 1].
What's more, every block should contain at least one peak. Notice that extreme elements of the blocks (for example A[K − 1] or A[K]) can also be peaks, but only if they have both neighbors (including one in an adjacent blocks).

The goal is to find the maximum number of blocks into which the array A can be divided.

Array A can be divided into blocks as follows:

one block (1, 2, 3, 4, 3, 4, 1, 2, 3, 4, 6, 2). This block contains three peaks.
two blocks (1, 2, 3, 4, 3, 4) and (1, 2, 3, 4, 6, 2). Every block has a peak.
three blocks (1, 2, 3, 4), (3, 4, 1, 2), (3, 4, 6, 2). Every block has a peak. Notice in particular that the first block (1, 2, 3, 4) has a peak at A[3], because A[2] < A[3] > A[4], even though A[4] is in the adjacent block.
However, array A cannot be divided into four blocks, (1, 2, 3), (4, 3, 4), (1, 2, 3) and (4, 6, 2), because the (1, 2, 3) blocks do not contain a peak. Notice in particular that the (4, 3, 4) block contains two peaks: A[3] and A[5].

The maximum number of blocks that array A can be divided into is three.

Write a function:

func Solution(A []int) int

that, given a non-empty array A consisting of N integers, returns the maximum number of blocks into which A can be divided.

If A cannot be divided into some number of blocks, the function should return 0.

For example, given:

    A[0] = 1
    A[1] = 2
    A[2] = 3
    A[3] = 4
    A[4] = 3
    A[5] = 4
    A[6] = 1
    A[7] = 2
    A[8] = 3
    A[9] = 4
    A[10] = 6
    A[11] = 2
the function should return 3, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [0..1,000,000,000].

Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
將 array 分成同樣長度的區塊, 買個區塊至少包含一個peak.

## 解題思路
先找出所有peak的index 寫入peaks array.
從peaks的長度開始往下找, 將 A 拆成區塊,
每個區塊檢查是否有有找到peak

## 來源
* https://app.codility.com/programmers/lessons/10-prime_and_composite_numbers/peaks/
* https://www.martinkysel.com/codility-peaks-solution/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0010.Prime-And-Composite-Numbers/Peaks/Peaks.go


```go
package peaks

/*
expected worst-case time complexity is O(N*log(log(N)));
expected worst-case space complexity is O(N)
*/
func Solution(A []int) int {
	// 先找出peaks
	peaks := []int{}
	for i := 1; i < len(A)-1; i++ {
		if A[i-1] < A[i] && A[i] > A[i+1] {
			peaks = append(peaks, i)
		}
	}

	if len(peaks) < 0 {
		return 0
	} else if len(peaks) == 1 {
		return 1
	}

	for size := len(peaks); size > 0; size-- {
		if len(A)%size == 0 {
			// 每個區塊的size
			blockSize := len(A) / size
			found := make(map[int]bool, size)
			foundCnt := 0
			for _, peak := range peaks {
				// 檢查每個區塊是否有找到 peak
				blockNr := peak / blockSize
				if ok := found[blockNr]; !ok {
					found[blockNr] = true
					foundCnt++
				}
			}
			if foundCnt == size {
				return size
			}
		}
	}
	return 0
}

/*
def solution(A):
    peaks = []

    for idx in range(1, len(A)-1):
        if A[idx-1] < A[idx] > A[idx+1]:
            peaks.append(idx)

    if len(peaks) == 0:
        return 0

    for size in range(len(peaks), 0, -1):
        if len(A) % size == 0:
            block_size = len(A) // size
            found = [False] * size
            found_cnt = 0
            for peak in peaks:
                block_nr = peak//block_size
                if found[block_nr] == False:
                    found[block_nr] = True
                    found_cnt += 1

            if found_cnt == size:
                return size

    return 0
*/
```