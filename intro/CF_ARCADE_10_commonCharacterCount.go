/*
intro 10/60

Given two strings, find the number of common characters between them.
*/

func commonCharacterCount(s1 string, s2 string) int {
	count := 0
	r1 := []rune(s1)
	r2 := []rune(s2)
	fmt.Println(r1, r2)
	for i := 0; i < len(r1); i++ {
		for j := 0; j < len(r2); j++ {
			if r1[i] == r2[j] {
                r1[i] = 0
				r2[j] = 0
				count++
                break
			}
		}
	}
	return count
}
