package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func first_step(string_value string) string {
	var encripted_string string

	for i := 0; i < len(string_value); i++ {
		if unicode.IsLower(rune(string_value[i])) || unicode.IsUpper(rune(string_value[i])) {
			encripted_string += string(int(string_value[i]) + 3)
		} else {
			encripted_string += string(string_value[i])
		}
	}

	return encripted_string
}

func second_step(string_value string) string {
	var encripted_string string

	for i := len(string_value); i > 0; i-- {
		encripted_string += string(string_value[i-1])
	}

	return encripted_string
}

func third_step(string_value string) string {
	var encripted_string string
	middle := int(len(string_value) / 2)

	for i := 0; i < middle; i++ {
		encripted_string += string(string_value[i])
	}

	for i := middle; i < len(string_value); i++ {
		encripted_string += string(int(string_value[i]) - 1)
	}

	return encripted_string
}

func encript_string(string_value string) string {
	var encripted_string string

	encripted_string = first_step(string_value)
	encripted_string = second_step(encripted_string)
	encripted_string = third_step(encripted_string)

	return encripted_string
}

func main() {
	var counter int

	scanner := bufio.NewReader(os.Stdin)

	fmt.Scanf("%d", &counter)

	for i := 0; i < counter; i++ {
		var encripted_string string

		string_input, _, err := scanner.ReadLine()

		if err != nil {
			break
		}

		encripted_string = encript_string(string(string_input))
		fmt.Println(encripted_string)
	}
}
