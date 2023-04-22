---
title: "Search Graph in Golang"
subtitle: "Search Graph in Golang"
date: 2023-04-22T18:01:25+08:00
lastmod: 2023-04-22T18:01:25+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: ""
license: ""
images: []

tags: [Golang, Algorithms, Search Graph, BFS]
categories: [Golang]

featuredImage: ""
featuredImagePreview: ""

hiddenFromHomePage: false
hiddenFromSearch: false
twemoji: false
lightgallery: true
ruby: true
fraction: true
fontawesome: true
linkToMarkdown: true
rssFullText: false

toc:
  enable: true
  auto: true
code:
  copy: true
  maxShownLines: 50
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

# Search Graph


     2
   /   \
  /     \
 1 - 3 - 5 - 7
  \        /
   \      6
    \   /
      4

一個Graph,隨便指定一個起點, 優先print出周圍的值

// 1 -> [2, 3, 4]
// 2 -> [1, 5]
// 3 -> [1, 5]
// 4 -> [1, 6]
// 5 -> [2, 3, 7]
// 6 -> [4, 7]
// 7 -> [5, 6]
// print : 1,2,3,4,5,6,7
// Also this is valid : 1,4,3,2,6,5,7

例如 開始位置1, print : 1,2,3,4,5,6,


## 解法
用BFS(Queue)
時間複雜度是O(V+E)，其中V是圖中節點的數量，E是圖中邊的數量

# Reference
* https://www.youtube.com/watch?v=BkszA-MvjXA