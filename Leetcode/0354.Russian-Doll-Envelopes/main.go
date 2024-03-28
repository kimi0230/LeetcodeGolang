package russiandollenvelopes

import (
	longestincreasingsubsequence "LeetcodeGolang/Leetcode/0300.Longest-Increasing-Subsequence"
	"sort"
)

type sortEnvelopes [][]int

func (s sortEnvelopes) Len() int {
	return len(s)
}

func (s sortEnvelopes) Less(i, j int) bool {
	if s[i][0] == s[j][0] {
		// 遇到w相同的情況, 則按照高度進行降序排序
		return s[i][1] > s[j][1]
	}
	// 對寬度w進行升序排序
	return s[i][0] < s[j][0]
}

func (s sortEnvelopes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func MaxEnvelopes(envelopes [][]int) int {
	// 先對寬度w進行升序排序, 如果遇到w相同的情況, 則按照高度進行降序排序.
	sort.Sort(sortEnvelopes(envelopes))
	// fmt.Println(envelopes)

	// 所有的h最為一個array, 對此array做 最長遞增子序列(Longest Increasing Subsequence, LIS)的長度就是答案
	height := make([]int, len(envelopes))
	for i := 0; i < len(envelopes); i++ {
		height[i] = envelopes[i][1]
	}

	//  DP + 二分搜尋:patience sorting. O(nlogn)
	return longestincreasingsubsequence.LengthOfLISPatience(height)
}

func MaxEnvelopes2(envelopes [][]int) int {
	// 先對寬度w進行升序排序, 如果遇到w相同的情況, 則按照高度進行降序排序.
	sort.Sort(sortEnvelopes(envelopes))
	// fmt.Println(envelopes)

	// 最長遞增子序列(Longest Increasing Subsequence, LIS)的長度就是答案
	dp := []int{}
	for _, e := range envelopes {
		left, right := 0, len(dp)
		for left < right {
			mid := left + (right-left)/2
			if dp[mid] > e[1] {
				// 現在的牌比堆小, 所小範圍
				right = mid
			} else if dp[mid] < e[1] {
				// 現在的牌比堆大
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 沒有找到堆, 建立一個新的堆
		if left == len(dp) {
			dp = append(dp, e[1])
		}

		// 再把這張牌放在堆的頂端
		dp[left] = e[1]
	}

	return len(dp)
}

/*
func Sort(data Interface) {
	n := data.Len()
	quickSort(data, 0, n, maxDepth(n))
}
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int

	// Less reports whether the element with index i
	// must sort before the element with index j.
	//
	// If both Less(i, j) and Less(j, i) are false,
	// then the elements at index i and j are considered equal.
	// Sort may place equal elements in any order in the final result,
	// while Stable preserves the original input order of equal elements.
	//
	// Less must describe a transitive ordering:
	//  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
	//  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
	//
	// Note that floating-point comparison (the < operator on float32 or float64 values)
	// is not a transitive ordering when not-a-number (NaN) values are involved.
	// See Float64Slice.Less for a correct implementation for floating-point values.
	Less(i, j int) bool

	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
*/
