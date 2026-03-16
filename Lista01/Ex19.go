package main

import "fmt"

func main() {
	var h, m, s int

	fmt.Print("Horas: ")
	fmt.Scan(&h)

	fmt.Print("Minutos: ")
	fmt.Scan(&m)

	fmt.Print("Segundos: ")
	fmt.Scan(&s)


	fmt.Printf("O TEMPO EM SEGUNDOS E = %d", (h*60*60) + (m*60) + s)
}