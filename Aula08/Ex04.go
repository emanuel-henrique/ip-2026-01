package main

import "fmt"

func main() {
	fmt.Print(fatorial(5))
}

func fatorial(n int) int {
	if n <= 0 {
		return 0
	}
	fmt.Println(n)
	return fatorial(n - 1)
}
