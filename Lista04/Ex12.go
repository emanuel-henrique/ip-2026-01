package main

import "fmt"

func main() {
	var notas [15]int
	frequenciaAbsoluta := make(map[int]int)

	for i := 0; i < 15; i++ {
		fmt.Scan(&notas[i])
		frequenciaAbsoluta[notas[i]]++
	}

	fmt.Println("Nota\tAbsoluta\tRelativa")
	for i := 0; i <= 10; i++ {
		abs := frequenciaAbsoluta[i]
		rel := float64(abs) / 15.0
		fmt.Printf("%d\t%d\t\t%.2f\n", i, abs, rel)
	}
}
