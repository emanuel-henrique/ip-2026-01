package main

import "fmt"

func main() {
	var n int
	var s float64 = 0
	numerador := 1000.0

	fmt.Print("Informe o numero de termos N: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Valor invalido!")
	} else {
		for i := 1; i <= n; i++ {
			if i%2 != 0 {
				s += numerador / float64(i)
			} else {
				s -= numerador / float64(i)
			}
			numerador -= 3
		}
		fmt.Printf("Resultado da serie: %.2f\n", s)
	}
}
