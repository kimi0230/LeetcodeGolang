---
title: 138. Copy List with Random Pointer
subtitle: "https://leetcode.com/problems/copy-list-with-random-pointer/description/"
date: 2024-03-12T22:30:00+08:00
lastmod: 2024-03-12T22:30:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0138.Copy-List-with-Random-Pointer"
license: ""
images: []

tags: [LeetCode, Go, Medium, Copy List with Random Pointer, Facebook ,Amazon ,Microsoft ,Bloomberg ,Google]
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
# [138. Copy List with Random Pointer](https://leetcode.com/problems/copy-list-with-random-pointer/description/)

## 題目

## 題目大意


## 解題思路
對於數據結構複製，甭管他怎麼變，你就記住最簡單的方式：一個哈希表 + 兩次遍歷
* 第一次遍歷專門clone節點，藉助 Hash表把原始節點和clone節點的映射存儲起來;
* 第二次專門組裝節點，照著原數據結構的樣子，把clone節點的指標組裝起來。
  
## Big O

* 時間複雜 : `O(n)` ，其中 n 是鏈表的長度。 對於每個節點，我們至多訪問其「後繼節點」和「隨機指標指向的節點」各一次，均攤每個點至多被訪問兩次。
* 空間複雜 : `O(n)` 其中 n 是鏈表的長度。 為哈希表的空間開銷

## 來源
* https://leetcode.com/problems/copy-list-with-random-pointer/description/
* https://leetcode.cn/problems/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0138.Copy-List-with-Random-Pointer/main.go

```go
package copylistwithrandompointer

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

var cacheMap map[*Node]*Node

// 時間複雜 O(n), 空間複雜 O(n)
func copyRandomList(head *Node) *Node {
	cacheMap = make(map[*Node]*Node)
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, ok := cacheMap[node]; ok {
		return n
	}
	newNode := &Node{Val: node.Val}
	cacheMap[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList2(head *Node) *Node {
	cMap := make(map[*Node]*Node)
	cur := head
	// 第一次遍歷專門clone節點，藉助 Hash表把原始節點和clone節點的映射存儲起來;
	for cur != nil {
		newNode := &Node{Val: cur.Val}
		cMap[cur] = newNode
		cur = cur.Next
	}
	// 第二次專門組裝節點，照著原數據結構的樣子，把clone節點的指標組裝起來。
	newHead := cMap[head]
	for head != nil {
		node := cMap[head]
		node.Next = cMap[head.Next]
		node.Random = cMap[head.Random]
		head = head.Next
	}
	return newHead
}

```

##  Benchmark

```sh

```