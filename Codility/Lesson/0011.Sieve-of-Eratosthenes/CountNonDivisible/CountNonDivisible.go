package countnondivisible

import (
	"math"
)

/*
Task Score 100%
Correctness 100%
Performance 100%
*/
func Solution(A []int) []int {
	// write your code in Go 1.4
	result := []int{}
	if len(A) < 0 {
		return result
	}

	elementDict := make(map[int]int)
	for _, i := range A {
		elementDict[i]++
	}

	// 用 map存起非因子數的個數, 空間換取時間
	nonDivisorsCountMap := make(map[int]int)
	for _, val := range A {
		if v, ok := nonDivisorsCountMap[val]; ok {
			result = append(result, v)
		} else {
			divisors := 0
			for factor := 1; factor*factor <= val; factor++ {
				if val%factor == 0 {
					// 檢查是否有在原先的 array 中, 並取得因子次數
					if v, ok := elementDict[factor]; ok {
						divisors += v
					}

					// 避免因子重複計算
					otherFactor := int(val / factor)
					if v, ok := elementDict[otherFactor]; ok && otherFactor != factor {
						divisors += v
					}
				}
			}
			// 推出非因子次數
			nonDivisors := len(A) - divisors
			result = append(result, nonDivisors)
			nonDivisorsCountMap[val] = nonDivisors
		}
	}
	return result
}

/*
Task Score 88%
Correctness 100%
Performance 75%
*/
func Solution2(A []int) []int {
	result := []int{}
	if len(A) < 0 {
		return result
	}

	elementDict := make(map[int]int)
	for _, i := range A {
		elementDict[i]++
	}

	// 用 map存起非因子數的個數, 空間換取時間
	nonDivisorsCountMap := make(map[int]int)
	for _, val := range A {
		if v, ok := nonDivisorsCountMap[val]; ok {
			result = append(result, v)
		} else {
			divisors := 0
			for factor := 1; factor <= int(math.Pow(float64(val), 0.5)); factor++ {
				if val%factor == 0 {
					// 檢查是否有在原先的 array 中, 並取得因子次數
					if v, ok := elementDict[factor]; ok {
						divisors += v
					}

					// 避免因子重複計算
					otherFactor := int(val / factor)
					if v, ok := elementDict[otherFactor]; ok && otherFactor != factor {
						divisors += v
					}
				}
			}
			// 推出非因子次數
			nonDivisors := len(A) - divisors
			result = append(result, nonDivisors)
			nonDivisorsCountMap[val] = nonDivisors
		}
	}
	return result
}

/*
Task Score 77%
Correctness 100%
Performance 50%
*/
func Solution1(A []int) []int {
	result := []int{}
	if len(A) < 0 {
		return result
	}

	elementDict := make(map[int]int)
	for _, i := range A {
		elementDict[i]++
	}

	for _, val := range A {
		divisors := 0
		for factor := 1; factor <= int(math.Pow(float64(val), 0.5)); factor++ {
			if val%factor == 0 {
				// 檢查是否有在原先的 array 中, 並取得因子次數
				if v, ok := elementDict[factor]; ok {
					divisors += v
				}

				// 避免因子重複計算
				otherFactor := int(val / factor)
				if v, ok := elementDict[otherFactor]; ok && otherFactor != factor {
					divisors += v
				}
			}
		}
		// 推出非因子次數
		nonDivisors := len(A) - divisors
		result = append(result, nonDivisors)
	}
	return result
}
