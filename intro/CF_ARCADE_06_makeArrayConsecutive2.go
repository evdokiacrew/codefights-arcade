/*
intro 06/60 makeArrayConsecutive2

Ratiorg got statues of different sizes as a present from CodeMaster for his birthday, each statue having an non-negative integer size.
Since he likes to make things perfect, he wants to arrange them from smallest to largest
so that each statue will be bigger than the previous one exactly by 1.
He may need some additional statues to be able to accomplish that. Help him figure out the minimum number of additional statues needed.
*/

func makeArrayConsecutive2(statues []int) int {
	min := statues[0]
	max := statues[0]
	for i := 1; i < len(statues); i++ {
		if statues[i] < min {
			min = statues[i]
		}
		if statues[i] > max {
			max = statues[i]
		}
	}
	fmt.Println(min, max)
	return max - min - len(statues) + 1
}
