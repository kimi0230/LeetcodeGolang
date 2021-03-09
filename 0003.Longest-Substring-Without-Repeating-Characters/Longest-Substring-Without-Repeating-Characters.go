/*
# [3. Longest Substring Without Repeating Characters](https://leetcode.com/problems/longest-substring-without-repeating-characters/)

## 题目
Given a string, find the length of the longest substring without repeating characters.

Example 1:

```c
Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
```

Example 2:

```c
Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
```

Example 3:

```c
Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
```

## 題目大意

在一個字符串重尋找沒有重複字母的最長子串。

## 解題思路
這一題和第 438 題，第 3 題，第 76 題，第 567 題類似，用的思想都是"滑動窗口"。

滑動窗口的右邊界不斷的右移，只要沒有重複的字符，就持續向右擴大窗口邊界。一旦出現了重複字符，就需要縮小左邊界，直到重複的字符移出了左邊界，然後繼續移動滑動窗口的右邊界。以此類推，每次移動需要計算當前長度，並判斷是否需要更新最大長度，最終最大的值就是題目中的所求。

O(n)
*/
package longestSubstringwithoutrepeatingcharacters

// LengthOfLongestSubstring 暴力解
func LengthOfLongestSubstring(s string) int {
	slength := len(s)
	if slength == 0 || slength == 1 {
		return slength
	}

	tmpLen := 1
	var maxLen = 1

	for i := 1; i < slength; i++ {
		// 往前找前幾個視窗
		j := i - tmpLen

		for ; j < i; j++ {
			if s[j] == s[i] { // 如果相同，那麼和S[J]到S[I-1]中間的肯定不相同，所以可以直接計算得到
				tmpLen = i - j
				break
			}
		}

		if j == i { // 都不相同
			tmpLen++
		}

		if tmpLen > maxLen {
			maxLen = tmpLen
		}
	}

	return maxLen
}

// LengthOfLongestSubstringPin 用map 紀錄是否重複 效能較好
func LengthOfLongestSubstringPin(s string) int {
	slength := len(s)
	if slength == 0 || slength == 1 {
		return slength
	}

	charMap := make(map[byte]bool)
	maxLen, left, right := 0, 0, 0
	for i := 0; i < slength; i++ {
		if ok := charMap[s[i]]; ok {
			// 有找到
			charMap[s[left]] = false
			left++
		} else {
			charMap[s[i]] = true
			right++
		}

		if maxLen < right-left {
			maxLen = right - left
		}
		if left+maxLen >= slength || right >= len(s) {
			break
		}
	}
	return maxLen
}
