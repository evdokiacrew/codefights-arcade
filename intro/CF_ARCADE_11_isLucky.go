/*
intro 11/60

Ticket numbers usually consist of an even number of digits.
A ticket number is considered lucky if the sum of the first half of the digits is equal to the sum of the second half.
Given a ticket number n, determine if it's lucky or not.
*/

func isLucky(n int) bool {
	number := strconv.Itoa(n)
	half := len(number)/2
	left, right := 0, 0
	for i := 0; i < half; i++ {
		one, _ := strconv.Atoi(string(number[i]))
		left += one
	}
	for i := half; i < len(number); i++ {
		two, _ := strconv.Atoi(string(number[i]))
		right += two
	}
	fmt.Println(left, right)
	return left == right
}
