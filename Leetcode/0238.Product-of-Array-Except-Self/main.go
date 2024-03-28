package productofarrayexceptself

// 時間複雜 O(), 空間複雜 O()
func productExceptSelf(nums []int) []int {
	result, left, right := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))

	// left 為左側所有的成績
	// 索引為'0' 的元素, 因為左側沒有元素,所以left[0]=1
	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		result[i] = left[i] * right[i]
	}
	return result
}

func productExceptSelf2(nums []int) []int {
	result := make([]int, len(nums))

	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}

	rightProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		result[i] = result[i] * rightProduct
		rightProduct = rightProduct * nums[i]
	}

	return result
}
