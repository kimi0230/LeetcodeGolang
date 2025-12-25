package subsets

// 時間複雜 O(n^2)	, 空間複雜 O(n)
func Subsets(nums []int) [][]int {
	path, result := []int{}, [][]int{}
	for i := 0; i <= len(nums); i++ {
		genSubsets(nums, i, 0, path, &result)
	}
	return result
}

func genSubsets(nums []int, elemSize, start int, path []int, result *[][]int) {
	if len(path) == elemSize {
		b := make([]int, len(path))
		copy(b, path)
		*result = append(*result, b)
		return
	}

	for k := start; k < len(nums); k++ {
		path = append(path, nums[k])
		genSubsets(nums, elemSize, k+1, path, result)
		path = path[:len(path)-1]
	}
}

// 時間複雜 O(n^2)	, 空間複雜 O(n)
func SubsetsDFS(nums []int) [][]int {
	path, result := []int{}, [][]int{}
	genSubsetsDFS(nums, path, &result)
	return result
}

func genSubsetsDFS(nums []int, path []int, result *[][]int) {
	b := make([]int, len(path))
	copy(b, path)
	*result = append(*result, b)
	for i := 0; i < len(nums); i++ {
		// n個 element 在slice中
		genSubsetsDFS(nums[i+1:], append(path, nums[i]), result)
	}
}
