package main

import "fmt"

func main() {
	var num int

	for {
		fmt.Print("Digite um numero: ")
		fmt.Scan(&num)

		if num <= 0 {
			break
		}

		quadradoPerfeito := false
		for i := 1; i*i <= num; i++ {
			if i*i == num {
				quadradoPerfeito = true
				break
			}
		}

		if quadradoPerfeito {
			fmt.Println("É um quadrado perfeito")
		} else {
			fmt.Println("Não é um quadrado perfeito")
		}
	}
}