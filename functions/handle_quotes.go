package reloaded

import (
	"errors"
	"strings"
)

func between(before rune, after rune) bool {
	return IsALphaNumeric(before) && IsALphaNumeric(after)
}

// func isSpace(r rune) bool {
// 	return r == ' '
// }

func isPair(number int) bool {
	return number%2 == 0
}

func HandleQuotes(sentence string) (string, error) {
	words := []rune(sentence)
	result := ""
	is_open := false
	var start int
	var end int
	var word string
	count_quotes := 0
	is_closed := true
	for i := 0; i < len(words); i++ {
		if words[i] == '\'' && i+1 < len(words) && i-1 >= 0 && between(words[i-1], words[i+1]) && is_closed {
			result += "'"
		} else if words[i] == '\'' && i+1 < len(words) && i-1 >= 0 && between(words[i-1], words[i+1]) && !is_closed {
			continue
		} else {
			if words[i] == '\'' && !is_open { // the debut of the quote
				if i-1 >= 0 { // if the quote b7al hakka ouma' don't'
					result += " "
				}

				is_open = true
				start = i
				is_closed = false
				count_quotes++
			} else if words[i] == '\'' && is_open {
				end = i
				word = string(words[start+1 : end])
				is_open = false

				result += "'" + strings.TrimSpace(word) + "'" + " " // if the character after is '

				word = ""
				is_closed = true
				count_quotes++
			} else if is_closed {
				result += string(words[i])
			}
		}
	}

	if !isPair(count_quotes) {
		return sentence, errors.New("error: there is no ending quote")
	} else {
		return Join(SplitWhiteSpaces(result), " "), nil
	}
}
