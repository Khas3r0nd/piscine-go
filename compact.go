package piscine

func Compact(ptr *[]string) int {
	new_ptr := []string{}
	for _, v := range *ptr {
		if v != "" {
			new_ptr = append(new_ptr, v)
		}
	}
	*ptr = new_ptr
	return len(new_ptr)
}
