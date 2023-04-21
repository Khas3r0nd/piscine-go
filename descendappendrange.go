package piscine

func DescendAppendRange(max, min int) []int {
	if max < min {
		return []int{}
	}
	intArr := []int{}
	for i := max; i > min; i-- {
		intArr = append(intArr, i)
	}
	return intArr
}
