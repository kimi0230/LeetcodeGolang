package frogjump

import (
	"math"
)

func Solution(X int, Y int, D int) int {
	if Y < X {
		return 0
	}
	remainDist := Y - X
	return int(math.Ceil(float64(remainDist) / float64(D)))
}
