package longestpalindrome

func LongestPalindrome(s string) int {
	counter := make(map[rune]int)
	for _, r := range s {
		counter[r]++
	}

	result := 0

	for _, v := range counter {
		result += v / 2 * 2

		// 字符出现了奇数次，我们可以選擇其中一個, 放在回文串的中間，這可以貢獻一個長度
		if result%2 == 0 && v%2 == 1 {
			result++
		}
	}

	return result
}
