package floodfill

// DFS
func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	oldColor := image[sr][sc]
	if color == oldColor {
		return image
	}
	dfsfill(image, sr, sc, oldColor, color)
	return image
}

func dfsfill(image [][]int, row, col, oldColor, newColor int) {
	// Check if the current pixel is out of bounds or does not have the old color
	if row < 0 || row >= len(image) || col < 0 || col >= len(image[0]) || image[row][col] != oldColor {
		return
	}

	// Update the current pixel with the new color
	image[row][col] = newColor

	// Recursively perform flood fill on the adjacent pixels
	dfsfill(image, row-1, col, oldColor, newColor) // Up
	dfsfill(image, row+1, col, oldColor, newColor) // Down
	dfsfill(image, row, col-1, oldColor, newColor) // Left
	dfsfill(image, row, col+1, oldColor, newColor) // Right
}

type Point struct {
	row, col int
}

func FloodFillBFS(image [][]int, sr int, sc int, newColor int) [][]int {
	// Check if the starting point is out of bounds or already has the new color
	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[0]) || image[sr][sc] == newColor {
		return image
	}

	// Get the old color at the starting point
	oldColor := image[sr][sc]

	// Create a queue and enqueue the starting point
	queue := []Point{{sr, sc}}

	// Define the directions (up, down, left, right)
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// Perform BFS
	for len(queue) > 0 {
		// Dequeue a point from the queue
		point := queue[0]
		queue = queue[1:]

		// Update the point with the new color
		image[point.row][point.col] = newColor

		// Explore the neighboring pixels
		for _, dir := range directions {
			newRow, newCol := point.row+dir[0], point.col+dir[1]

			// Check if the neighboring pixel is within bounds and has the old color
			if newRow >= 0 && newRow < len(image) && newCol >= 0 && newCol < len(image[0]) && image[newRow][newCol] == oldColor {
				// Enqueue the neighboring pixel
				queue = append(queue, Point{newRow, newCol})
			}
		}
	}

	return image
}
