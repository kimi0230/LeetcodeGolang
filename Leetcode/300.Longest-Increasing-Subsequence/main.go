package longestincreasingsubsequence

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//  DP 解法 O(n^2)
func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for idx := 0; idx < len(dp); idx++ {
		dp[idx] = 1
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// 找到前面比現在結尾還小的子序列
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	fmt.Println(dp)
	return res
}

func LengthOfLIS2(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	res := 0

	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[i-1] > nums[j-1] {
				// 找到前面比現在結尾還小的子序列
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] = dp[i] + 1
		res = max(res, dp[i])
	}
	fmt.Println(dp)
	return res
}
