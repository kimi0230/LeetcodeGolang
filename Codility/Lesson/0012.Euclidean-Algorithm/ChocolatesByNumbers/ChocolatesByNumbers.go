package chocolatesbynumbers

/*
Task Score 75%
Correctness 100%
Performance 50%
input (947853, 4453) the solution exceeded the time limit.
*/
func Solution(N int, M int) int {
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
