package main

import "fmt"

// Conversões para o Sistema Métrico

func main() {
	var temperature float64
	var rain float64

	fmt.Print("Digite a temperatura a ser convertida (FAHRENHEIT): ")
	fmt.Scan(&temperature)

	fmt.Print("Digite o volume de chuva a ser convertido (POLEGADA): ")
	fmt.Scan(&rain)

	temperature = ((temperature - 32) * 5) / 9 
	rain = rain * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", temperature)
	
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f", rain)
}