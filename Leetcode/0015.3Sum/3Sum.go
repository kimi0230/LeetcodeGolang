package threesum

import (
	"sort"
)

// ThreeSumBurst : 暴力解 : O(n^3)
func ThreeSumBurst(nums []int) [][]int {
	result := [][]int{}
	sort.Ints(nums) // O(n log n)
	for i := 0; i < len(nums); i++ {
		// 需要跟上一次不同
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums); j++ {
			// 需要跟上一次不同
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return result
}

/*
ThreeSumDoublePoint : 最佳解, 排序 + 雙指針法 (滑動視窗) O(n^2)
1. 特判，對於數組長度 n，如果數組為 null 或者數組長度小於 3，返回 []。
2. 對數組進行排序。
3. 遍歷排序後數組：
	* 對於重複元素：跳過，避免出現重複解
	* 令左指針 L=i+1，右指針 R=n−1，當 L<R 時，執行循環：
		* 當nums[i]+nums[L]+nums[R]==0，執行循環，判斷左界和右界是否和下一位置重複，
		去除重複解。並同時將 L,R 移到下一位置，尋找新的解
		* 若和大於 0，說明 nums[R] 太大，R 左移
		* 若和小於 0，說明 nnums[L] 太小，L 右移
*/
func ThreeSumDoublePoint(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}
	sort.Ints(nums) // O(n log n)
	start, end, addNum := 0, 0, 0
	for i := 1; i < len(nums)-1; i++ {
		start, end = 0, len(nums)-1
		if i > 1 && nums[i] == nums[i-1] {
			// 去掉重複
			start = i - 1
		}
		for start < i && end > i {
			if start > 0 && nums[start] == nums[start-1] {
				// 去掉重複
				start++
				continue
			}
			if end < (len(nums)-1) && nums[end] == nums[end+1] {
				// 去掉重複
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[i]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[i], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
