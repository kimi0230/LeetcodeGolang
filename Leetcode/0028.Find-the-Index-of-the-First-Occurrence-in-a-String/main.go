package findtheindexofthefirstoccurrenceinastring

// 暴力解
// 時間複雜 O(M*N), 空間複雜 O()
func strStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	index := 0
	for i := 0; i <= (haystackLen - needleLen); i++ {
		j := 0
		for j = 0; j < needleLen; j++ {
			if haystack[i+j] == needle[j] {
				index = i
			} else {
				break
			}
		}
		if j == needleLen {
			return index
		}
	}
	return -1
}

// Slice 解法
func strStrSlice(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)
	if haystackLen == 0 || haystackLen < needleLen {
		return -1
	}
	if needleLen == 0 {
		return 0
	}
	for i := 0; i <= (haystackLen - needleLen); i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}
