package twosumiiinputarrayissorted

// 雙指針. 時間複雜度： O(n)，其中 n 是陣列的長度。 兩個指標移動的總次數最多為 n 次。
// 空間複雜度： O(1)

func TwoSum(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		}
		if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}

	return nil
}

// 使用 O(n^2) 的時間複雜度和 O(1)的空間複雜度暴力
// 或者藉助哈希表使用 O(n) 的時間複雜度和 O(n) 的空間複雜度求解

func TwoSum2(numbers []int, target int) []int {
	m := make(map[int]int, len(numbers))

	for i, v := range numbers {
		if j, ok := m[target-v]; ok {
			return []int{j + 1, i + 1}
		}
		m[v] = i
	}
	return nil
}
