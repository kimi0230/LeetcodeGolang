---
title: 381.Insert Delete GetRandom O(1) Duplicates allowed
subtitle: "https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed"
date: 2024-02-07T21:26:00+08:00
lastmod: 2024-02-07T21:26:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0381.Insert-Delete-GetRandom-O1-Duplicates-allowed"
license: ""
images: []

tags:
  - LeetCode
  - Go
  - Hard
  - Insert Delete GetRandom O(1) Duplicates allowed
  - Array
  - Hash
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
# [381.Insert Delete GetRandom O(1) Duplicates allowed](https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/description/)

## 題目


## 題目大意

## 解題思路
RandomizedCollection is a data structure that contains a collection of numbers, possibly duplicates (i.e., a multiset). It should support inserting and removing specific elements and also reporting a random element.

Implement the RandomizedCollection class:

RandomizedCollection() Initializes the empty RandomizedCollection object.
bool insert(int val) Inserts an item val into the multiset, even if the item is already present. Returns true if the item is not present, false otherwise.
bool remove(int val) Removes an item val from the multiset if present. Returns true if the item is present, false otherwise. Note that if val has multiple occurrences in the multiset, we only remove one of them.
int getRandom() Returns a random element from the current multiset of elements. The probability of each element being returned is linearly related to the number of the same values the multiset contains.
You must implement the functions of the class such that each function works on average O(1) time complexity.

Note: The test cases are generated such that getRandom will only be called if there is at least one item in the RandomizedCollection.

 

Example 1:

Input
["RandomizedCollection", "insert", "insert", "insert", "getRandom", "remove", "getRandom"]
[[], [1], [1], [2], [], [1], []]
Output
[null, true, false, true, 2, true, 1]

Explanation
RandomizedCollection randomizedCollection = new RandomizedCollection();
randomizedCollection.insert(1);   // return true since the collection does not contain 1.
                                  // Inserts 1 into the collection.
randomizedCollection.insert(1);   // return false since the collection contains 1.
                                  // Inserts another 1 into the collection. Collection now contains [1,1].
randomizedCollection.insert(2);   // return true since the collection does not contain 2.
                                  // Inserts 2 into the collection. Collection now contains [1,1,2].
randomizedCollection.getRandom(); // getRandom should:
                                  // - return 1 with probability 2/3, or
                                  // - return 2 with probability 1/3.
randomizedCollection.remove(1);   // return true since the collection contains 1.
                                  // Removes 1 from the collection. Collection now contains [1,2].
randomizedCollection.getRandom(); // getRandom should return 1 or 2, both equally likely.
 

Constraints:

-231 <= val <= 231 - 1
At most 2 * 105 calls in total will be made to insert, remove, and getRandom.
There will be at least one element in the data structure when getRandom is called.

## Big O
時間複雜 : `O(1)`
空間複雜 : `O(n)`

## 來源
* https://leetcode.com/problems/insert-delete-getrandom-o1-duplicates-allowed/description/
* https://leetcode.cn/problems/insert-delete-getrandom-o1-duplicates-allowed/description/
* https://www.youtube.com/watch?v=46dZH7LDbf8

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0381.Insert-Delete-GetRandom-O1-Duplicates-allowed/main.go

```go
package insertdeletegetrandomo1duplicatesallowed

import (
	"math/rand"
)

// 時間複雜 O(1), 空間複雜 O(n)
type RandomizedCollection struct {
	arr  []int
	set  map[int]map[int]struct{}
	size int
}

func Constructor() RandomizedCollection {
	arr := make([]int, 0)
	set := make(map[int]map[int]struct{})
	size := 0
	return RandomizedCollection{arr, set, size}
}

func (this *RandomizedCollection) Insert(val int) bool {
	ids, ok := this.set[val]
	if !ok {
		// 沒有此數字了, 建立一個index的map
		ids = map[int]struct{}{}
		this.set[val] = ids
	}
	// index的map, key為當前最後的index, value為空的struct{}{}
	ids[this.size] = struct{}{}
	this.arr = append(this.arr, val)
	this.size++
	return !ok
}

func (this *RandomizedCollection) Remove(val int) bool {
	ids, ok := this.set[val]
	if !ok {
		return false
	}

	// 找出此val的的index, 並且把最後一個index的map刪除, 並且把此index的value設為空的struct{}{}
	var i int
	for id := range ids {
		i = id
		break
	}

	// 將最後一筆移到要替換的id
	this.arr[i] = this.arr[this.size-1]
	delete(ids, i)
	// 因為把最後一個元素移到前面了
	delete(this.set[this.arr[i]], this.size-1)

	if i < this.size-1 {
		// fmt.Printf("i =%d, this.size =%d\t", i, this.size)
		this.set[this.arr[i]][i] = struct{}{}
	}
	if len(ids) == 0 {
		// 沒有此數字了
		delete(this.set, val)
	}

	this.arr = this.arr[:this.size-1]
	this.size--
	return true
}

func (this *RandomizedCollection) GetRandom() int {
	index := rand.Intn(this.size)
	return this.arr[index]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

```

##  Benchmark

```sh

```
