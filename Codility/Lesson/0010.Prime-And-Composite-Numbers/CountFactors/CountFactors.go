package CountFactors

import (
	"math"
)

func Solution(N int) int {
	result := 0
	for i := 1; i <= int(math.Pow(float64(N), 0.5)); i++ {
		if N%i == 0 {
			if i*i == N {
				// fmt.Println("+1 : ", i)
				result++
			} else {
				// fmt.Println("+2 : ", i)
				result += 2
			}
		}
	}
	return result
}
