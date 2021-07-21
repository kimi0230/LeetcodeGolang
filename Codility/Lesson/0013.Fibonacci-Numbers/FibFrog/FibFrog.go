package fibfrog

import "fmt"

/**
 * @description: 產生不大於n的斐波那契數的列表
 * @param {int} N
 * @return {*}
 */
func Fib(N int) (fibArr []int) {
	fibArr = append(fibArr, 0)
	fibArr = append(fibArr, 1)
	fibArr = append(fibArr, 1)
	i := 2
	for fibArr[i] < N {
		i = i + 1
		fibArr = append(fibArr, fibArr[i-1]+fibArr[i-2])
	}
	return fibArr
}

func Solution(A []int) int {
	// 終點
	A = append(A, 1)
	N := len(A)
	fibArr := Fib(N)
	// 一次就可以從 -1 跳到 N
	if fibArr[len(fibArr)-1] == N {
		return 1
	}
	fibArr = fibArr[1 : len(fibArr)-1]
	fmt.Println(fibArr)

	// get the leafs that can be reached from the starting shore
	reachable := make([]int, N)
	for _, v := range fibArr {
		if A[v-1] == 1 {
			reachable[v-1] = 1
		}
	}
	// 一開始只能到 index: 4
	// fmt.Println("re", reachable) // [0 0 0 0 1 0 0 0 0 0 0 0]

	// iterate all the positions until you reach the other shore
	for i := 0; i < N; i++ {
		// 忽略不是葉子或已經找過的path
		if A[i] != 1 || reachable[i] > 0 {
			continue
		}

		// get the optimal jump count to reach this leaf
		if A[i] == 1 {
			// 有樹葉
			// 遍歷斐波那契數列, 尋找最少的跳躍次數
			minJump := i + 1
			canJump := false
			for _, f := range fibArr {
				previousIdx := i - f

				if previousIdx < 0 || reachable[previousIdx] == 0 {
					fmt.Printf("[No] %d :previousIdx = %d reachable = %v \n", i, previousIdx, reachable)
					continue
				}
				fmt.Printf("%d :previousIdx = %d reachable = %v \n", i, previousIdx, reachable)

				if minJump > reachable[previousIdx] {
					// 此 previousIdx 位置可以到達
					minJump = reachable[previousIdx]
					canJump = true
				}
			}
			if canJump {
				reachable[i] = minJump + 1
			}
		}
		fmt.Printf("i=%d , reachable = %v \n", i, reachable)
	}

	if reachable[len(reachable)-1] < N {
		return reachable[len(reachable)-1]
	} else {
		return -1
	}
}
