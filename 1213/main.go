package main

import (
	"fmt"
	"math"
)

func main() {
	var n int64

	for {
		_, err := fmt.Scanf("%d", &n)
		if err != nil {
			break
		}

		aux := 1
		i := 1
		if n%2 != 0 && n%5 != 0 {
			for aux%int(n) != 0 {
				aux = aux % int(n)
				aux += int(math.Pow(10, float64(i)))
				i++
			}
			fmt.Println(i)
		}
	}
}
