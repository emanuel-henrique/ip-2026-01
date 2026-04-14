package main

import "fmt"

func main() {
	var n int
	fmt.Print("Quantos termos de Fetuccine deseja ver? ")
	fmt.Scan(&n)

	a, b := 1, 1
	fmt.Println("Série de Fetuccine:")
	fmt.Print(a, " ", b, " ")

	for i := 3; i <= n; i++ {
		var c int
		if i%2 == 0 {
			c = a - b
		} else {
			c = a + b
		}
		fmt.Print(c, " ")
		a = b
		b = c
	}
}
