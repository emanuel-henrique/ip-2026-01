package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	var n int
	fmt.Print("Digite a quantidade de termos (N >= 3): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		if i == 0 {
			fmt.Printf("%d", fibonacci(i))
		} else {
			fmt.Printf(" – %d", fibonacci(i))
		}
	}
	fmt.Println()
}