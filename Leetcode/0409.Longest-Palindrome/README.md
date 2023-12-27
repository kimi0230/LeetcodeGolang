---
title: 0409. Longest Palindrome
tags: Easy, String
author: Kimi Tsai <kimi0230@gmail.com>
description:
---
# [409. Longest Palindrome](https://leetcode.com/problems/longest-palindrome/)

## 題目
Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.

Letters are case sensitive, for example, "Aa" is not considered a palindrome here.

 

Example 1:

```
Input: s = "abccccdd"
Output: 7
Explanation: One longest palindrome that can be built is "dccaccd", whose length is 7.
```

Example 2:
```
Input: s = "a"
Output: 1
Explanation: The longest palindrome that can be built is "a", whose length is 1.
```

Constraints:

* 1 <= s.length <= 2000
* s consists of lowercase and/or uppercase English letters only.

## 題目大意
給定一個包含大寫字母和小寫字母的字符串，找到通過這些字母構造成的最長的回文串。
在構造過程中，請注意區分大小寫。比如 `Aa` 不能當做一個回文字符串。
注意:假設字符串的長度不會超過 1010。

## 解題思路
* 給出一個字符串，要求用這個字符串裡面的字符組成一個回文串，問回文串最多可以組合成多長的？
* 這也是一道題水題，然後先統計每個字符的頻次，每個字符能取2個的取2個，不足2個的並且當前構造中的回文串是偶數的情況下（即每2個都陣容了），可以取1個。最後組合出來的就是終止回文串。

## 來源
* https://leetcode.com/problems/longest-palindrome/

## 解答

```go
package longestpalindrome

func LongestPalindrome(s string) int {
	counter := make(map[rune]int)
	for _, r := range s {
		counter[r]++
	}

	result := 0

	for _, v := range counter {
		result += v / 2 * 2

		// 字符出現奇數次，我們可以選擇其中一個, 放在回文串的中間，這可以貢獻一個長度
		if result%2 == 0 && v%2 == 1 {
			result++
		}
	}

	return result
}
```
