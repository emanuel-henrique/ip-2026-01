package main

import (
	"fmt"
)

func main() {
	var n1, n2 int
	fmt.Print("Digite N1 e N2 (positivos): ")
	fmt.Scan(&n1, &n2)

	if n1 <= 0 || n2 <= 0 {
		fmt.Println("Os números devem ser inteiros positivos.")
		return
	}

	resto := n1
	quociente := 0

	for resto >= n2 {
		resto = resto - n2
		quociente++
	}

	fmt.Printf("Quociente(%d,%d) = %d\n", n1, n2, quociente)
	fmt.Printf("Resto(%d,%d) = %d\n", n1, n2, resto)
}
