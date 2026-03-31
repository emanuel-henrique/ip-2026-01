package main

import f "fmt"

func main () {

	var n [5]int
	var i, soma int

	for i = 0; i < 5 ; i++ {
		f.Println("insira o ", i + 1, "º valor inteiro")
		f.Scan(&n[i])
	}
	for i = 0; i < 5; i++{
       soma = soma + n[i]
	}
	f.Println("A soma dos valores é: ", soma)
}
