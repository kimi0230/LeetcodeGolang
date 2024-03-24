package searcha2dmatrix

// 時間複雜 O(), 空間複雜 O()
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m <= 0 {
		return false
	}
	n := len(matrix[0])

	left, right := 0, m*n-1
	for left <= right {
		mid := int(uint(right+left) >> 1) //left + (right-left)/2
		val := getMatrix(matrix, mid)
		if val == target {
			return true
		} else if val < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}

// 將 mxn的二維陣列 映射成一維陣列
func getMatrix(matrix [][]int, index int) int {
	n := len(matrix[0])
	i, j := index/n, index%n
	return matrix[i][j]
}
