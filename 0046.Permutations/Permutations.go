package permutations

func Permute(nums []int) [][]int {
	numsLen := len(nums)
	if numsLen == 0 {
		return [][]int{}
	}
	used, path, res := make([]bool, numsLen), []int{}, [][]int{}
	dfs(nums, numsLen, 0, path, &used, &res)
	return res
}

// dfs: 輸入數組, 數組長度, 遞迴到第幾層depth, path, 使用過的, 結果
func dfs(nums []int, numsLen int, depth int, path []int, used *[]bool, res *[][]int) {
	if depth == numsLen {
		temp := make([]int, len(path))
		copy(temp, path)
		*res = append(*res, temp)
		return
	}

	for i := 0; i < numsLen; i++ {
		if !(*used)[i] {
			// 沒使用過, 將其紀錄走過
			(*used)[i] = true
			path = append(path, nums[i])
			dfs(nums, numsLen, depth+1, path, used, res)
			path = path[:len(path)-1]
			// 回朔
			(*used)[i] = false
		}
	}
}
