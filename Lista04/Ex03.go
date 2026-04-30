package main

import "fmt"

func main() {
	fmt.Print("Digite 10 números inteiros: ")
	var nums []int

	for range 10{
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}
	sum := 0

	var pares []int
	var impares []int

	for _,n := range nums{
		if n%2 == 0 {	
			pares = append(pares, n)
			sum += n
		} else {
			impares = append(impares, n)
		}
	}

	fmt.Println("Pares: ", pares)
	fmt.Println("Soma dos Pares: ", sum)
	fmt.Println("Impares: ", impares)
	fmt.Print("Quantidade de Impares: ", len(impares))
}
