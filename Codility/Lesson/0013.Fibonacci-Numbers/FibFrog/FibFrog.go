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
	route := make([]int, N)
	for i, _ := range route {
		route[i] = N
	}
	fmt.Println(route)
	for i := 0; i < N; i++ {
		if A[i] == 0 {
			// 有樹葉
			// 遍歷斐波那契數列, 尋找最少的跳躍次數
			for _, f := range fibArr {

			}
		}
	}

	fmt.Println(route)
	if route[len(route)-1] < N {
		return route[len(route)-1]
	} else {
		return -1
	}
}
