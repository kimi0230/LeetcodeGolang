package countsemiprimes

import "fmt"

func Solution(N int, P []int, Q []int) []int {
	semiPrime := []int{}

	// 半質數:兩個質數的乘積所得的自然數我們稱之為半質數.
	// 4, 6, 9, 10, 14, 15, 21,22,25,26,33,34,35,38,39,46,49,51,55,57,58,62,65,69,74,77,82,85,86,87,91,93,94,95,106, ...
	// 它們包含1及自己在內合共有3或4個因數
	for i := 1; i <= N; i++ {
		factorCount := 0
		sign := 0
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				factorCount++
				f := i / j
				if f != j {
					if f == j*j {
						// 3個相同: ex i=27, j=3, f=9
						sign = 1
						semiPrime = append(semiPrime, 0)
						break
					} else {
						factorCount++
					}
				}
			}
			if factorCount > 4 {
				sign = 1
				semiPrime = append(semiPrime, 0)
			}
		}
		if sign != 1 {
			if factorCount >= 3 {
				semiPrime = append(semiPrime, i)
			} else {
				semiPrime = append(semiPrime, 0)
			}
		}
	}
	fmt.Println("semiPrime", semiPrime)

	// 當前array和前面一共有幾個半質數
	indexMap := make(map[int]int)
	// 如果是半質數添加到 map
	semiMap := make(map[int]int)
	count := 0
	for i := 0; i < len(semiPrime); i++ {
		if semiPrime[i] != 0 {
			count++
			indexMap[semiPrime[i]] = count
			semiMap[semiPrime[i]] = 0
		} else {
			indexMap[i+1] = count
		}
	}
	fmt.Println("indexMap : ", indexMap)
	fmt.Println("semiMap : ", semiMap)

	result := []int{}
	for i := 0; i < len(P); i++ {
		if _, ok := semiMap[P[i]]; ok {
			result = append(result, indexMap[Q[i]]-indexMap[P[i]]+1)
		} else {
			result = append(result, indexMap[Q[i]]-indexMap[P[i]])
		}
	}
	return result
}
