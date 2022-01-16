package binarysearch

import "sort"

func Search(nums []int, target int) int {
	lenght := len(nums)
	if lenght <= 0 {
		return -1
	}
	left, right := 0, lenght-1

	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// 找右邊
			left = mid + 1
		} else if nums[mid] > target {
			// 找左邊
			right = mid - 1
		}
	}
	// 都沒找到
	return -1
}

func Search2(nums []int, target int) int {
	lenght := len(nums)
	if lenght <= 0 {
		return -1
	}
	left, right := 0, lenght-1

	for left <= right {
		// 除以2
		// mid := left + (right-left)>>1
		mid := int(uint(right+left) >> 1)
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// 找右邊
			left = mid + 1
		} else if nums[mid] > target {
			// 找左邊
			right = mid - 1
		}
	}
	// 都沒找到
	return -1
}

// 內建sort
func BinarySearch(nums []int, target int) int {
	length := len(nums)

	index := sort.Search(length, func(i int) bool {
		return nums[i] >= target
	})

	if index < length && nums[index] == target {
		return index
	} else {
		return -1
	}
}

// 使用遞迴
func BinarySearchRecur(nums []int, target int) (index int) {
	return BinarySearchRecursively(nums, target, 0, len(nums)-1)
}

func BinarySearchRecursively(nums []int, target int, start int, end int) int {
	if end < start {
		return -1
	}

	middle := int(uint(end+start) >> 1)

	if nums[middle] == target {
		return middle
	} else if target > nums[middle] {
		return BinarySearchRecursively(nums, target, middle+1, end)
	} else {
		return BinarySearchRecursively(nums, target, start, middle-1)
	}
}
