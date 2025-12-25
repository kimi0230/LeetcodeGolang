package mergesortedarray

func Merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	num1Idx, num2Idx, tail := m-1, n-1, m+n-1
	// 從每一個array的後面開始比較, 並放入最後面的位子. 這樣只要跑一次
	for ; num1Idx >= 0 && num2Idx >= 0; tail-- {
		if nums1[num1Idx] > nums2[num2Idx] {
			nums1[tail] = nums1[num1Idx]
			num1Idx--
		} else {
			nums1[tail] = nums2[num2Idx]
			num2Idx--
		}
	}

	// 將尚未跑完的補齊
	for ; num2Idx > 0; tail-- {
		nums1[tail] = nums2[num2Idx]
		num2Idx--
	}
}
