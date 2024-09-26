package reloaded


func isVowel(r string) bool {
	return r == "a" || r == "e" || r == "i" || r == "o" || r == "u" || r == "h" || r == "A" || r == "E" || r == "I" || r == "O" || r == "U" || r == "H"
}

func HandleVowels(sentence string) string {
	words := SplitWhiteSpaces(sentence)
	for i := 0; i < len(words); i++ {
		if words[i] == "a" && i+1 < len(words) && isVowel(string(words[i+1][0])) {
			words[i] = "an"
		} else if words[i] == "A" && i+1 < len(words) && isVowel(string(words[i+1][0])) {
			words[i] = "An"
		} else {
			continue
		}
	}
	return Join(words, " ")
}
