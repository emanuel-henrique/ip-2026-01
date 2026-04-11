package main

import "fmt"

func main() {
	var salario float64

	fmt.Print("Digite o salario de carlos: ")
	fmt.Scan(&salario)

	joao := (1.0/3.0) * salario
	meses := 0
	fmt.Println("Salario Joao: ", joao)

	for joao < salario{
		salario *= 1.02
		joao *= 1.05
		meses++
	}

	fmt.Print(meses)
}