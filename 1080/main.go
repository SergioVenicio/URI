package main

import "fmt"

func Max(numbers map[int]int) (int, int) {
	var maxValue, position = 0, 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > maxValue {
			maxValue = numbers[i]
			position = i + 1
		}
	}

	return maxValue, position
}

func main() {
	var n int
	numbers := make(map[int]int)
	for i := 0; i <= 100; i++ {
		fmt.Scanf("%d", &n)
		numbers[i] = n
	}

	max, index := Max(numbers)
	fmt.Println(max)
	fmt.Println(index)
}
