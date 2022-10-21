package longestpalindromicsubstring

func LongestPalindromeSubstring(s string) string {

	return ""
}

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func max[T numbers](a T, b T) T {
	if a > b {
		return a
	}
	return b
}
