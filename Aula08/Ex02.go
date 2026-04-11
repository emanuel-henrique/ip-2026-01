package main

import (
	"fmt"
	"sort"
)

func main() {
	var n1, n2, n3 int
	fmt.Print("Digite 3 numeros e te direi o maior deles: ")
	fmt.Scan(&n1, &n2, &n3)
	err, biggest := returnTheBiggest(n1, n2, n3)

	if err {
		fmt.Println("O valores são iguais e não há um maior!")
	} else {

		fmt.Println("O maior número é:", biggest)
	}
}

func returnTheBiggest(n1, n2, n3 int) (bool, int) {
	err := false
	nums := [3]int{n1, n2, n3}

	if n1 == n2 && n2 == n3 {
		err = true
	}

	sort.Ints(nums[:])
	return err, nums[len(nums)-1]
}
