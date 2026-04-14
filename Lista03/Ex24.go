package main

import "fmt"

func main() {
	fmt.Println("Angulo (rad) | Seno (Mac-Laurin)")
	fmt.Println("-------------------------------")

	for a := 0.0; a <= 6.31; a += 0.1 {
		seno := a - (a*a*a)/6.0 + (a*a*a*a*a)/120.0 - (a*a*a*a*a*a*a)/5040.0
		fmt.Printf("%.1f          | %.4f\n", a, seno)
	}
}
