package main

import "fmt"

func main() {
	var codigo int
	var v [10]float64

	fmt.Scan(&codigo)
	if codigo == 0 { return }

	for i := 0; i < 10; i++ { fmt.Scan(&v[i]) }

	if codigo == 1 {
		fmt.Println(v)
	} else if codigo == 2 {
		for i := 9; i >= 0; i-- {
			fmt.Print(v[i], " ")
		}
		fmt.Println()
	}
}
