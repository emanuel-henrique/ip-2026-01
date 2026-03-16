package main

import "fmt"

// Série de pares
func main(){
	var x,y int
	fmt.Print("Digite os dois numeros: ")
	fmt.Scan(&x, &y)

	if x%2 != 0 {
		fmt.Print("O PRIMEIRO NUMERO NAO E PAR")
		return
	}

	for i := 1; i <= y * 2; i++ {
		if i%2 == 0 {
			fmt.Print((i-2) + x, " ")
		}
	}
}