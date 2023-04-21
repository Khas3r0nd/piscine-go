package piscine

func MakeRange(min, max int) []int {
	if max-min < 1 {
		return nil
	}
	answer := make([]int, max-min)
	counter := 0
	for i := min; i < max; i++ {
		answer[counter] = i
		counter++
	}
	return answer
}
