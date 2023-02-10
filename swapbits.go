package piscine

func swapbits(oct byte) byte {
	return oct<<4 | oct>>4
}
