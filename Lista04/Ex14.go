package main

import "fmt"

func main() {
	var v1, v2 [10]int
	var resultante [20]int

	for i := 0; i < 10; i++ { fmt.Scan(&v1[i]) }
	for i := 0; i < 10; i++ { fmt.Scan(&v2[i]) }

	for i := 0; i < 10; i++ {
		resultante[i*2] = v1[i]
		resultante[i*2+1] = v2[i]
	}

	fmt.Println(resultante)
}
