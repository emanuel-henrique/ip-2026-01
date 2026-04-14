package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Digite o valor de X: ")
	fmt.Scan(&x)
	var s float64
	f := 1.0
	for i := 1; i <= 20; i++ {
		f *= float64(i)
		s += x / f
	}
	fmt.Printf("Resultado da série: %.6f\n", s)
}
