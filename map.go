package piscine

func Map(f func(int) bool, a []int) []bool {
	answers := make([]bool, len(a))
	for i, value := range a {
		answers[i] = f(value)
	}
	return answers
}
