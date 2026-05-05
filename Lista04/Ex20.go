package main

import "fmt"

func main() {
	var jogadas [20]int
	freq := make(map[int]int)

	for i := 0; i < 20; i++ {
		fmt.Scan(&jogadas[i])
		freq[jogadas[i]]++
	}

	fmt.Println("Sorteados:", jogadas)
	for face, f := range freq {
		fmt.Printf("Face %d: %d vezes\n", face, f)
	}
}
