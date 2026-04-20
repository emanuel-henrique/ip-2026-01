package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%-10s | %-20s\n", "Raio (cm)", "Volume (cm³)")
	fmt.Println("---------------------------------")

	for r := 0.0; r <= 20.0; r += 0.5 {
		volume := (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
		fmt.Printf("%-10.2f | %-20.4f\n", r, volume)
	}
}
