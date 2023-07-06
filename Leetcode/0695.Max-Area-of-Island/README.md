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

tags: [LeetCode, Go, Medium, Max Area of Island, DFS, BFS, Array & String, Matrix]
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

## 來源
* https://leetcode.com/problems/max-area-of-island/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0695.Max-Area-of-Island/main.go

```go

```

##  Benchmark

```sh

```
