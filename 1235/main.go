package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverse(text string) string {
	len_text := len(text)
	result := make([]rune, len_text)
	for idx, _ := range text {
		len_text--
		result[idx] = rune(text[len_text])
	}

	return string(result)
}

func UndoHack(text string) string {
	return reverse(text[:len(text)/2]) + reverse(text[len(text)/2:])
}

func main() {
	var cases int
	scanner := bufio.NewReader(os.Stdin)

	fmt.Scanf("%d", &cases)

	for i := cases; i > 0; i-- {

		text, _, _ := scanner.ReadLine()

		fmt.Println(UndoHack(string(text)))

	}
}
