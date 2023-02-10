package piscine

func reverseBits(num byte) byte {
	var res byte
	for i := 0; i < 8; i++ {
		res = res<<1 + num&1
		num = num >> 1
	}
	return res
}
