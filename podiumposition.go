package piscine

func PodiumPosition(podium [][]string) [][]string {
	return SortString(podium)
}

func SortString(a [][]string) [][]string {
	for i := 0; i < len(a)-1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i][0] > a[j][0] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
	return a
}
