package main

import (
	"fmt"
)

func main() {
	var n1, n2 int

	for {

		_, err := fmt.Scanf("%d %d", &n1, &n2)

		if err != nil {
			break
		}

		fmt.Println(n1 ^ n2)
	}

	return

}
