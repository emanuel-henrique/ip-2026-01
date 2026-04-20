package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Print("Digite o valor de x (em radianos): ")
	fmt.Scan(&x)
  
	cosSerie := 1.0
	termo := 1.0

	for i := 1; i < 20; i++ {
		termo = termo * (-1.0) * x * x / (float64(2*i) * float64(2*i-1))
		cosSerie += termo
	}

	cosNativo := math.Cos(x)
	diferenca := math.Abs(cosSerie - cosNativo)

	fmt.Printf("Cosseno pela série (20 termos): %.15f\n", cosSerie)
	fmt.Printf("Cosseno pela função COS(x):     %.15f\n", cosNativo)
	fmt.Printf("Diferença entre os valores:     %.15f\n", diferenca)
}
