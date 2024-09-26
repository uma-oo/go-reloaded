package reloaded

import (
	"strconv"
	"strings"
	"unicode"
)

func isValidBin(num string) bool {
	runes := []rune(num)
	result := true
	for i := 0; i < len(runes); i++ {
		if runes[i] == '0' || runes[i] == '1' {
			continue
		} else {
			return false
		}
	}
	return result
}

func isValidHex(num string) bool {
	runes := []rune(num)
	result := true
	for i := 0; i < len(runes); i++ {
		if i == 0 && (runes[i] == '-' || runes[i] == '+') { // weird fact if we replace it with ( runes[0] == '-' || runes[0] == '+')
			continue
		} else if (runes[i] >= '0' && runes[i] <= '9') || (runes[i] >= 'A' && runes[i] <= 'F') || (runes[i] >= 'a' && runes[i] <= 'f') {
			continue
		} else {
			return false
		}
	}
	return result
}

func hex(number string) string {
	if isValidHex(number) {
		nbr_str, _ := strconv.ParseInt(number, 16, 64)
		s := strconv.FormatInt(nbr_str, 10)

		return s
	} else {
		return number
	}
}

func bin(number string) string {
	if isValidBin(number) {
		nbr_str, _ := strconv.ParseInt(number, 2, 64)
		s := strconv.FormatInt(nbr_str, 10)

		return s

	} else {
		return number
	}
}

// func isALphaNumeric(r rune) bool {
// 	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9')
// }

func cap(s string) string {
	runes := []rune(s)
	start := true
	for i, r := range runes {
		{
			if start {
				if unicode.IsLetter(r) {
					runes[i] = unicode.ToUpper(r)
					start = false
				} 
				
			} else {
				runes[i] = unicode.ToLower(r)
			}
		}
	}
	return string(runes)
}

func up(word string) string {
	word_runes := []rune(word)
	for i := 0; i < len(word_runes); i++ {
		{
			if unicode.IsUpper(word_runes[i]) {
				continue
			} else {
				word_runes[i] = unicode.ToUpper(word_runes[i])
			}
		}
	}
	return string(word_runes)
}

func low(word string) string {
	word_runes := []rune(word)
	for i := 0; i < len(word_runes); i++ {
		{
			if unicode.IsLower(word_runes[i]) {
				continue
			} else {
				word_runes[i] = unicode.ToLower(word_runes[i])
			}
		}
	}
	return string(word_runes)
}

func executeOperation(k string, word string) string {
	var result string
	sp := "\\'(){}[]"

	switch k {
	case "up":
		result = up(word)

	case "low":
		result = low(word)
	case "hex":
		result = hex(word)
	case "bin":
		result = bin(word)
	case "cap":
		if strings.Contains(sp, string(word[0])) {
			result = string(word[0]) + cap(word[1:])
		} else {
			result = cap(word)
		}

	}
	return result
}

func applyOneOperation(flag string, data_splitted []string, index int) {
	res := executeOperation(flag, data_splitted[index])
	data_splitted[index] = res
}

func applyMultipleOperations(flag string, data_splitted []string, index int, num_of_iterations int) {
	var res string
	if num_of_iterations > len(data_splitted[0:index+1]) {
		for j := 0; j < len(data_splitted[0:index+1]); j++ {
			res = executeOperation(flag, data_splitted[j])
			data_splitted[j] = res

		}
	} else {
		count := 0
		for j := len(data_splitted[:index+1]); j >= 0; j-- {
			if count <= num_of_iterations {
				res = executeOperation(flag, data_splitted[j])
				data_splitted[j] = res
				count++
			}
		}
	}
}

func isFlag(str string) bool {
	return str == "(cap)" || str == "(up)" || str == "(low)" || str == "(hex)" || str == "(bin)"
}

func isFlagOp(operation string, number string) bool {
	var result bool
	if (operation == "(cap," || operation == "(up," || operation == "(low,") && string(number[len(number)-1]) == ")" {
		_, err := strconv.Atoi(number[:len(number)-1])
		if err != nil {
			result = false
		} else {
			result = true
		}

	}
	return result
}

func ApplyOperations(data string) string {
	data_splitted := SplitWhiteSpaces(data)

	// new_str:=""
	for i := 0; i < len(data_splitted); i++ {
		if isFlag(data_splitted[i]) && i == 0 {
			data_splitted = append(data_splitted[:i], data_splitted[i+1:]...)
			i = -1
		} else if i == 0 && i+1 < len(data_splitted) && isFlagOp(data_splitted[i], data_splitted[i+1]) {
			data_splitted = append(data_splitted[:i], data_splitted[i+2:]...)
			i = -1
		} else if i+1 < len(data_splitted) && data_splitted[i+1] == "(cap)" {
			applyOneOperation("cap", data_splitted, i)
			data_splitted = append(data_splitted[:i+1], data_splitted[i+2:]...)
			i = -1

		} else if i+1 < len(data_splitted) && data_splitted[i+1] == "(up)" {
			applyOneOperation("up", data_splitted, i)
			data_splitted = append(data_splitted[:i+1], data_splitted[i+2:]...)
			i = -1
		} else if i+1 < len(data_splitted) && data_splitted[i+1] == "(low)" {
			applyOneOperation("low", data_splitted, i)
			data_splitted = append(data_splitted[:i+1], data_splitted[i+2:]...)
			i = -1
		} else if i+1 < len(data_splitted) && data_splitted[i+1] == "(hex)" {
			applyOneOperation("hex", data_splitted, i)
			data_splitted = append(data_splitted[:i+1], data_splitted[i+2:]...)
			i = -1
		} else if i+1 < len(data_splitted) && data_splitted[i+1] == "(bin)" {
			applyOneOperation("bin", data_splitted, i)
			data_splitted = append(data_splitted[:i+1], data_splitted[i+2:]...)
			i = -1
		} else if i+1 < len(data_splitted) && data_splitted[i+1] == "(cap," && i+2 < len(data_splitted) && data_splitted[i+2][len(data_splitted[i+2])-1] == ')' {
			num_of_iterations, err := strconv.Atoi(data_splitted[i+2][:len(data_splitted[i+2])-1])
			if err != nil {
				continue
			} else {
				applyMultipleOperations("cap", data_splitted, i, num_of_iterations)
				data_splitted = append(data_splitted[0:i+1], data_splitted[i+3:]...)
				i = -1
			}
		} else if i+1 < len(data_splitted) && data_splitted[i+1] == "(up," && i+2 < len(data_splitted) && data_splitted[i+2][len(data_splitted[i+2])-1] == ')' {

			num_of_iterations, err := strconv.Atoi(data_splitted[i+2][:len(data_splitted[i+2])-1])
			if err != nil {
				continue
			} else {
				applyMultipleOperations("up", data_splitted, i, num_of_iterations)
				data_splitted = append(data_splitted[0:i+1], data_splitted[i+3:]...)
				i = -1
			}

		} else if i+1 < len(data_splitted) && data_splitted[i+1] == "(low," && i+2 < len(data_splitted) && data_splitted[i+2][len(data_splitted[i+2])-1] == ')' {
			num_of_iterations, err := strconv.Atoi(data_splitted[i+2][:len(data_splitted[i+2])-1])
			if err != nil {
				continue
			} else {
				applyMultipleOperations("low", data_splitted, i, num_of_iterations)
				data_splitted = append(data_splitted[0:i+1], data_splitted[i+3:]...)
				i = -1

			}
		}
	}
	return Join(data_splitted, " ")
}
