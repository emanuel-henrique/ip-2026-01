package main

import "fmt"

// Divisível por 3 e 5
func main() {
	var n int
	var isDivisibleByFive bool
	var isDivisibleByThree bool

	fmt.Print("Digite o número a ser verificado: ")
	fmt.Scan(&n)

	if n % 3 == 0 {
		isDivisibleByThree = true
	}

	if n % 5 == 0 {
		isDivisibleByFive = true
	}

	if isDivisibleByThree && isDivisibleByFive{
		fmt.Print("O NUMERO É DIVISIVEL")
		return
	}
	fmt.Print("O NUMERO NÃO É DIVISIVEL")
}