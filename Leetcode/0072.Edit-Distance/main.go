package editdistance

import "fmt"

// 遞迴 (暴力解)
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
			// word1[0..i] 和 word2[0..j]的最小編輯距離等於 word1[0..i-1] 和 word2[0..j-1]
			// 本來就相等所以不需要任何操作
			// 也就是說 dp(i,j)等於 dp(i-1,j-1)
			return dp(i-1, j-1)
		} else {
			return min(
				dp(i, j-1)+1,   // insert: 直接在 word1[i]中插入一個和word2[j]一樣的字符, 那麼word2[j]就被匹配了,往前j, 繼續和i對比, 操作次數+1
				dp(i-1, j)+1,   // delete: 直接把 word1[i] 這個字符串刪除, 往前 i 繼續和 j 對比, 操作次數+1
				dp(i-1, j-1)+1, // replace: 直接把 word1[i] 替換成 word2[j], 這樣他們就匹配了, 同時往前 i, j 繼續對比, 操作次數+1
			)
		}
	}

	return dp(len(word1)-1, len(word2)-1)
}

// Memo優化
func MinDistanceMemo(word1 string, word2 string) int {
	var dp func(int, int) int
	memo := map[string]int{}
	dp = func(i, j int) int {
		key := fmt.Sprintf("%d,%d", i, j)
		// 查詢備忘錄 避免重複
		if _, ok := memo[key]; ok == true {
			return memo[key]
		}

		// base case
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if word1[i] == word2[j] {
			// word1[0..i] 和 word2[0..j]的最小編輯距離等於 word1[0..i-1] 和 word2[0..j-1]
			// 本來就相等所以不需要任何操作
			// 也就是說 dp(i,j)等於 dp(i-1,j-1)
			memo[key] = dp(i-1, j-1)
		} else {
			memo[key] = min(
				dp(i, j-1)+1,   // insert: 直接在 word1[i]中插入一個和word2[j]一樣的字符, 那麼word2[j]就被匹配了,往前j, 繼續和i對比, 操作次數+1
				dp(i-1, j)+1,   // delete: 直接把 word1[i] 這個字符串刪除, 往前 i 繼續和 j 對比, 操作次數+1
				dp(i-1, j-1)+1, // replace: 直接把 word1[i] 替換成 word2[j], 這樣他們就匹配了, 同時往前 i, j 繼續對比, 操作次數+1
			)
		}
		return memo[key]
	}

	return dp(len(word1)-1, len(word2)-1)
}

type Number interface {
	int | int64 | float64
}

func min[T Number](vars ...T) T {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}
