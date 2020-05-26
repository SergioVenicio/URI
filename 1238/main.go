package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Combinator(line []string) string {
	var return_arr []byte
	first_str, second_str := line[0], line[1]
	first_len, second_len := len(first_str), len(second_str)

	if first_len >= second_len {
		for i := 0; i < second_len; i++ {
			return_arr = append(return_arr, first_str[i])
			return_arr = append(return_arr, second_str[i])
		}

		for i := second_len; i < first_len; i++ {
			return_arr = append(return_arr, first_str[i])
		}

	} else {
		for i := 0; i < first_len; i++ {
			return_arr = append(return_arr, first_str[i])
			return_arr = append(return_arr, second_str[i])
		}

		for i := first_len; i < second_len; i++ {
			return_arr = append(return_arr, second_str[i])
		}
	}

	return string(return_arr)
}

func main() {
	var counter int
	reader := bufio.NewReader(os.Stdin)

	fmt.Scanf("%d", &counter)

	for i := counter; i > 0; i-- {
		line, _, _ := reader.ReadLine()
		str_line := strings.Split(string(line), " ")
		fmt.Println(Combinator(str_line))
	}
}
