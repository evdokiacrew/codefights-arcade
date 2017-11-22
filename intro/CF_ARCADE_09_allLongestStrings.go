/*
intro 09/60

Given an array of strings, return another array containing all of its longest strings.
*/

func allLongestStrings(inputArray []string) []string {
	var resultArray []string
	maximumLength := 0
	for i := 0; i < len(inputArray); i++ {
		if len(inputArray[i]) > maximumLength {
			maximumLength = len(inputArray[i])
		}
	}
	for i := 0; i < len(inputArray); i++ {
		if len(inputArray[i]) == maximumLength {
			resultArray = append(resultArray, inputArray[i])
		}
	}
	return resultArray
}
