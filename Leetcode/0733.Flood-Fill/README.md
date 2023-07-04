---
title: 0733.Flood Fill
subtitle: "https://leetcode.com/problems/flood-fill/"
date: 2023-01-22T21:20:00+08:00
lastmod: 2023-01-22T21:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0733.Flood Fill"
license: ""
images: []

tags: [LeetCode, Go, Easy, LEETCODETITLE]
categories: [LeetCode]

featuredImage: ""
featuredImagePreview: ""

hiddenFromHomePage: false
hiddenFromSearch: false
twemoji: false
lightgallery: true
ruby: true
fraction: true
fontawesome: true
linkToMarkdown: false
rssFullText: false

toc:
  enable: true
  auto: true
code:
  copy: true
  maxShownLines: 200
math:
  enable: false
  # ...
mapbox:
  # ...
share:
  enable: true
  # ...
comment:
  enable: true
  # ...
library:
  css:
    # someCSS = "some.css"
    # located in "assets/"
    # Or
    # someCSS = "https://cdn.example.com/some.css"
  js:
    # someJS = "some.js"
    # located in "assets/"
    # Or
    # someJS = "https://cdn.example.com/some.js"
seo:
  images: []
  # ...
---
# [0733.Flood Fill](https://leetcode.com/problems/flood-fill/)

## 題目
An image is represented by an `m x n` integer grid image where `image[i][j]` represents the pixel value of the image.

You are also given three integers `sr`,`sc`, and `color`. You should perform a flood fill on the image starting from the pixel `image[sr][sc]`.

To perform a flood fill, consider the starting pixel, plus any pixels connected 4-directionally to the starting pixel of the same color as the starting pixel, plus any pixels connected 4-directionally to those pixels (also with the same color), and so on. Replace the color of all of the aforementioned pixels with `color`.

Return the modified image after performing the flood fill.


Example 1:
![](images/flood1-grid.jpg)

```sh
Input: image = [[1,1,1],[1,1,0],[1,0,1]], sr = 1, sc = 1, color = 2
Output: [[2,2,2],[2,2,0],[2,0,1]]
Explanation: From the center of the image with position (sr, sc) = (1, 1) (i.e., the red pixel), all pixels connected by a path of the same color as the starting pixel (i.e., the blue pixels) are colored with the new color.
Note the bottom corner is not colored 2, because it is not 4-directionally connected to the starting pixel.
```

Example 2:
```sh
Input: image = [[0,0,0],[0,0,0]], sr = 0, sc = 0, color = 0
Output: [[0,0,0],[0,0,0]]
Explanation: The starting pixel is already colored 0, so no changes are made to the image.
```

Constraints:

* m == image.length
* n == image[i].length
* 1 <= m, n <= 50
* 0 <= image[i][j], color < 216
* 0 <= sr < m
* 0 <= sc < n



## 題目大意
有一個以二維整數陣列表示的圖畫，每個整數表示該圖畫的像素值大小，數值在 0 到 65535 之間。
給你一個坐標 (sr, sc) 表示圖像渲染開始的像素值（行 ，列）和一個新的顏色值新顏色，讓您重新上色該幅圖像。

為了完成上顏色工作，從主板坐標開始，記錄主板坐標的上下四個方向上像素值與主板坐標相同的完整像素點，連接再記錄這四個方向上條件符合的像素點與它們對應的四個像素點方向上像素值與主板坐標的連通像素點相同，……，重複該過程。將所有有記錄的像素點的顏色值改為新的顏色值。最後返回經過上顏色渲染後的圖像。

注意：

* image 和 image[0] 的長度在範圍 [1, 50] 內。
* 給出的初始點將滿足 0 <= sr < image.length 和 0 <= sc < image[0].length。
* image[i][j] 和 newColor 表示的顏色值在範圍 [0, 65535] 內。

## 解題思路
一個給出二維的圖片點陣，每個點陣都有一個數字。給出起點一個坐標，要求從這個起點坐標開始，把所有與這個起點設置的點都染色成新的顏色。
這題是標準的洪水填充算法。可以用 DFS 也可以用 BFS 。

## 來源
* https://leetcode.com/problems/flood-fill/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0733.Flood-Fill/main.go

```go
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

```

## Benchmark

```sh
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0733.Flood-Fill
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkFloodFillDFS-8         384756555                4.486 ns/op           0 B/op          0 allocs/op
BenchmarkFloodFillBFS-8         309088303                3.642 ns/op           0 B/op          0 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0733.Flood-Fill 3.572s
```