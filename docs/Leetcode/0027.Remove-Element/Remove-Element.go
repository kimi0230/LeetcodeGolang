package removeelement

/*
雙指針法
雙指針法（快慢指針法）在數組和鍊錶的操作中是非常常見的，很多考察數組和鍊錶操作的面試題，都使用雙指針法
*/
func RemoveElementDoublePoint(nums []int, val int) int {
	if len(nums) <= 0 {
		return 0
	}
	slowIndex := 0
	for fastIndex := 0; fastIndex < len(nums); fastIndex++ {
		if nums[fastIndex] != val {
			if fastIndex != slowIndex {
				nums[fastIndex], nums[slowIndex] = nums[slowIndex], nums[fastIndex]
			}
			slowIndex++
		}
	}
	return slowIndex
}

/*
RemoveElement :
*/
func RemoveElement(nums []int, val int) int {
	size := len(nums)
	i := 0
	for i < size {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			size--
			// fmt.Println(nums)
		} else {
			i++
		}
	}
	return len(nums)
}
