package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func Rot13(user_input string) string {
	var return_string string
	alphabet := [...]rune{
		'a',
		'b',
		'c',
		'd',
		'e',
		'f',
		'g',
		'h',
		'i',
		'j',
		'k',
		'l',
		'm',
		'n',
		'o',
		'p',
		'q',
		'r',
		's',
		't',
		'u',
		'v',
		'w',
		'x',
		'y',
		'z',
	}

	for _, char := range user_input {

		aux := false
		for i := 0; i < len(alphabet); i++ {
			new_char := ""
			is_upper := unicode.IsUpper(char)

			if strings.ToLower(string(char)) == string(alphabet[i]) {

				if i+13 > 25 {
					new_char = string(alphabet[(i+13)-26])
				} else {
					new_char = string(alphabet[i+13])
				}
				aux = true

				if is_upper {
					return_string += strings.ToUpper(string(new_char))
				} else {
					return_string += string(new_char)
				}
			}
		}

		if !aux {
			return_string += string(char)
		}
	}

	return return_string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		user_input := scanner.Text()
		fmt.Println(Rot13(user_input))
	}
}
