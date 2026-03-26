package main

import "fmt"

func main() {
	var valueC float64
	fmt.Print("Digite o valor da compra: ")
	fmt.Scan(&valueC)
	
	if valueC < 10 {
	    fmt.Printf("Valor de Venda: R$ %.2f", (valueC * 1.7))
	} else if valueC >= 10 && valueC < 30 {
	    fmt.Printf("Valor de Venda: R$ %.2f", (valueC * 1.5))
	} else if valueC >= 30 && valueC < 50 {
	    fmt.Printf("Valor de Venda: R$ %.2f", (valueC * 1.4))
	} else if valueC > 50 {
	    fmt.Printf("Valor de Venda: R$ %.2f", (valueC * 1.3)) 
	}
}
