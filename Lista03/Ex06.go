package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número: ")
	fmt.Scan(&num)

	triangular := false
	for i := 1; i*(i+1)*(i+2) <= num; i++ {
		if i*(i+1)*(i+2) == num {
			fmt.Printf("%d é triangular (%d x %d x %d)\n", num, i, i+1, i+2)
			triangular = true
			break
		}
	}

	if !triangular {
		fmt.Println("Não é um número triangular.")
	}
}