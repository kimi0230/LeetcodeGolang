package merging2packages

func getIndicesOfItemWeights(arr []int, limit int) []int {

	m := make(map[int]int)
	for i, v := range arr {
		if _, ok := m[limit-v]; ok {
			return []int{i, m[limit-v]}
		}
		m[v] = i
	}
	return nil
}
