package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scan(&t)

	for i := 1; i <= t; i++ {
		var n, popular, geral, arquibancada, cadeiras float64
		
		fmt.Scan(&n, &popular, &geral, &arquibancada, &cadeiras)

		renda := (n * popular / 100 * 1.00) +
			(n * geral / 100 * 5.00) +
			(n * arquibancada / 100 * 10.00) +
			(n * cadeiras / 100 * 20.00)

		fmt.Printf("A RENDA DO JOGO N. %d E = %.2f\n", i, renda)
	}
}
