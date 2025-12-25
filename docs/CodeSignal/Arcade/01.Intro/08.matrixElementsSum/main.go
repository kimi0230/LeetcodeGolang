package matrixelementssum

// 從左上往左下走, 遇到0就結束. 一路由左邊統計到右邊
func matrixElementsSum(matrix [][]int) int {
	sum := 0
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			if matrix[j][i] <= 0 {
				break
			}
			sum += matrix[j][i]
		}
	}
	return sum
}
