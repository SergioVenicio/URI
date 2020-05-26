package main

import "fmt"

func main() {
	var cases int
	var ric, vic int

	fmt.Scanf("%d", &cases)

	for i := cases; i > 0; i-- {
		fmt.Scanf("%d %d", &ric, &vic)

		if ric%vic == 0 {
			fmt.Println(vic)
		} else {
			rest := 0
			for ric%vic != 0 {
				rest = ric % vic
				ric, vic = vic, rest
			}
			fmt.Println(rest)
		}
	}
}
