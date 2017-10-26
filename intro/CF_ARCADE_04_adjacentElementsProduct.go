/*
intro 04/60 adjacentElementsProduct

Given an array of integers, find the pair of adjacent elements that has the largest product and return that product.
*/

func adjacentElementsProduct(inputArray []int) int {
	var largestProduct, resultProduct int
	largestProduct = inputArray[0] * inputArray[1]
	for i := 0; i < len(inputArray)-1; i++ {
		resultProduct = inputArray[i] * inputArray[i+1]
		if largestProduct < resultProduct {
			largestProduct = resultProduct
		}
	}
	return largestProduct
}
