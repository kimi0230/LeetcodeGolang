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
