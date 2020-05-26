package main

import "fmt"

func main() {
	var tests int
	var positions int
	var value string
	var resturn_value string

	fmt.Scanf("%d", &tests)

	for i := 0; i < tests; i++ {
		resturn_value = ""

		fmt.Scanf("%s", &value)
		fmt.Scanf("%d", &positions)

		for _, char := range value {
			char_pos := int(char) - positions

			if char_pos < 65 {
				char_pos += 26
			}

			resturn_value += string(char_pos)
		}

		fmt.Println(resturn_value)
	}
}
