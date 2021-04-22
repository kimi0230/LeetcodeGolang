package GenomicRangeQuery

func Solution(S string, P []int, Q []int) []int {
	A, C, G, T := prefixSSums(S)
	result := make([]int, len(P))
	// fmt.Println(A)
	// fmt.Println(C)
	// fmt.Println(G)
	// fmt.Println(T)
	for k, _ := range P {
		if A[Q[k]+1]-A[P[k]] > 0 {
			result[k] = 1
		} else if C[Q[k]+1]-C[P[k]] > 0 {
			result[k] = 2
		} else if G[Q[k]+1]-G[P[k]] > 0 {
			result[k] = 3
		} else if T[Q[k]+1]-T[P[k]] > 0 {
			result[k] = 4
		}
	}

	return result
}

// 數算從開始到每個固定索引的A,C,G,T個數. 開頭插入0
func prefixSSums(S string) ([]int, []int, []int, []int) {
	n := len(S)
	A := make([]int, n+1)
	C := make([]int, n+1)
	G := make([]int, n+1)
	T := make([]int, n+1)

	for i := 1; i < n+1; i++ {
		s := string(S[i-1])
		A[i] = A[i-1]
		C[i] = C[i-1]
		G[i] = G[i-1]
		T[i] = T[i-1]

		switch s {
		case "A":
			A[i]++
		case "C":
			C[i]++
		case "G":
			G[i]++
		case "T":
			T[i]++
		}

	}

	return A, C, G, T
}
