package oddoccurrencesinarray

func Solution(A []int) int {
	intMap := make(map[int]int)
	for i := 0; i < len(A); i++ {
		intMap[A[i]] += 1
	}

	for k, v := range intMap {
		if v%2 != 0 {
			return k
		}
	}
	return -1
}

// 所有的整數XOR起來, 若是兩個整數相同XOR得到0, 最後剩下基數次的數字
// 前提只能有一個基數數字
func Solution2(A []int) int {
	result := 0
	for i := 0; i < len(A); i++ {
		result ^= A[i]
	}
	return result
}
