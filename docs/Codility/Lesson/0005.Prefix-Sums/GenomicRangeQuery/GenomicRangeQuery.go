package GenomicRangeQuery

func Solution(S string, P []int, Q []int) []int {
	A, C, G, T := prefixSums(S)
	result := make([]int, len(P))
	/*
		// fmt.Println("A: ", A)
		// fmt.Println("C: ", C)
		// fmt.Println("G: ", G)
		// fmt.Println("T: ", T)
		idx  0 1 2 3 4 5 6 7
		S:  [C A G C C T A]
		A:  [0 0 1 1 1 1 1 2]
		C:  [0 1 1 1 2 3 3 3]
		G:  [0 0 0 1 1 1 1 1]
		T:  [0 0 0 0 0 0 1 1]
		P:  [2 5 0]
		Q:  [4 5 6]
	*/

	for k, _ := range P {
		// 判斷 A[end of slice]-A[Begin of Slice]是否大於零 即可判斷是否 A 出現過
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
func prefixSums(S string) ([]int, []int, []int, []int) {
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

func inLoop(arr string, s string) bool {
	for _, v := range arr {
		if string(v) == s {
			return true
		}
	}
	return false
}

// O(N * M)
func SolutionBurst(S string, P []int, Q []int) []int {
	result := make([]int, len(P))
	for i := 0; i < len(P); i++ {
		tmp := S[P[i] : Q[i]+1]
		if inLoop(tmp, "A") {
			result[i] = 1
		} else if inLoop(tmp, "C") {
			result[i] = 2
		} else if inLoop(tmp, "G") {
			result[i] = 3
		} else if inLoop(tmp, "T") {
			result[i] = 4
		}
	}
	return result
}

/*
def solutionBySlice(S, P, Q):
  result = []
  length = len(P)
  for i in range(length):
    temp = (S[P[i]:Q[i]+1])
    if "A" in temp:
      result.append(1)
    elif "C" in temp:
      result.append(2)
    elif "G" in temp:
      result.append(3)
    elif "T" in temp:
      result.append(4)
  return result
*/
