package main

import "fmt"

func main() {
	var a, b float64
	fmt.Print("Digite dois números e te darei a soma deles: ")
	fmt.Scan(&a, &b)

	fmt.Printf("A soma deles é: %.2f\n", soma(a, b))
}

func soma(a, b float64) float64 {
	return a + b
}
