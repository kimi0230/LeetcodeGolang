---
title: 0695.Max Area of Island
subtitle: "https://leetcode.com/problems/max-area-of-island/"
date: 2023-07-06T17:58:00+08:00
lastmod: 2023-07-06T17:58:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Max Area of Island"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Medium
  - Max Area of Island
  - DFS
  - BFS
  - Array & String
  - Matrix
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
# [695. Max Area of Island](https://leetcode.com/problems/max-area-of-island/)

## 題目
You are given an m x n binary matrix grid. An island is a group of 1's (representing land) connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid are surrounded by water.

The area of an island is the number of cells with a value 1 in the island.

Return the maximum area of an island in grid. If there is no island, return 0.


Example 1:

![](images/maxarea1-grid.jpg)

```
Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]]
Output: 6
Explanation: The answer is not 11, because the island must be connected 4-directionally.
```

Example 2:
```
Input: grid = [[0,0,0,0,0,0,0,0]]
Output: 0
```

Constraints:

* m == grid.length
* n == grid[i].length
* 1 <= m, n <= 50
* grid[i][j] is either 0 or 1.

## 題目大意
給定一個由 '1'（陸地）和 '0'（水域）組成的二維矩陣，計算矩陣中連續的1所組成的最大區域的面積。

## 解題思路
使用BFS或DFS遍歷矩陣，計算每個島嶼的面積，並找到最大的面積值。
給出一個地圖，要求計算上面島嶼的。注意島嶼的定義是要求都是海（為0的點），如果土地（為1的點）靠在地圖邊緣，不能是珊瑚島。
這題和第200題，第1254題解題思路是一致的。
DPS深搜。不過這題需要多處理2件事情，一是注意靠邊緣的島嶼不能計算關係，二是動態維護島嶼的最大面積。

使用DFS的解法：

1. 創建一個變量maxArea，用於記錄最大面積，初始化為0。
2. 遍歷二維網格，對於每個為1的位置，調用dfs函數計算以該位置為起點的島嶼面積。
3. 在dfs函數中，首先檢查當前位置的合法性，如果超出網格範圍或者該位置已經訪問過或者是水域（0），則返回0。
4. 將當前位置標記為已訪問，並初始化面積為1。
5. 遞迴地對當前位置的上、下、左、右四個相鄰位置進行dfs，並將返回的面積加到當前面積上。
6. 返回當前面積。
7. 更新maxArea為當前面積和maxArea中的較大值。
8. 遍歷完整個網格後，返回maxArea作為結果。

時間複雜度: 遍歷整個網格的時間複雜度為O(m*n)，其中m和n分別為網格的行數和列數。
空間複雜度: 使用了遞迴調用的栈空間，空間複雜度為O(m*n)。



## 來源
* https://leetcode.com/problems/max-area-of-island/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0695.Max-Area-of-Island/main.go

```go
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

```

##  Benchmark

```sh

```
