package maxareaofisland

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
	return 0
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
