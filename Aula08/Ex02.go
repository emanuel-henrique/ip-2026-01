package main

import (
	"fmt"
	"sort"
)

func main() {
	var n1, n2, n3 int
	fmt.Print("Digite 3 numeros e te direi o maior deles: ")
	fmt.Scan(&n1, &n2, &n3)
	biggest := returnTheBiggest(n1, n2, n3)

	fmt.Println("O maior número é:", biggest)
}

func returnTheBiggest(n1, n2, n3 int) int {
	nums := [3]int{n1, n2, n3}

	sort.Ints(nums[:])
	return nums[len(nums)-1]
}
