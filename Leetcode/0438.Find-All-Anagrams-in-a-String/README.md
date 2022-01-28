---
title: 0438.Find All Anagrams in a String
tags: Medium, Sliding Window
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [438. Find All Anagrams in a String](https://leetcode.com/problems/find-all-anagrams-in-a-string/)
## 题目
Given two strings s and p, return an array of all the start indices of p's anagrams in s. You may return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
```
Input: s = "cbaebabacd", p = "abc"
Output: [0,6]
Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
```

Example 2:
```
Input: s = "abab", p = "ab"
Output: [0,1,2]
Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
```

Constraints:

1 <= s.length, p.length <= 3 * 104
s and p consist of lowercase English letters.

## 題目大意
找所有字母異位詞, 就像[全排列](https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0567.Permutation-in-String/main.go)
給定一個字符串 S 和非空的字符串 P, 找到 S 中所有是 P 得排列, 並返回他的起始 index
## 解題思路
跟 [0567.Permutation-in-String](https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0567.Permutation-in-String/main.go)類似, 只是把找到的答案記錄起來

## 來源
* https://leetcode.com/problems/find-all-anagrams-in-a-string/
* 
## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0438.Find-All-Anagrams-in-a-String/main.go

```go
package findallanagramsinastring

func FindAnagrams(s string, p string) []int {
	need, window := make(map[rune]int), make(map[rune]int)
	for _, c := range p {
		need[c]++
	}

	left, right := 0, 0
	valid := 0
	res := []int{} //紀錄結果
	for right < len(s) {
		c := rune(s[right])
		right++

		if need[c] > 0 {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		// fmt.Printf("[%d,%d) \n", left, right)

		// 判斷左視窗是否收縮, 看看視窗長度是否同要找的字串的長度
		if (right - left) >= len(p) {
			if valid == len(need) {
				// 想要的字元都找到了, 紀錄index
				res = append(res, left)
			}
			d := rune(s[left])
			left++
			if need[d] > 0 {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

// 用 slice 取代 map 來優化
func FindAnagramsSlice(s string, p string) []int {
	need := [256]int{}
	for _, c := range p {
		need[c-'a']++
	}
	left, right := 0, 0
	count := len(p)
	res := []int{}

	for right < len(s) {
		c := s[right] - 'a'
		if need[c] > 0 {
			count--
		}
		need[c]--
		right++

		if count == 0 {
			res = append(res, left)
		}

		if (right - left) >= len(p) {
			d := s[left] - 'a'
			if need[d] >= 0 {
				count++
			}
			need[d]++
			left++
		}
	}
	return res
}

```