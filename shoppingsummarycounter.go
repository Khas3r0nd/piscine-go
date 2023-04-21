package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	var stringArray []string
	m := make(map[string]int)
	lastIndex := 0
	for index, char := range str {
		if char == ' ' {
			stringArray = append(stringArray, str[lastIndex:index])
			lastIndex = index + 1
		}
		if index == len(str)-1 {
			stringArray = append(stringArray, str[lastIndex:])
		}
	}
	for _, str := range stringArray {
		m[str] = m[str] + 1
	}
	if str == "" {
		m[""] = 1
	}
	return m
}
