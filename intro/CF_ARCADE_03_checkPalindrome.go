/*
intro 03/60 checkPalindrome

Given the string, check if it is a palindrome.
*/

func checkPalindrome(inputString string) bool {
	for i := 0; i < len(inputString)/2; i++ {
		if inputString[i] != inputString[len(inputString)-i-1] {
			return false
		}
	}
	return true
}
