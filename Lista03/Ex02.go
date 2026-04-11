package main

import "fmt"

func main() {
	var soma int
	var count int
	for i := 50; i <= 70; i++ {
		if i%2 == 0 {
			soma += i
			count++
		}
	}

	fmt.Println("Soma: ", soma)
	fmt.Print("Media: ", soma/count)
}