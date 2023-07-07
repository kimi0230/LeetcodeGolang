package maxareaofisland

var (
	dir = [][]int{
		{-1, 0}, // 上
		{0, 1},  // 右
		{1, 0},  // 下
		{0, -1}, // 左
	}
)

func MaxAreaOfIsland(grid [][]int) int {
	result := 0

	for i, row := range grid {
		for j, col := range row {
			if col == 1 {
				result = max(result, areaOfIsland(grid, i, j))
			}
		}
	}
	return result
}

// DFS
func areaOfIsland(grid [][]int, x, y int) int {
	if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) || grid[x][y] == 0 {
		return 0
	}
	// 設為0，避免重複計算
	grid[x][y] = 0
	total := 1

	// 四個方向查找
	for i := 0; i < 4; i++ {
		nx := x + dir[i][0]
		ny := y + dir[i][1]
		total += areaOfIsland(grid, nx, ny)
	}
	return total
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
