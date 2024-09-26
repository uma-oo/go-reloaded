package reloaded

func isPonctuation(r string) bool {
	return r == "," || r == ";" || r == "!" || r == "?" || r == "." || r == ":"
}

func HandlePonctuation(sentence string) string {
	words := SplitWhiteSpaces(sentence)

	for i := 0; i < len(words); i++ {
		if i+1 < len(words) && isPonctuation(string(words[i+1][0])) {
			words[i] = words[i] + string(words[i+1][0])
			words[i+1] = words[i+1][1:]
			sentence = Join(words, " ")
			words = SplitWhiteSpaces(sentence)
			i = -1
		}
	}
	return Join(words, " ")
}
