package adjacentelementsproduct

func adjacentElementEproduct(inputArray []int) int {
	maxProduct := inputArray[0] * inputArray[1]
	for i := 1; i < len(inputArray)-1; i++ {
		tmpProduct := inputArray[i] * inputArray[i+1]
		if tmpProduct > maxProduct {
			maxProduct = tmpProduct
		}
	}
	return maxProduct
}
