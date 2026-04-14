package main

import "fmt"

func main() {
	var s float64 = 0

	for i := 0; i < 20; i++ {
		var fat float64 = 1
		for j := 1; j <= i; j++ {
			fat *= float64(j)
		}
		s += (100.0 - float64(i)) / fat
	}

	fmt.Printf("Soma dos 20 termos: %.4f\n", s)
}
