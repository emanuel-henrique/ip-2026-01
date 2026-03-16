package main

import "fmt"

// Locadora de charretes
func main() {
	var hours int

	fmt.Print("Digite a quantidade de horas de uso: ")
	fmt.Scan(&hours)

	fmt.Printf("O VALOR A PAGAR E = %d.00", calculatePrice(hours))
}

func calculatePrice(hours int) int {
	threeHoursBlocks := hours / 3
	oneHourBlocks := hours % 3

	price := (threeHoursBlocks * 10) + (oneHourBlocks * 5)
	return price
}