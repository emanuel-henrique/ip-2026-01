package main

import (
	"fmt"
	"math"
)

func main() {
	var b [100]float64
	for i := 0; i < 100; i++ {
		fmt.Scan(&b[i])
	}

	var s float64
	for i := 0; i < 50; i++ {
		s += math.Pow(b[i]-b[99-i], 3)
	}

	fmt.Println(s)
}
