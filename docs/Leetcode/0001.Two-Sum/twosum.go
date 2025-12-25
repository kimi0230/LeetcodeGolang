package twosum

// 時間複雜 O(n^2), 空間複雜 O(1)
func Twosum(nums []int, target int) []int {
	for i, _ := range nums {
		for j := i + 1; j < len(nums); j++ {
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func Twosum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if idx, ok := m[target-v]; ok {
			return []int{idx, i}
		}
		m[v] = i
	}
	return []int{0, 0}
}

// 如果nums是有序 可以使用左右指針
// func Twosum3(nums []int, target int) []int {
// 	left, right := 0, len(nums)-1
// 	sort.Ints(nums)

// 	for left < right {
// 		sum := nums[left] + nums[right]
// 		if sum == target {
// 			return []int{left, right}
// 		} else if sum < target {
// 			left++
// 		} else if sum > target {
// 			right--
// 		}
// 	}
// 	return []int{0, 0}
// }
