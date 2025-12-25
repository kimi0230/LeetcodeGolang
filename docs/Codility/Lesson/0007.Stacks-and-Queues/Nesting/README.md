# [Nesting](https://app.codility.com/programmers/lessons/7-stacks_and_queues/nesting/)
Determine whether a given string of parentheses (single type) is properly nested.

A string S consisting of N characters is called properly nested if:

S is empty;
S has the form "(U)" where U is a properly nested string;
S has the form "VW" where V and W are properly nested strings.
For example, string "(()(())())" is properly nested but string "())" isn't.

Write a function:

func Solution(S string) int

that, given a string S consisting of N characters, returns 1 if string S is properly nested and 0 otherwise.

For example, given S = "(()(())())", the function should return 1 and given S = "())", the function should return 0, as explained above.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [0..1,000,000];
string S consists only of the characters "(" and/or ")".
Copyright 2009–2021 by Codility Limited. All Rights Reserved. Unauthorized copying, publication or disclosure prohibited.

## 題目大意
括號配對

## 解題思路
與Bracket類似

## 來源
https://app.codility.com/programmers/lessons/7-stacks_and_queues/nesting/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Codility/Lesson/0007.Stacks-and-Queues/Nesting/Nesting.go


```go
package Nesting

import "LeetcodeGolang/Utility/structures"

func Solution(S string) int {
	if len(S) == 0 {
		return 1
	}
	if len(S)%2 != 0 {
		return 0
	}

	stack := structures.NewArrayStack()
	for _, v := range S {
		val := string(v)
		if val == "(" {
			stack.Push(val)
		} else if val == ")" {
			stack.Pop()
		}
	}
	if stack.IsEmpty() {
		return 1
	} else {
		return 0
	}
}
```