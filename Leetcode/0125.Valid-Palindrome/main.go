package validpalindrome

import (
	"strings"
	"unicode"
)

// 時間複雜 O(), 空間複雜 O()
func IsPalindrome(s string) bool {
	// 將字符轉成小寫,
	s = strings.ToLower(s)

	// 使用雙指針, 一左一右相向而行, 判斷回文
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isChar(s[left]) {
			left++
		}
		for left < right && !isChar(s[right]) {
			right--
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isChar(c byte) bool {
	if unicode.IsLetter(rune(c)) || unicode.IsDigit(rune(c)) {
		return true
	}
	return false
}

func isCharAZ09(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}

func IsPalindromeStrBuilder(s string) bool {
	s = strings.ToLower(s)
	strLen := len(s)
	// 將字符轉成小寫, 並濾掉空格和標點符號等字符
	var sb strings.Builder
	for i := 0; i < strLen; i++ {
		c := rune(s[i])
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			sb.WriteRune(c)
		}
	}

	sbStr := sb.String()

	// 使用雙指針, 一左一右相向而行, 判斷回文
	left, right := 0, len(sbStr)-1
	for left < right {
		if sbStr[left] != sbStr[right] {
			return false
		}
		left++
		right--
	}
	return true
}
