package main

import "fmt"

// Conta de Água
func main(){
	var codConta int

	fmt.Printf("Digite o código da conta: ")
	fmt.Scan(&codConta)

	var consumoAgua float64

	fmt.Printf("Digite o consumo de água (metros cúbicos): ")
	fmt.Scan(&consumoAgua)

	var tipoCosumidor string

	fmt.Printf("Digite o tipo de consumidor (Ex: R, C ou I): ")
	fmt.Scan(&tipoCosumidor)
	
	var valorConta float64

	if tipoCosumidor == "R" {
		valorConta = 5 + (0.05 * consumoAgua)
	} else if tipoCosumidor == "C" {

		if consumoAgua < 80 {
			valorConta = 6.25 * consumoAgua
		} else {
			valorConta = 500 + ((consumoAgua - 80) * 0.25)
		}

	} else if tipoCosumidor == "I" {

		if consumoAgua < 100 {
			valorConta = 8 * consumoAgua
			return
		} else { 
			valorConta = 800 + ((consumoAgua - 100) * 0.04)
		}
	}

	fmt.Println("CONTA =", codConta)
	fmt.Printf("VALOR DA CONTA = %.2f", valorConta)
}