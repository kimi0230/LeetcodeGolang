package editdistance

// 遞迴
func MinDistance(word1 string, word2 string) int {
	var dp func(int, int) int
	dp = func(i, j int) int {
		// base case
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		} else {
			return min(
				dp(i, j-1)+1,   // insert
				dp(i-1, j)+1,   // delete
				dp(i-1, j-1)+1, // replace
			)
		}
	}

	return dp(len(word1)-1, len(word2)-1)
}

type Number interface {
	int | int64 | float64
}

func min[nums Number](vars ...nums) nums {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
