# [46. Permutations](https://leetcode.com/problems/permutations/)

## 題目
Given a collection of **distinct** integers, return all possible permutations(排列).
**Example:**
    Input: [1,2,3]
    Output:
    [
      [1,2,3],
      [1,3,2],
      [2,1,3],
      [2,3,1],
      [3,1,2],
      [3,2,1]
    ]

## 題目大意
給定一個沒有重複數字的序列，返回其所有可能的全排列。

## 解題思路
解決回朔問題可用一個`決策樹`的遍歷過程
1. 路徑: 也就是已經做的選擇
2. 選擇列表: 也就是當前可以做的選擇
3. 結束條件: 也就是達到決策樹底層, 無法再做選擇的條件

```python
result = []
def backtrack(路徑, 選擇列表):
	if 滿足結束條件:
		result.add(路徑)
		return
	
	for 選擇 in 選擇列表:
		做選擇
		backtrack(路徑, 選擇列表)
		撤銷選擇
```
```
                       選擇:[1,2,3]
                            []
          [1]/              |[2]            \[3]
        [2]/  \[3]      [1]/  \[3]       [1]/  \[2]
        |[3]   |[2]     |[3]   |[1]      |[2]    |[1]   
結果  [1,2,3]  [1,3,2] [2,1,3] [2,3,1]  [3,1,2]  [3,2,1]
```
![](/asset/images/0046_permutations.png)

- 求出一個數組的排列組合中的所有排列，用 DFS 深搜即可。
這個問題可以看作有 ñ 個排列成一行的空格，我們需要從左往右依此填入題目給定的 ñ個數，每個數只能使用一次。
那麼很直接的可以想到一種窮舉的算法，即從左往右每一個位置都依此嘗試填入一個數，
看能不能填完這ñ 個空格，在程序中我們可以用「回溯法」來模擬這個過程
回溯法：
一種通過探索所有可能的候選解來找出所有的解的算法。如果候選解被確認不是一個解（或者至少不是最後一個解），
回溯算法會通過在上一步進行一些變化拋棄該解，即回溯並且再次嘗試。

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/permutations/solution/quan-pai-lie-by-leetcode-solution-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0046.Permutations/
* https://leetcode-cn.com/problems/permutations/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0046.Permutations/Permutations.go


時間複雜 O(n)
```go
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
```