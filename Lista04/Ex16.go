package main

import "fmt"

func main() {
	var idades [50]int
	contagem := make(map[int]int)
	for i := 0; i < 50; i++ {
		fmt.Scan(&idades[i])
		contagem[idades[i]]++
	}

	moda, max := 0, 0
	for idade, freq := range contagem {
		if freq > max {
			max = freq
			moda = idade
		}
	}
	fmt.Println(moda)
}
