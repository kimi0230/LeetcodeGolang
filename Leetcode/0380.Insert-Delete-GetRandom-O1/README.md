---
title: 380. Insert Delete GetRandom O(1)
subtitle: "https://leetcode.com/problems/insert-delete-getrandom-o1/description/"
date: 2024-01-22T21:20:00+08:00
lastmod: 2024-01-22T21:20:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "0380.Insert-Delete-GetRandom-O(1)"
license: ""
images: []

tags: [LeetCode, Go, Medium, Insert Delete GetRandom O(1), Array, Hash]
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
# [380. Insert Delete GetRandom O(1)](https://leetcode.com/problems/insert-delete-getrandom-o1/description/)

## 題目
Implement the RandomizedSet class:

RandomizedSet() Initializes the RandomizedSet object.
bool insert(int val) Inserts an item val into the set if not present. Returns true if the item was not present, false otherwise.
bool remove(int val) Removes an item val from the set if present. Returns true if the item was present, false otherwise.
int getRandom() Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called). Each element must have the same probability of being returned.
You must implement the functions of the class such that each function works in average O(1) time complexity.

Example 1:

Input
["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove", "insert", "getRandom"]
[[], [1], [2], [2], [], [1], [2], []]
Output
[null, true, false, true, 2, true, false, 2]

Explanation
RandomizedSet randomizedSet = new RandomizedSet();
randomizedSet.insert(1); // Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomizedSet.remove(2); // Returns false as 2 does not exist in the set.
randomizedSet.insert(2); // Inserts 2 to the set, returns true. Set now contains [1,2].
randomizedSet.getRandom(); // getRandom() should return either 1 or 2 randomly.
randomizedSet.remove(1); // Removes 1 from the set, returns true. Set now contains [2].
randomizedSet.insert(2); // 2 was already in the set, so return false.
randomizedSet.getRandom(); // Since 2 is the only number in the set, getRandom() will always return 2.
 

Constraints:

-231 <= val <= 231 - 1
At most 2 * 105 calls will be made to insert, remove, and getRandom.
There will be at least one element in the data structure when getRandom is called.

## 題目大意

## 解題思路
用map紀錄每個元素的index

## Big O
時間複雜 : `O(1)`
空間複雜 : `O(n)`

## 來源
* https://leetcode.com/problems/insert-delete-getrandom-o1/description/
* https://leetcode.cn/problems/insert-delete-getrandom-o1/description/
* https://www.youtube.com/watch?v=46dZH7LDbf8

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0380.Insert-Delete-GetRandom-O1/main.go

```go
package insertdeletegetrandom

import "math/rand"

// 時間複雜 O(1), 空間複雜 O(n)
type RandomizedSet struct {
	arr  []int
	set  map[int]int
	size int
}

func Constructor() RandomizedSet {
	arr := make([]int, 0)
	set := make(map[int]int)
	size := 0
	return RandomizedSet{arr, set, size}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.set[val]; ok {
		return false
	}
	this.set[val] = this.size
	this.size++
	this.arr = append(this.arr, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.set[val]
	if !ok {
		return false
	}
	// swapping
	lastValue := this.arr[this.size-1]
	this.arr[index] = lastValue
	this.set[lastValue] = index

	// Remove last value
	this.arr = this.arr[:this.size-1]
	delete(this.set, val)
	this.size--
	return true
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(this.size)
	return this.arr[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

```

##  Benchmark

```sh

```

## Follow up
如果允許重複, 譬如多個1

```go
```