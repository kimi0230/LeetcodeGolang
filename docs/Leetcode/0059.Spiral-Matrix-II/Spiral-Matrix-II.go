package spiralmatrixii

// GenerateMatrix : 按層模擬, 時間複雜 O(n^2) 空間複雜 O(1)
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

// 模擬 : O(n)
// https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0059.Spiral-Matrix-II/
func GenerateMatrix2(n int) [][]int {
	if n == 0 {
		return [][]int{}
	}
	if n == 1 {
		return [][]int{{1}}
	}
	result, visit, round, x, y, spDir := make([][]int, n), make([][]int, n), 0, 0, 0, [][]int{
		{0, 1},  // 朝右
		{1, 0},  // 朝下
		{0, -1}, // 朝左
		{-1, 0}, // 朝上
	}
	for i := 0; i < n; i++ {
		visit[i] = make([]int, n)
		result[i] = make([]int, n)
	}
	visit[x][y] = 1
	result[x][y] = 1

	for i := 0; i < n*n; i++ {
		x += spDir[round%4][0]
		y += spDir[round%4][1]
	}

	return result
}
