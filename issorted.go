package piscine

func CompareInt(a, b int) int {
	return a - b
}

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) == 1 {
		return true
	}
	isAscending := a[0] < a[1]
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			switch isAscending {
			case true:
				if f(a[j], a[j+1]) > 0 {
					return false
				}
			case false:
				if f(a[j], a[j+1]) < 0 {
					return false
				}
			}
		}
	}
	return true
}
