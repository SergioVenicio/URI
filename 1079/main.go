package main

import "fmt"

func calc_average(n1 float32, n2 float32, n3 float32) float32 {
	return ((n1 * 2) + (n2 * 3) + (n3 * 5)) / 10
}

func main() {
	var n1, n2, n3 float32
	var max_inputs int

	fmt.Scanf("%d", &max_inputs)

	for i := 0; i < max_inputs; i++ {
		fmt.Scanf("%f", &n1)
		fmt.Scanf("%f", &n2)
		fmt.Scanf("%f", &n3)

		fmt.Printf("%.1f\n", calc_average(n1, n2, n3))
	}

}
