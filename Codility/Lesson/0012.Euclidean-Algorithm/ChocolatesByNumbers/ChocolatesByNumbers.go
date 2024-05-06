package chocolatesbynumbers

func gcd(N int, M int) int {
	if N%M == 0 {
		return M
	} else {
		return gcd(M, N%M)
	}
}

/*
可以吃到的巧克力的數量就是總的巧克力顆數 N 除以 N 和 M 的最大公因數
計算 N和M的最大公因數P, N除以P得到商即為答案
O(log(N + M))
*/
func Solution(N int, M int) int {
	return N / gcd(N, M)
}

/*
Task Score 75%
Correctness 100%
Performance 50%
input (947853, 4453) the solution exceeded the time limit.
從0號開始吃, 下一個號碼+M-1號
*/
func SolutionBurst(N int, M int) int {
	eaten := make(map[int]struct{})
	eatCount := 0

	if N == 1 || M == 1 {
		return N
	}

	for {
		sumNum := eatCount * M
		startNum := sumNum % N

		if _, ok := eaten[startNum]; !ok {
			eaten[startNum] = struct{}{}
			eatCount++
		} else {
			// 找到已吃過的巧克力
			break
		}
	}
	return eatCount
}
