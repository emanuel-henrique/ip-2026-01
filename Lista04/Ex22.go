package main

import "fmt"

func main() {
	var codigos [10]int
	var saldos [10]float64
	var opcao int

	for i := 0; i < 10; i++ {
		fmt.Scan(&codigos[i], &saldos[i])
	}

	for {
		fmt.Scan(&opcao)
		if opcao == 4 {
			break
		}

		switch opcao {
		case 1:
			var cod int
			var valor float64
			fmt.Scan(&cod, &valor)
			achou := false
			for i := 0; i < 10; i++ {
				if codigos[i] == cod {
					saldos[i] += valor
					achou = true
					break
				}
			}
			if !achou {
				fmt.Println("Conta não encontrada")
			}
		case 2:
			var cod int
			var valor float64
			fmt.Scan(&cod, &valor)
			achou := false
			for i := 0; i < 10; i++ {
				if codigos[i] == cod {
					achou = true
					if saldos[i] >= valor {
						saldos[i] -= valor
					} else {
						fmt.Println("Saldo insuficiente")
					}
					break
				}
			}
			if !achou {
				fmt.Println("Conta não encontrada")
			}
		case 3:
			var total float64
			for _, s := range saldos {
				total += s
			}
			fmt.Println(total)
		}
	}
}
