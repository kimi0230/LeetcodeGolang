package checkpalindrome

func checkPalindrome(inputString string) bool {
	strLen := len(inputString)
	if strLen < 2 {
		return true
	}
	left, right := 0, strLen-1

	for left < right {
		if inputString[left] != inputString[right] {
			return false
		}
		left++
		right--
	}
	return true
}
