package main

import "fmt"

func main() {
	var janela [24]int
	var corredor [24]int
	var opcao int

	for {
		fmt.Scan(&opcao)
		if opcao == 0 {
			break
		}

		cheioJanela := true
		cheioCorredor := true
		for i := 0; i < 24; i++ {
			if janela[i] == 0 { cheioJanela = false }
			if corredor[i] == 0 { cheioCorredor = false }
		}

		if cheioJanela && cheioCorredor {
			fmt.Println("Ônibus lotado")
			continue
		}

		if opcao == 1 {
			if cheioJanela {
				fmt.Println("Janelas lotadas")
			} else {
				for i := 0; i < 24; i++ {
					if janela[i] == 0 {
						fmt.Printf("Poltrona %d disponível na Janela. Reservar? (1-Sim/0-Não): ", i+1)
						var res int
						fmt.Scan(&res)
						if res == 1 {
							janela[i] = 1
							fmt.Println("Venda efetuada")
							break
						}
					}
				}
			}
		} else if opcao == 2 {
			if cheioCorredor {
				fmt.Println("Corredores lotados")
			} else {
				for i := 0; i < 24; i++ {
					if corredor[i] == 0 {
						fmt.Printf("Poltrona %d disponível no Corredor. Reservar? (1-Sim/0-Não): ", i+1)
						var res int
						fmt.Scan(&res)
						if res == 1 {
							corredor[i] = 1
							fmt.Println("Venda efetuada")
							break
						}
					}
				}
			}
		}
	}
}
