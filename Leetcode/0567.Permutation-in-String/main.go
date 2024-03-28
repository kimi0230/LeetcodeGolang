package permutationinstring

func CheckInclusion(s1 string, s2 string) bool {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range s1 {
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		c := rune(s2[right])
		right++

		// 進行窗口內數據的一系列更新
		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				// 該字符長度達到
				valid++
			}
		}

		// fmt.Printf("[%d,%d) \n", left, right)

		// 判斷左視窗是否要收縮
		// for (right - left) >= len(s1)
		if (right - left) >= len(s1) {

			if valid == len(need) {
				// 全找到
				return true
			}
			d := rune(s2[left])
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

// 用 slice 取代 map 來優化
func CheckInclusionSlice(s1 string, s2 string) bool {
	need := [256]int{}
	for _, c := range s1 {
		need[c-'a']++
	}

	left, right := 0, 0
	count := len(s1)
	for right < len(s2) {
		c := s2[right] - 'a'

		if need[c] > 0 {
			// 有找到
			count--
		}
		need[c]--
		right++

		// fmt.Printf("[%d,%d)\n", left, right)
		if count == 0 {
			return true
		}

		// 判斷左視窗是否要收縮
		if (right - left) == len(s1) {
			d := s2[left] - 'a'
			if need[d] >= 0 {
				// 符合預期的長度, 但是卻沒找到預期的結果
				count++
			}
			need[d]++
			left++
		}
	}
	return false
}
