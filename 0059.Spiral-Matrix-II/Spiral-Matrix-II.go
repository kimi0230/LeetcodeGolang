package spiralmatrixii

func GenerateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	left, right, top, botton := 0, n-1, 0, n-1
	num := 1
	target := n * n

	for num <= target {
		// 上層  left to right, 上層邊界++
		for i := left; i <= right; i++ {
			result[top][i] = num
			num++
		}
		top++

		// 右層 top to botton , 右層邊界--
		for i := top; i <= botton; i++ {
			result[i][right] = num
			num++
		}
		right--

		// 下層 right to left , 下層邊界--
		for i := right; i >= left; i-- {
			result[botton][i] = num
			num++
		}
		botton--

		// 左層  botton to top, 左層邊界++
		for i := botton; i >= top; i-- {
			result[i][left] = num
			num++
		}
		left++
	}

	return result
}
