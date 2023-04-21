package piscine

func UltimateDivMod(a *int, b *int) {
	extraA := (*a) / (*b)
	extraB := (*a) % (*b)
	*a = extraA
	*b = extraB
}
