package piscine

func Abort(a, b, c, d, e int) int {
	var intArr []int
	intArr = append(intArr, a, b, c, d, e)
	for i := 0; i < len(intArr)-1; i++ {
		for j := 0; j < len(intArr)-i-1; j++ {
			if intArr[j] > intArr[j+1] {
				intArr[j], intArr[j+1] = intArr[j+1], intArr[j]
			}
		}
	}
	return intArr[2]
}
