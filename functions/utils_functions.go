package reloaded

func IsAlphaString(str string) bool {
	result := true
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		if string(runes[i]) != " " {
			continue
		} else {
			return false
		}
	}

	return result
}

func IsAlpha(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

func IsALphaNumeric(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
}

func SplitWhiteSpaces(s string) []string {
	var chunks []string
	chunk := ""
	for _, el := range s {
		if string(el) != " " && string(el) != "\n"  && string(el) != "\t" {
			chunk += string(el)
		} else if chunk != "" { // skip the spaces if there are multiple ones
			chunks = append(chunks, chunk)
			chunk = ""
		}
	}

	if chunk != "" { // added because space may occur in the last of a ljumlaa ;)
		chunks = append(chunks, chunk) // added because the last one doesn't have a space after so we need to add it after the loop
	}
	return chunks
}

func Join(strs []string, sep string) string {
	new := ""
	for i := 0; i < len(strs); i++ {
		if i != len(strs)-1 {
			new += strs[i] + sep
		} else {
			new += strs[i]
		}
	}
	return new
}


func SplitPath(s string) []string {
	var chunks []string
	chunk := ""
	for _, el := range s {
		if string(el) != "/" {
			chunk += string(el)
		} else if chunk != "" { // skip the spaces if there are multiple ones
			chunks = append(chunks, chunk)
			chunk = ""
		}
	}

	if chunk != "" { // added because space may occur in the last of a ljumlaa ;)
		chunks = append(chunks, chunk) // added because the last one doesn't have a space after so we need to add it after the loop
	}
	return chunks
}