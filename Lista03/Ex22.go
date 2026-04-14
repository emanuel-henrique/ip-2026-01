package main

import "fmt"

func main() {
	var s float64 = 0
	for i := 1; i <= 37; i++ {
		s += float64((38-i)*(39-i)) / float64(i)
	}
	fmt.Printf("Soma S: %.2f\n", s)
}
