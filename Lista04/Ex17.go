package main

import "fmt"

func ePrimo(n int) bool {
	if n < 2 { return false }
	for i := 2; i*i <= n; i++ {
		if n%i == 0 { return false }
	}
	return true
}

func main() {
	var v [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&v[i])
		if ePrimo(v[i]) {
			fmt.Printf("Valor: %d, Posição: %d\n", v[i], i)
		}
	}
}
