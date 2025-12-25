// 時間複雜度:
// 空間複雜度:
/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	// 找出最大的香蕉堆
	left, right := 1, maxPile(piles)
	for left < right {
		// mid 代表消耗香蕉的速度(k)
		mid := int(uint(left+right) >> 1)
		if executeTime(piles, mid) <= h {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// 假設消耗速度k, 算出要花多少時間
func executeTime(piles []int, k int) int {
	time := 0
	for _, pile := range piles {
		time += pile / k
		if pile%k > 0 {
			time++
		}
		// time += int(math.Ceil(float64(pile) / float64(k)))
	}
	return time
}

func maxPile(piles []int) int {
	result := 0
	for _, pile := range piles {
		if result < pile {
			result = pile
		}
	}
	return result
}

// @lc code=end

