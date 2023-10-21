package twosumiiinputarrayissorted

// 時間複雜 O(), 空間複雜 O()

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
