package main

import (
	"fmt"
	"strconv"
)

func main() {
	var cases int
	var number string
	var necessary_leds int

	leds := map[rune]int{
		'1': 2,
		'2': 5,
		'3': 5,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 3,
		'8': 7,
		'9': 6,
		'0': 6,
	}

	fmt.Scanf("%d", &cases)

	for i := cases; i > 0; i-- {
		necessary_leds = 0

		fmt.Scanf("%s", &number)

		for _, char := range number {
			necessary_leds += leds[char]
		}

		fmt.Println(strconv.Itoa(necessary_leds) + " leds")
	}
}
