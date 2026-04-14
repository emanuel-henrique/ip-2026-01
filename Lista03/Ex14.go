package main

import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Digite N1 e N2: ")
	fmt.Scan(&n1, &n2)
	if n1 > n2 {
		n1, n2 = n2, n1
	}
	fmt.Println("Pares entre os números:")
	for i := n1 + 1; i < n2; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}
