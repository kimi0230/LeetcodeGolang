package makearrayconsecutive2

import "sort"

func makeArrayConsecutive2(statues []int) int {
	sort.Ints(statues)
	max := statues[len(statues)-1]
	min := statues[0]
	total := (max - min) + 1

	minAdditional := total - len(statues)

	return minAdditional

	// sort.Ints(statues)
	// max:=statues[len(statues)-1]
	// min:=statues[0]
	// currentIdx,result := 0,0
	// for i:= min; i<max ; i++{
	//     if statues[currentIdx] > i {
	//         result++
	//     }else{
	//         currentIdx++
	//     }
	// }
	// return result
}
