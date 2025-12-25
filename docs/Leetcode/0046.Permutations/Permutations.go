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

/*
generatePermutation: (輸入數組, 數組長度, 遞迴到第幾層depth, path, 使用過的, 結果)
找最短路徑用**BFS**, 其他時用**DFS**用得多一些, 因為遞迴較好寫
假設有棵滿的二叉樹,節點數為 N. 對DFS來說空間複雜度就是遞迴, 最壞的情況就是樹的高度 O(log N)
BFS算法, Queue每次都會存二叉樹一層的節點, 最壞的情況下空間複雜度應該就是樹的最下層的數量, 也就是 N/2. 空間複雜度 O(N)
DFS（深度優先搜索）通常使用堆棧（Stack）來實現。在DFS中，您首先處理一個節點，然後將其子節點按某種順序推入堆棧中，接著繼續處理堆棧頂部的節點，直到堆棧為空。
*/
func generatePermutation(nums []int, numsLen int, depth int, path []int, used *[]bool, res *[][]int) {
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
			generatePermutation(nums, numsLen, depth+1, path, used, res)
			path = path[:len(path)-1]
			// 回朔
			(*used)[i] = false
		}
	}
}
