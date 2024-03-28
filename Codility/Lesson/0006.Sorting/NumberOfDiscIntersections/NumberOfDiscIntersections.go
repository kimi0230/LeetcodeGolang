package NumberOfDiscIntersections

import "sort"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 時間複雜 O(n^2)
func SolutionDirect(A []int) int {
	count := 0
	for indexI, valueI := range A {
		tmpArr := A[indexI+1:]
		for indexJ, valueJ := range tmpArr {
			if valueI+valueJ >= indexJ+indexI+1-indexI {
				count++
			}
		}
	}

	return count
}

// 時間複雜 O(nlogn) or O(n)
// TODO: 待研究
func Solution(A []int) int {
	iLimit := make([]int, len(A)) // 左
	jLimit := make([]int, len(A)) // 右
	result := 0

	for i := 0; i < len(A); i += 1 {
		iLimit[i] = i - A[i]
		jLimit[i] = i + A[i]
	}
	// 針對iLimit中的每個元素，利用二分查找算法，找到其不小於jLimit中元素的個數
	sort.Ints(iLimit)
	sort.Ints(jLimit)
	for idx, _ := range iLimit {
		end := jLimit[idx]

		// Binary search for index of element of the rightmost value less than to the interval-end
		count := sort.Search(len(iLimit), func(i int) bool {
			return iLimit[i] > end
		})

		// 因為i=j時，A[i]+i 肯定不小於j-A[j],也就是說多算了一個，因此要減去1。
		// 減去idx是因為圓盤A和圓盤B相交，次數加上1了，圓盤B和圓盤A相交就不用再加1了。
		count = count - idx - 1
		result += count

		if result > 10000000 {
			return -1
		}

	}

	return result
}
