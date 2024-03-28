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
