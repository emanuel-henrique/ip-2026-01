package main

import (
	"fmt"
)

func main() {
	var dia, mes, ano int

	fmt.Print("Digite o dia: ")
	fmt.Scan(&dia)
	if dia > 31 || dia <= 0{
		fmt.Print("Dia Inválido")
		return
	}

	fmt.Print("Digite o mês: ")
	fmt.Scan(&mes)
	if mes > 12 || mes <= 0{
		fmt.Print("Mês Inválido")
		return
	}

	fmt.Print("Digite o ano: ")
	fmt.Scan(&ano)
	if ano < 0{
		fmt.Print("Ano Inválido")
		return
	}

	fmt.Printf("dia %d de %s de %d", dia, nomeMesPT(mes), ano)
}

func nomeMesPT(n int) string {
    meses := []string{
        "Janeiro", "Fevereiro", "Março", "Abril", "Maio", "Junho",
        "Julho", "Agosto", "Setembro", "Outubro", "Novembro", "Dezembro",
    }
    if n < 1 || n > 12 {
        return "Mês inválido"
    }
    return meses[n-1]
}