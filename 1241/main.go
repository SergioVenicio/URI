package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var cases int

	fmt.Scanf("%d", &cases)

	scanner := bufio.NewReader(os.Stdin)

	for i := cases; i > 0; i-- {
		fmt.Printf(string(cases))
		line, _, err := scanner.ReadLine()

		if err != nil {
			panic(err)
		}

		arr_strings := strings.Split(string(line), " ")

		if arr_strings[1] == "" {
			fmt.Println("nao encaixa")
		} else if strings.HasSuffix(arr_strings[0], arr_strings[1]) {
			fmt.Println("encaixa")
		} else {
			fmt.Println("nao encaixa")
		}

	}
}
