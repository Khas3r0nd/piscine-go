package piscine

func StringToIntSlice(str string) []int {
	var intArr []int
	for _, value := range str {
		intArr = append(intArr, int(value))
	}
	return intArr
}
