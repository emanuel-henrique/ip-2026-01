package main

import (
	"fmt"
	"sort"
)

func main() {
	var v [10]int
	for i := 0; i < 10; i++ {
		fmt.Scan(&v[i])
	}
	sort.Ints(v[:])
	fmt.Println(v)
}
