package main

import (
	"fmt"
	"math"
)

func main() {
	var S float64 = 0.0
	sinal := 1.0

	for i := 0; i < 51; i++ {
		impar := float64(2*i + 1)
		termo := sinal * (1.0 / math.Pow(impar, 3))
		S += termo
		sinal *= -1.0 
	}

	pi_aprox := math.Cbrt(S * 32.0)

	fmt.Printf("Valor aproximado de PI (51 termos): %.15f\n", pi_aprox)
	fmt.Printf("Valor real de PI (para comparar):   %.15f\n", math.Pi)
}
