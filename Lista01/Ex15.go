package main

import "fmt"

// Reajuste salarial
func main(){
	var salario float64

	fmt.Print("Digite o salário a ser reajustado: ")
	fmt.Scan(&salario)

	if salario <= 300{
		fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salario * 1.5)
	} else if salario > 300{
		fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", salario * 1.3)
	}
}