# [Brackets](https://app.codility.com/programmers/lessons/7-stacks_and_queues/brackets/)
Determine whether a given string of parentheses (multiple types) is properly nested.

A string S consisting of N characters is considered to be properly nested if any of the following conditions is true:

S is empty;
S has the form "(U)" or "[U]" or "{U}" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.
For example, the string "{[()()]}" is properly nested but "([)()]" is not.

Write a function:

func Solution(S string) int

that, given a string S consisting of N characters, returns 1 if S is properly nested and 0 otherwise.

For example, given S = "{[()()]}", the function should return 1 and given S = "([)()]", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..200,000];
string S consists only of the following characters: "(", "{", "[", "]", "}" and/or ")".
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.


## 題目大意
括號配對, 可配對回傳1 反之回傳0.
## 解題思路
將左括號都放入stack. 遇到右括號時將stack pop出來並檢查pop出來的左括號是否跟右括號配對.
## 來源
https://app.codility.com/programmers/lessons/7-stacks_and_queues/brackets/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0007.Stacks-and-Queues/Brackets/Brackets.go


```go
package Brackets

import (
	"LeetcodeGolang/Utility/structures"
)

func Solution(S string) int {
	if len(S) == 0 {
		return 1
	}
	if len(S)%2 != 0 {
		return 0
	}

	BracketMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	stack := structures.NewArrayStack()
	for _, v := range S {
		val := string(v)
		if val == "(" || val == "[" || val == "{" {
			stack.Push(val)
		} else if val == ")" || val == "]" || val == "}" {
			if stack.IsEmpty() {
				return 0
			}

			topVal := stack.Top()
			if topVal == BracketMap[val] {
				stack.Pop()
			} else {
				// 找不到可配對的括號
				return 0
			}
		}
	}
	if stack.IsEmpty() {
		return 1
	} else {
		return 0
	}
}
```