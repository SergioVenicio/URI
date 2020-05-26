package main

import (
	"bufio"
	"fmt"
	"os"
)

func ParseLine(line []byte) string {
	var return_str string
	prev := " "

	for _, char := range string(line) {
		if prev == " " && string(char) != " " {
			return_str += string(char)
			prev = string(char)
		} else {
			prev = string(char)
		}
	}

	return return_str
}

func main() {
	var counter int
	reader := bufio.NewReader(os.Stdin)

	fmt.Scanf("%d", &counter)

	for i := counter; i > 0; i-- {
		line, _, _ := reader.ReadLine()
		fmt.Println(ParseLine(line))
	}
}
