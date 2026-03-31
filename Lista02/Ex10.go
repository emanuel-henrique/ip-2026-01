package main

import "fmt"

func main() {
	var destino int
	var isIdaeVolta int
	
	var custo int

	fmt.Println("Digite o destino para a sua viagem!")
	fmt.Println("[1] - Região Norte")
	fmt.Println("[2] - Região Nordeste")
	fmt.Println("[3] - Região Centro-Oeste")
	fmt.Println("[4] - Região Sul")
	fmt.Print("=========================================\nDestino: ")
	fmt.Scan(&destino)
	fmt.Print("=========================================\n")
	fmt.Println("É uma viagem de ida e volta?")
	fmt.Println("[0] - Não")
	fmt.Println("[1] - Sim")
	fmt.Print("=========================================\nResposta: ")
	fmt.Scan(&isIdaeVolta)
	
	if isIdaeVolta == 0 {
	    if destino == 1 {
	        custo = 500
	    } else if destino == 2 || destino == 3{
	        custo = 350
	    } else if destino == 4 {
	        custo = 300
	    }
	    
	} else if isIdaeVolta == 1 {
	    if destino == 1 {
	        custo = 900
	    } else if destino == 2 {
	        custo = 650
	    } else if destino == 3 {
	        custo = 600
	    } else if destino == 4 {
	        custo = 550
	    }
	}
	
	fmt.Print(custo)
}
