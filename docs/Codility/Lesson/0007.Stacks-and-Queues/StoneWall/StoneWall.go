package StoneWall

import "LeetcodeGolang/Utility/structures"

func Solution(H []int) int {
	stack := structures.NewArrayStack()
	result := 0
	for _, v := range H {
		for !stack.IsEmpty() && stack.Top().(int) > v {
			stack.Pop()
		}
		if stack.IsEmpty() || stack.Top().(int) < v {
			stack.Push(v)
			result++
		}
	}
	return result
}
