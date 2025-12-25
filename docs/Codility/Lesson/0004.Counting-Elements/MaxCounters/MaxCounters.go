package MaxCounters

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 時間 O(N+M) , 空間 O(N)
func Solution(N int, A []int) []int {
	result := make([]int, N)
	maxNum := 0
	nowMaxNum := 0
	for i := 0; i < len(A); i++ {
		if A[i] > N {
			// 如果A[i] 大於 N 則將計數器中所有的數更新為計數器當前的最大數值
			maxNum = nowMaxNum
		} else {
			// 如果A[i] 小於 N 則將計數器中對應位置的數+1,
			if result[A[i]-1] < maxNum {
				result[A[i]-1] = maxNum
			}
			result[A[i]-1]++

			if nowMaxNum < result[A[i]-1] {
				nowMaxNum = result[A[i]-1]
			}
		}
	}
	for i := 0; i < N; i++ {
		if result[i] < maxNum {
			result[i] = maxNum
		}
	}
	return result
}
