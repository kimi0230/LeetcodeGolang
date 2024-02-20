package longestconsecutivesequence

// 時間複雜 O(), 空間複雜 O()
func longestConsecutive(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		m[num] = struct{}{}
	}
	result := 0
	for v := range m {
		// 如果沒找到該數字的前一個數字, 則把該數字刀做連續序列的第一個數
		if _, ok := m[v-1]; !ok {
			length := 1
			for _, exit := m[v+length]; exit; _, exit = m[v+length] {
				length++
			}
			if result < length {
				result = length
			}
		}
	}
	return result
}
