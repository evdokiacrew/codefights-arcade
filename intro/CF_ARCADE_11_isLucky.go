/*
intro 11/60

Ticket numbers usually consist of an even number of digits.
A ticket number is considered lucky if the sum of the first half of the digits is equal to the sum of the second half.
Given a ticket number n, determine if it's lucky or not.
*/

func isLucky(n int) bool {
	number := strconv.Itoa(n)
	half := len(number)/2
	left := 0
	right := 0
	for i := 0; i < half; i++ {
        one := strconv.Atoi(number[i])
        left += one
	}
	for i := half - 1; i < len(number); i++ {
		right += strconv.Atoi(number[i])
	}
    if left == right {
    	return true
    } else {
    	return false
    }
}
