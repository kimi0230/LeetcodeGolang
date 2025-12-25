package PermCheck

func Solution(A []int) int {
	intMap := make(map[int]bool)

	for _, v := range A {
		if !intMap[v] {
			intMap[v] = true
		} else {
			// 重複出現
			return 0
		}
	}

	for i := 1; i <= len(A); i++ {
		if !intMap[i] {
			return 0
		}
	}
	return 1
}

func Solution2(A []int) int {
	intMap := make(map[int]bool)
	sum := 0
	for _, v := range A {
		if !intMap[v] {
			intMap[v] = true
			sum += v
		} else {
			// 重複出現
			return 0
		}
	}

	if sum == (len(A)+1)*len(A)/2 {
		return 1
	}
	return 0
}
