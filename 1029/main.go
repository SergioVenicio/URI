package main

import "fmt"

func fibo(n int, calls int) int {
	calls++

	if n < 2 {
		return 1
	}

	return fibo(n, calls)
}

func main() {
	fmt.Println("vim-go")
}
