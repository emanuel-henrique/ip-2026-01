package main

import "fmt"

func main() {
	var num [10]int
	var divis [5]int

	for i := 0; i < 10; i++ { fmt.Scan(&num[i]) }
	for i := 0; i < 5; i++ { fmt.Scan(&divis[i]) }

	for _, n := range num {
		fmt.Printf("Número %d:\n", n)
		for i, d := range divis {
			if d != 0 && n%d == 0 {
				fmt.Printf(" Divisível por %d na posição %d\n", d, i)
			}
		}
	}
}
