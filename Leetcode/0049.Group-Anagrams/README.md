---
title: 0049.Group Anagrams
subtitle: "https://leetcode.com/problems/group-anagrams/"
date: 2023-10-11T17:42:00+08:00
lastmod: 2023-10-11T17:42:00+08:00
draft: false
author: "Kimi.Tsai"
authorLink: "https://kimi0230.github.io/"
description: "Group Anagrams"
license: ""
images: []

tags: [LeetCode, Go, Medium, Group Anagrams]
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
# [0049.Grop Anagrams](https://leetcode.com/problems/group-anagrams/)

## 題目
Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

Example 2:

Input: strs = [""]
Output: [[""]]

Example 3:

Input: strs = ["a"]
Output: [["a"]]
 

Constraints:

* 1 <= strs.length <= 104
* 0 <= strs[i].length <= 100
* strs[i] consists of lowercase English letters.

## 題目大意

分組出只用同樣字符產生的的單字

## 解題思路
方法一: 計數
由於互為字母異位詞的兩個字串包含的字母相同，因此兩個字串中的相同字母出現的次數一定是相同的，故可以將每個字母出現的次數使用字串表示，作為哈希表的鍵。

方法二: 排序
由於互為字母異位詞的兩個字串包含的字母相同，因此對兩個字串分別進行排序之後得到的字串一定是相同的，故可以將排序之後的字串作為哈希表的鍵。

## 來源
* https://leetcode.com/problems/group-anagrams/
* https://leetcode.com/problems/group-anagrams/solutions/3645797/hash-solution-with-explanation-smart-solution-o-n-m/
* https://leetcode.cn/problems/group-anagrams/description/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0049.Group-Anagams/main.go

```go
package groupanagrams

import "sort"

// 使用計數: 時間複雜 O(nk), 空間複雜 O(nk)
// n 是單詞的個數 , k是單字的最大長度
// O(n)遍歷單詞, 遍歷字符O(nk)
func GroupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string, len(strs))

	for _, str := range strs {
		// 建立每一個string的特徵`count`, 藉由統計每個字母出現的次數
		count := [26]int{}
		for _, b := range str {
			count[b-'a']++
		}
		// 將同樣的次數的string放置在 map中
		m[count] = append(m[count], str)
	}

	// 把map中的string放入到結果的 `[][]string` array 中
	ans := [][]string{}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}

// 排序: 時間複雜 O(nklog(k)), 空間複雜 O(klog(k))
// n 是單詞的個數 , k是單字的最大長度
// 遍歷單詞 O(n)
// 排序字符 O(klog(k))
func GroupAnagramsBySort(strs []string) [][]string {
	m := make(map[string][]string, len(strs))

	for _, str := range strs {
		// 建立每一個string的特徵, 藉由統計排序
		s := []byte(str)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		sortedStr := string(s)

		// 將同樣的特徵的string放置在 map中
		m[sortedStr] = append(m[sortedStr], str)
	}

	// 把map中的string放入到結果的 `[][]string` array 中
	ans := [][]string{}
	for _, v := range m {
		ans = append(ans, v)
	}
	return ans
}
```

##  Benchmark

```sh
go test -benchmem -run=none LeetcodeGolang/Leetcode/0049.Group-Anagrams -bench=.
goos: darwin
goarch: amd64
pkg: LeetcodeGolang/Leetcode/0049.Group-Anagrams
cpu: Intel(R) Core(TM) i5-8259U CPU @ 2.30GHz
BenchmarkGroupAnagrams-8                  899431              1615 ns/op             968 B/op         12 allocs/op
BenchmarkGroupAnagramsBySort-8            592148              3566 ns/op             776 B/op         33 allocs/op
PASS
ok      LeetcodeGolang/Leetcode/0049.Group-Anagrams     3.611s

```

```go
package groupanagrams

import (
	"reflect"
	"sort"
	"testing"
)

var tests = []struct {
	arg1 []string
	want [][]string
}{
	{
		[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
		[][]string{
			{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
	},
	{
		[]string{},
		[][]string{},
	}, {
		[]string{"a"},
		[][]string{
			{"a"}},
	},
}

// 創建一個輔助函數，將字符串數組排序，以便比較
func sortSubStrings(groups [][]string) {
	for _, group := range groups {
		sort.Strings(group)
	}
	// 排序整個切片
	sort.Slice(groups, func(i, j int) bool {
		return groups[i][0] < groups[j][0]
	})
}

func TestGroupAnagrams(t *testing.T) {
	for _, tt := range tests {
		result := GroupAnagrams(tt.arg1)
		// 對結果進行排序，以便比較
		sortSubStrings(result)

		// 對期望結果進行排序，以便比較
		sortSubStrings(tt.want)

		// 使用反射檢查結果和期望是否相同
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("got = %v, want = %v", result, tt.want)
		}
	}
}

func TestGroupAnagramsBySort(t *testing.T) {
	for _, tt := range tests {
		result := GroupAnagramsBySort(tt.arg1)
		// 對結果進行排序，以便比較
		sortSubStrings(result)

		// 對期望結果進行排序，以便比較
		sortSubStrings(tt.want)

		// 使用反射檢查結果和期望是否相同
		if !reflect.DeepEqual(result, tt.want) {
			t.Errorf("got = %v, want = %v", result, tt.want)
		}
	}
}

func BenchmarkGroupAnagrams(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GroupAnagrams(tests[0].arg1)
	}
}

func BenchmarkGroupAnagramsBySort(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		GroupAnagramsBySort(tests[0].arg1)
	}
}
```