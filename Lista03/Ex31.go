package main

import (
	"fmt"
	"math/big"
)

func main() {
	somaTotal := big.NewInt(0)
	graosNaCasa := big.NewInt(1)
	multiplicador := big.NewInt(2)

	for i := 1; i <= 64; i++ {
		somaTotal.Add(somaTotal, graosNaCasa)
		// Próxima casa tem o dobro
		if i < 64 {
			graosNaCasa.Mul(graosNaCasa, multiplicador)
		}
	}

	fmt.Printf("Total de grãos no tabuleiro de xadrez: %s\n", somaTotal.String())
}
