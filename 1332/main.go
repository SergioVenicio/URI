package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var cases int
	scanner := bufio.NewReader(os.Stdin)

	fmt.Scanf("%d", &cases)

	for i := cases; i > 0; i-- {
		number, _, _ := scanner.ReadLine()

		if len(number) == 5 {
			fmt.Println("3")
		} else if number[0] == 'o' || number[2] == 'e' {
			fmt.Println("1")
		} else {
			fmt.Println("2")
		}
	}
}
