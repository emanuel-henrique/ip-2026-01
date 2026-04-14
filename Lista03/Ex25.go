package main

import "fmt"

func main() {
	var s float64 = 0
	num := 1.0
	denBase := 15
	i := 0

	for denBase >= 1 {
		termo := num / float64(denBase*denBase)
		if i%2 == 0 {
			s += termo
		} else {
			s -= termo
		}
		num *= 2
		denBase -= 1
		i++
	}

	fmt.Printf("Valor de S: %.4f\n", s)
}
