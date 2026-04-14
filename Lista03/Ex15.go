package main

import "fmt"

func main() {
	fmt.Println("Primeiros 6 termos da série:")
	for i := 1; i <= 6; i++ {
		fmt.Print(i*i, " ")
	}
}
