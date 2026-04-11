package main

import "fmt"

func main(){
	var base float64
	var expoente int

	fmt.Print("Digite a base: ")
	fmt.Scan(&base)

	fmt.Print("Digite o expoente: ")
	fmt.Scan(&expoente)

	fmt.Print(calcularPotencia(base,expoente))
}

func calcularPotencia(b float64, e int) float64{
	var result float64 = 1
	for range e {
		result *= b
	}
	return result
}