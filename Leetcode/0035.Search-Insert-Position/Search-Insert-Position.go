package searchinsertposition

// 暴力解 時間複雜  O(n) 空間複雜 O(1)
func SearchInsertBurst(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

//二分法 時間複雜 O(log n) 空間複雜 O(1)
func SearchInsertBisection(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// / 防止溢出 同(left + right)/2
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	// 分別處理如下四種情況
	// targe在所有元素之前 [0, -1]
	// targe等於數組中某一個元素 return middle;
	// targe在數組中的位置 [left, right]，return right + 1
	// targe在數組所有元素之後的情況 [left, right]， return right + 1
	return right + 1
}

//二分法 時間複雜 O(log n) 空間複雜 O(1)
func SearchInsertBisection2(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		// / 防止溢出 同(left + right)/2
		mid := left + (right-left)>>1
		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
