package main

import f "fmt"

func main() {

	var n [10]float64
	var n_inv [10]float64
	var i int

	for i = 0; i < 10; i++ {
		f.Println("insira o ", i+1, "º valor inteiro")
		f.Scan(&n[i])
	}
	for i = 0; i < 10; i++ {
		n_inv[i] = n[9-i]
	}

	f.Print("A ordem inversa é: ", n_inv)
}
