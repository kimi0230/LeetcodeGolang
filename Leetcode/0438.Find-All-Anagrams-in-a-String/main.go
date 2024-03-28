package findallanagramsinastring

func FindAnagrams(s string, p string) []int {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range p {
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	res := []int{} //紀錄結果
	for right < len(s) {
		c := rune(s[right])
		right++

		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// fmt.Printf("[%d,%d) \n", left, right)

		// 判斷左視窗是否收縮, 看看視窗長度是否同要找的字串的長度
		if (right - left) >= len(p) {
			if valid == len(need) {
				// 想要的字元都找到了, 紀錄index
				res = append(res, left)
			}
			d := rune(s[left])
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

// 用 slice 取代 map 來優化
func FindAnagramsSlice(s string, p string) []int {
	need := [256]int{}
	for _, c := range p {
		need[c-'a']++
	}
	left, right := 0, 0
	count := len(p)
	res := []int{}

	for right < len(s) {
		c := s[right] - 'a'
		if need[c] > 0 {
			count--
		}
		need[c]--
		right++

		if count == 0 {
			res = append(res, left)
		}

		if (right - left) >= len(p) {
			d := s[left] - 'a'
			if need[d] >= 0 {
				count++
			}
			need[d]++
			left++
		}
	}
	return res
}
