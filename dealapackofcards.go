package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	for i := 0; i <= 3; i++ {
		fmt.Printf("Player %v: %v, %v, %v\n", i+1, deck[i*3], deck[i*3+1], deck[i*3+2])
	}
}
