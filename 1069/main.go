package main

import (
	"bufio"
	"fmt"
	"os"
)

func ExtracDiamond(line []byte, initial_index int, final_index int) []byte {
	var new_line []byte

	for i, _ := range line {
		if i >= initial_index && i <= final_index {
			continue
		} else {
			new_line = append(new_line, line[i])
		}
	}
	return new_line
}

func GetDiamond(line []byte) ([]byte, bool) {
	var prev_char byte
	var has_diamond bool
	var has_init bool
	initial_index, final_index := 0, 0

	for index, char := range line {
		if char == '<' {
			initial_index = index
			has_init = true
		} else if char == '>' && (prev_char == '.' || prev_char == '<') && has_init {
			final_index = index
			has_diamond = true
		}

		if final_index != 0 {
			break
		}

		prev_char = char
	}

	var new_line []byte
	if has_diamond {
		new_line = ExtracDiamond(line, initial_index, final_index)
	}
	return new_line, has_diamond
}

func main() {
	var counter int
	reader := bufio.NewReader(os.Stdin)

	fmt.Scanf("%d", &counter)

	for i := counter; i > 0; i-- {
		diamonds := 0
		has_diamond := false

		line, _, _ := reader.ReadLine()

		for {
			line, has_diamond = GetDiamond(line)

			if has_diamond {
				diamonds++
			} else {
				fmt.Println(diamonds)
				break
			}
		}
	}
}
