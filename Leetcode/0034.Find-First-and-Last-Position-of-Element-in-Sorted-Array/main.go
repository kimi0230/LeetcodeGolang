package findfirstandlastpositionofelementinsortedarray

// 時間複雜 O(log(n)), 空間複雜 O(1)
func searchRange(nums []int, target int) []int {
	result := []int{}

	// first index
	left, right := 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right) >> 1) // [0,1] => 0
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			right = mid
		}
	}
	if left == right && nums[left] == target {
		result = append(result, left)
	} else {
		result = append(result, -1)
	}

	// last index
	left, right = 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right+1) >> 1) // [0,1] => 1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid
		}
	}
	if left == right && nums[left] == target {
		result = append(result, left)
	} else {
		result = append(result, -1)
	}
	return result
}
