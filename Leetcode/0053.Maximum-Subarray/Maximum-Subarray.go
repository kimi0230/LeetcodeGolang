package maximumsubarray

// MaxSubArrayDP : DP (dynamic programming)
func MaxSubArrayDP(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp, res := make([]int, len(nums)), nums[0]
	dp[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			// 前一個和是正的 繼續加下去
			dp[i] = nums[i] + dp[i-1]
		} else {
			// 前一個和是小於等於0 直接拿現在值取代
			dp[i] = nums[i]
		}
		res = max(res, dp[i])
	}
	return res
}

// MaxSubArray1 : 模擬, 比DP快
func MaxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	maxSum := 0
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp += nums[i]
		if tmp > maxSum {
			maxSum = tmp
		}
		if tmp < 0 {
			tmp = 0
		}
	}
	return maxSum
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

//TODO: 分治法, 這個分治方法類似於「線段樹求解最長公共上升子序列問題」的pushUp操作
