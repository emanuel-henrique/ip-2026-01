package main

import "fmt"

func main() {
    var destino int
    var isIdaeVolta int
	
	fmt.Println("Digite o destino para a sua viagem!")
	fmt.Println("[1] - Região Norte")
	fmt.Println("[2] - Região Nordeste")
	fmt.Println("[3] - Região Centro-Oeste")
	fmt.Println("[4] - Região Sul")
	fmt.Print("=========================================\n Destino: ")
	fmt.Scan(&destino)
	fmt.Print("=========================================\n")
	fmt.Println("É uma viagem de ida e volta?")
	fmt.Println("[0] - Sim")
	fmt.Println("[1] - Não")
	fmt.Print("=========================================\n Resposta: ")
	fmt.Scan(&isIdaeVolta)
	
	
	
}
