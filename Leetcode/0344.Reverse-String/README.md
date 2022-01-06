# [344. Reverse String](https://leetcode.com/problems/reverse-string/)

## 題目

Write a function that reverses a string. The input string is given as an array of characters char[].

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

You may assume all the characters consist of printable ascii characters.

Example 1:

```c
Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
```

Example 2:

```c
Input: ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
```

## 題目大意

題目要求我們反轉一個字符串。

## 解題思路

這一題的解題思路是用 2 個指針，指針對撞的思路，來不斷交換首尾元素，即可。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0300~0399/0344.Reverse-String/
* https://leetcode-cn.com/problems/reverse-string/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0344.Reverse-Stringm/Reverse-String.go

```go
package reversestring

func ReverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

```