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
