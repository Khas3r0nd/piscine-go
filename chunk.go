package main

import "fmt"

func Chunk(slice []int, size int) {
	var rnn [][]int
	if size == 0 {
		fmt.Print("\n")
		return
	} else if len(slice) <= size {
		fmt.Println(slice)
		return
	}
	for i := 0; i < len(slice); i += size { // i = 0
		end := i + size // end = 3
		if end > len(slice) {
			end = len(slice)
		}
		rnn = append(rnn, (slice[i:end])) // slice[i:end]
	}
	fmt.Println(rnn)
}

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}
