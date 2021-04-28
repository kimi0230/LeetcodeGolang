package Fish

import "LeetcodeGolang/Utility/structures"

func Solution(A []int, B []int) int {
	stack := structures.NewArrayStack()
	aliveFish := 0
	for idx, val := range B {
		if val == 1 {
			stack.Push(A[idx])
		} else {
			// 繼續往下游
			for !stack.IsEmpty() {
				if stack.Top().(int) < A[idx] {
					// stack的魚比遇到的魚還小, stack被吃掉
					stack.Pop()
				} else {
					break
				}
			}
			if stack.IsEmpty() {
				aliveFish++
			}
		}
	}
	return aliveFish + stack.Size()
}
