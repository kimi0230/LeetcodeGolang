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
