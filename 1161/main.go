package main

import (
	"fmt"
)

func main() {
	var n, n2 int64

	for {
		_, err := fmt.Scanf("%d %d", &n, &n2)
		if err != nil {
			break
		}

		fmt.Printf("%d\n", factorial(n)+factorial(n2))
	}
}

func factorial(n int64) int64 {
	if n <= 1 {
		return 1
	}

	return n * factorial(n-1)
}
