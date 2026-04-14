package main

import "fmt"

func main() {
	var n int
	fmt.Print("Quantos termos de Fibonacci deseja ver? ")
	fmt.Scan(&n)
	a, b := 0, 1
	fmt.Println("Sequência:")
	for i := 0; i < n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
}
