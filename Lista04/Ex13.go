package main

import (
	"fmt"
	"sort"
)

type Empregado struct {
	ID    int
	Meses int
}

func main() {
	var lista []Empregado
	for i := 0; i < 100; i++ {
		var id, meses int
		fmt.Scan(&id, &meses)
		if id == 0 && meses == 0 {
			break
		}
		lista = append(lista, Empregado{id, meses})
	}

	sort.Slice(lista, func(i, j int) bool {
		return lista[i].Meses < lista[j].Meses
	})

	for i := 0; i < 3 && i < len(lista); i++ {
		fmt.Println(lista[i].ID)
	}
}
