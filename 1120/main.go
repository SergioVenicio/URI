package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var broken_digit string
	var contract_value string

	for {
		fmt.Scanf("%s %s", &broken_digit, &contract_value)

		if broken_digit == "0" || contract_value == "0" {
			break
		}

		string_value := strings.Replace(contract_value, broken_digit, "", -1)

		if string_value == "" {
			string_value = "0"
		}

		value, err := strconv.Atoi(string_value)

		if err != nil {
			fmt.Println(string_value)
		} else {
			fmt.Println(value)
		}
	}
}
