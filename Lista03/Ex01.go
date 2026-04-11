// Escreva um programa que calcule potências. O usuário
// deve digitar a base e o expoente, e o programa deve
// apresentar o resultado (sem usar o comando pow).
// Assuma que o usuário irá digitar valores positivos.
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
	for i := 0; i < e; i++ {
		result *= b
	}
	return result
}