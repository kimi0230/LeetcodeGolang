package MissingInteger

func Solution(A []int) int {
	smallNum := 1
	intMap := make(map[int]bool)

	// 將出現的數字塞入map
	for _, v := range A {
		if v > 0 && !intMap[v] {
			intMap[v] = true
		}
	}

	for i := 1; i <= len(intMap); i++ {
		if !intMap[i] {
			// 此正整數沒在map找到
			return i
		}
		smallNum = i + 1
	}

	return smallNum
}
