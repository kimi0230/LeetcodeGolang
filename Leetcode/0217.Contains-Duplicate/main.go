package containsduplicate

func ContainsDuplicate(nums []int) bool {
	numsMap := make(map[int]bool, len(nums))
	for _, v := range nums {
		if numsMap[v] {
			return true
		}
		numsMap[v] = true
	}
	return false
}
