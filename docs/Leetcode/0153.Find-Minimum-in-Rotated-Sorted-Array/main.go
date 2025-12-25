package findminimuminrotatedsortedarray

// 時間複雜 O(logN), 空間複雜 O(1)
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := int(uint(left+right) >> 1)
		if nums[mid] >= nums[left] {
			// mid 一定不是最小值了
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}
