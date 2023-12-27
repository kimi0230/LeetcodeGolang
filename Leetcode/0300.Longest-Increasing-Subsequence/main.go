package longestincreasingsubsequence

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// DP 解法 O(n^2)
func LengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for idx := 0; idx < len(dp); idx++ {
		dp[idx] = 1
	}

	res := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// 找到下一個數值比現在的大
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	// fmt.Println(dp)
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
	// fmt.Println(dp)
	return res
}

// DP + 二分搜尋:patience sorting. O(nlogn)
func LengthOfLISPatience(nums []int) int {
	top := make([]int, len(nums))

	// 牌堆數初始化為0
	piles := 0

	for i := 0; i < len(nums); i++ {
		poker := nums[i]

		// 搜尋左側邊界的二元搜尋
		left, right := 0, piles
		// fmt.Printf("i:%d\tL:%d\tR:%d\n", i, left, right)
		for left < right {
			mid := left + (right-left)/2
			if top[mid] > poker {
				// 現在的牌比堆小, 縮小範圍
				right = mid
			} else if top[mid] < poker {
				// 現在的牌比堆大
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 沒有找到堆, 建立一個新的堆
		if left == piles {
			piles++
		}
		// 再把這張牌放在堆的頂端
		top[left] = poker
	}
	return piles
}
