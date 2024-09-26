package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	re "reloaded/functions"
)

func main() {
	var result string
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Incorrect arguments, Please provide input and output file!")
		return
	} else if !strings.HasSuffix(args[0], ".txt") {
		fmt.Println("The input file given has the wrong extension!")
		return
	} else {
		input_file := args[0]

		file, err := os.Open(input_file)
		if err != nil {
			fmt.Println("There was a problem while Opening this file!")
			return
		}
		defer file.Close()
        output_file:=re.SplitPath(args[1])
		newfile, err := re.CreateFile(output_file[len(output_file)-1], "test")
		if err != nil {
			return
		}
		defer newfile.Close()

		content := bufio.NewScanner(file)
        
		lines := 0
		for content.Scan() {

			if lines == 0 {

				_, err := io.WriteString(newfile, "")
				lines += 1
				if err != nil {
					return
				}
			} else {
				_, err := io.WriteString(newfile, "\n")
				lines += 1
				if err != nil {
					return
				}
			}

			scantext := content.Text()

			res := re.ApplyOperations(string(scantext))

			res__ponctuation := re.HandlePonctuation(res)
			res_quotes, error_quotes := re.HandleQuotes(res__ponctuation)
			if error_quotes != nil {
				result = re.HandleVowels(res_quotes)
				_, err := io.WriteString(newfile, result)
				fmt.Printf("%s at line %d", error_quotes, lines)
				fmt.Println()
				if err != nil {
					return
				}

			} else {
				result = re.HandleVowels(res_quotes)
				_, err := io.WriteString(newfile, result)
				if err != nil {
					return
				}
			}

		}

		if err := content.Err(); err != nil {
			return
		}

	}
}
