package main

import "fmt"

func main() {
	var number int

	fmt.Scanf("%d", &number)

	for i := 1; i <= 10000; i++ {
		if i%number == 2 {
			fmt.Printf("%d", i)
			fmt.Printf("\n")
		}
	}
}
