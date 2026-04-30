package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	count := 10
	var nums []int
	var positions []int
	
	for i := range count {
		num := rand.IntN(100)
		if num > 50{
			positions = append(positions, i)
		}
		nums = append(nums, num)
	}

  fmt.Println("Array: ", nums)

	if len(nums) > 0{
		for _,n := range positions{
			fmt.Printf("%d - Posição: %d\n", nums[n], n)
		}
	} else {
			fmt.Print("Não há nenhum número superior à 50!")
	}
}
