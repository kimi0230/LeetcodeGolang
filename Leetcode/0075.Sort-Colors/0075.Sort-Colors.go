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
