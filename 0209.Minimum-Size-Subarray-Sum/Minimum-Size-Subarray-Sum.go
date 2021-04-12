package minimumsizesubarraysum

import (
	"math"
	"sort"
)

// MinSubArrayLenBurst : 暴力解 時間複雜 O(n^2), 空間複雜 O(1)
func MinSubArrayLenBurst(target int, nums []int) int {
	lens := len(nums)
	if lens <= 0 {
		return 0
	}

	minSize := math.MaxInt32
	for i := 0; i < lens; i++ {
		sum := 0
		for j := i; j < lens; j++ {
			sum += nums[j]
			if sum >= target {
				minSize = min(minSize, j-i+1)
			}
		}
	}
	if minSize == math.MaxInt32 {
		minSize = 0
	}
	return minSize
}

// MinSubArrayLenSlidingWindow : 滑動視窗 時間複雜 O(n), 空間複雜 O(1)
func MinSubArrayLenSlidingWindow(target int, nums []int) int {
	lens := len(nums)
	if lens <= 0 {
		return 0
	}
	minSize := math.MaxInt32

	start, end, sum := 0, 0, 0
	for end < lens {
		sum += nums[end]
		for sum >= target {
			minSize = min(minSize, end-start+1)
			sum -= nums[start]
			start++
		}
		end++

	}

	if minSize == math.MaxInt32 {
		minSize = 0
	}
	return minSize
}

/*
MinSubArrayLenBinarysearch : 前缀和 + 二分查找
為了使用二分查找，需要額外創建一個數組 sums 用於存儲數組nums 的前綴和，其中 sums[i] 表示從 nums[0] 到 nums[i−1] 的元素和。
得到前綴和之後，對於每個開始下標i，可通過二分查找得到大於或等於 i 的最小下標 bound，
使得 sums[bound]-sums[i-1] >= s，
並更新子數組的最小長度（此時子數組的長度是 bound -(i-1) )。
C++ 的 lower_bound，Java 的 Arrays.binarySearch

因為這道題保證了數組中每個元素都為正，所以前綴和一定是遞增的，這一點保證了二分的正確性。如果題目沒有說明數組中每個元素都為正，這裡就不能使用二分來查找這個位置了。
*/
func MinSubArrayLenBinarysearch(target int, nums []int) int {
	lens := len(nums)
	if lens <= 0 {
		return 0
	}
	minSize := math.MaxInt32

	sums := make([]int, lens+1)
	// 為了方便計算 size = lens + 1
	// sums[0] = 0, 前0個的元素和
	// sums[1] = nums[0] , 前1個元素和為 nums[0]
	for i := 1; i <= lens; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	// fmt.Println("sums", sums)
	for i := 1; i <= lens; i++ {
		/*
			target      7
			input       2  3  1  2  4  3
			-----------------------------------
			i           0  1  2  3  4  5  6
			-----------------------------------
			sums        0  2  5  6  8 12 15	// input 的前 n 個元素和
			tmpTarge       7  9 12 13 15 19	// tmpTarge = target + sums[i-1]
			bound          4  5  5  6  6  7	// bound = sort.SearchInts(sums, tmpTarge)
			bound-(i-1)    4  4  3  3  2  2
			minSize        4  4  3  3  2  2
		*/
		tmpTarge := target + sums[i-1]
		bound := sort.SearchInts(sums, tmpTarge)
		// if bound < 0 {
		// 	bound = -bound - 1
		// }
		if bound <= lens {
			minSize = min(minSize, bound-(i-1))
		}
		// fmt.Println("i:", i, " tmpTarge", tmpTarge, "  bound:", bound, "  bound-(i-1)", bound-(i-1), " minSize:", minSize)
	}

	if minSize == math.MaxInt32 {
		minSize = 0
	}
	return minSize
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
