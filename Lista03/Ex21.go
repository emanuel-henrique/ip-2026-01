package main

import "fmt"

func main() {
	var b, n int
	var resultado int64 = 1

	fmt.Print("Informe o valor de b (base >= 2): ")
	fmt.Scan(&b)
	fmt.Print("Informe o valor de n (expoente > 1): ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		resultado *= int64(b)
	}

	fmt.Printf("Resultado: %d\n", resultado)
}
