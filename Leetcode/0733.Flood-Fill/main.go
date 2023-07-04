package floodfill

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
