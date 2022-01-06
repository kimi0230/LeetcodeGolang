# [75. Sort Colors](https://leetcode.com/problems/sort-colors/)

## 题目

Given an array with n objects colored red, white or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white and blue.

Here, we will use the integers 0, 1, and 2 to represent the color red, white, and blue respectively.

Note: You are not suppose to use the library's sort function for this problem.

Example 1:  

```
Input: [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

Follow up:

- A rather straight forward solution is a two-pass algorithm using counting sort.  
First, iterate the array counting number of 0's, 1's, and 2's, then overwrite array with total number of 0's, then 1's and followed by 2's.
- Could you come up with a one-pass algorithm using only constant space?

---
給定一個包含紅色、白色和藍色，一共 n 個元素的數組，原地對它們進行排序，使得相同顏色的元素相鄰，並按照紅色、白色、藍色順序排列。

此題中，我們使用整數 0、 1 和 2 分別表示紅色、白色和藍色。


範例 1：
```
輸入：nums = [2,0,2,1,1,0]
輸出：[0,0,1,1,2,2]
```

範例 2：
```
輸入：nums = [2,0,1]
輸出：[0,1,2]
```
範例 3：
```
輸入：nums = [0]
輸出：[0]
```
範例 4：
```
輸入：nums = [1]
輸出：[1]
```

提示：
* n == nums.length
* 1 <= n <= 300
* nums[i] 為 0、1 或 2
 

進階：
* 你可以不使用代碼庫中的排序函數來解決這道題嗎？
* 你能想出一個僅使用常數空間的一趟掃描算法嗎？


來源：力扣（LeetCode）
鏈接：https://leetcode-cn.com/problems/sort-colors
著作權歸領扣網絡所有。商業轉載請聯繫官方授權，非商業轉載請註明出處。

## 題目大意

抽象題意其實就是排序。這題可以用快排一次通過。

## 解題思路

題目末尾的 Follow up 提出了一個更高的要求，能否用一次循環解決問題？這題由於數字只會出現 0，1，2 這三個數字，所以用游標移動來控制順序也是可以的。具體做法：0 是排在最前面的，所以只要添加一個 0，就需要放置 1 和 2。1 排在 2 前面，所以添加 1 的時候也需要放置 2 。至於最後的 2，只用移動游標即可。

這道題可以用計數排序，適合待排序數字很少的題目。用一個 3 個容量的數組分別計數，記錄 0，1，2 出現的個數。然後再根據個數排列 0，1，2 即可。時間複雜度 O(n)，空間複雜度 O(K)。這一題 K = 3。

這道題也可以用一次三路快排。數組分為 3 部分，第一個部分都是 0，中間部分都是 1，最後部分都是 2 。

## 來源
* https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0075.Sort-Colors/
* https://leetcode-cn.com/problems/sort-colors/

## 解答
https://github.com/kimi0230/LeetcodeGolang/blob/master/Leetcode/0075.Sort-Colors/0075.Sort-Colors.go

```go
package sortcolors

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	var index0, index1, index2 = 0, 0, 0
	for _, num := range nums {
		switch num {
		case 0:
			nums[index2] = 2
			index2++
			nums[index1] = 1
			index1++
			nums[index0] = 0
			index0++
		case 1:
			nums[index2] = 2
			index2++
			nums[index1] = 1
			index1++
		case 2:
			index2++
		}
	}
}

/*
二路快排與三路快排
先掌握二路快排， https://leetcode-cn.com/problems/sort-an-array/solution/golang-er-lu-kuai-pai-by-rqb-2/
三路快排稍微改下即可

區別：
二路快排，分兩部分：小於等於、大於二部分。
三路快排分為，小於、等於、大於三部分。

三路快排：
相比二路快排增加一個變量記錄相等的起始元素。
假設選擇數組第一個元素作為參考元素。則起始下邊為0，即為equalHead

作者：hibrq
链接：https://leetcode-cn.com/problems/sort-an-array/solution/golang-san-lu-kuai-pai-by-rqb-2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func sortColorsQuickSort(nums []int) {
	if len(nums) < 2 {
		return
	}

	pivot := nums[0]

	head, tail := 0, len(nums)-1
	i := 1
	equalHead := 0
	for head < tail {
		if nums[i] < pivot {
			nums[i], nums[equalHead] = nums[equalHead], nums[i]
			head++
			i++
			equalHead++
		} else if nums[i] == pivot {
			i++
			head++
		} else {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail--
		}
	}
	sortColorsQuickSort(nums[:equalHead])
	sortColorsQuickSort(nums[tail+1:])
}
```