package findtargetlastindex

// RightBound,  有點類似 nums 大於 target的元素有幾個

// 二分搜尋
func Solution(nums []int, target int) int {

	left, right := 0, len(nums)-1

	for left <= right {
		mid := int(uint(left+right) >> 1)
		if nums[mid] == target {
			if nums[right] == target {
				return right
			} else {
				right--
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func SolutionRecursive(nums []int, target int) int {
	left, right := 0, len(nums)-1
	return findTarget(nums, left, right, target)
}

func findTarget(nums []int, left, right, target int) int {
	if left > right {
		return -1
	}

	mid := int(uint(left+right) >> 1)
	if nums[mid] == target {
		if nums[right] == target {
			return right
		} else {
			return findTarget(nums, mid, right-1, target)
		}
	} else if nums[mid] < target {
		return findTarget(nums, mid+1, right, target)
	} else {
		return findTarget(nums, left, mid-1, target)
	}

}
