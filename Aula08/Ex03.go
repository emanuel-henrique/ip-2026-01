package main

import "fmt"

func main() {
	var nt1, nt2, nt3 float64
	fmt.Print("Digite 3 notas e te direi a média delas: ")
	fmt.Scan(&nt1, &nt2, &nt3)

	fmt.Printf("A média é: %.2f\n", calculateMedia(nt1, nt2, nt3))

}

func calculateMedia(nt1, nt2, nt3 float64) float64 {
	media := (nt1 + nt2 + nt3) / 3
	return media
}
