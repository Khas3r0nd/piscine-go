package piscine

func Split(str, charset string) []string {
	var res []string
	word := ""
	for i := 0; i < len(str); i++ {
		if i+len(charset) <= len(str) && str[i:i+len(charset)] == charset {
			if word != "" {
				res = append(res, word)
				word = ""
			}
			i += len(charset) - 1
			continue
		}
		word += string(str[i])
	}
	if word != "" {
		res = append(res, word)
	}
	return res
}
