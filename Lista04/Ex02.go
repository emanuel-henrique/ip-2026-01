package main

import "fmt"

func main() {
	fmt.Print("Digite 10 números inteiros: ")
	var nums1 []int

	for range 10{
		var num int
		fmt.Scan(&num)
		nums1 = append(nums1, num)
	}

	fmt.Print("Digite 5 números inteiros: ")
	var nums2 []int

	for range 5{
		var num int
		fmt.Scan(&num)
		nums2 = append(nums2, num)
	}

	var vetorResultante1 []int
	var vetorResultante2 []int

	for i := 0; i < len(nums1); i++{
		if nums1[i]%2 == 0{
			sum := nums1[i]
			for _,n := range nums2{
				sum += n
			}
			vetorResultante1 = append(vetorResultante1, sum)
		} else {
			sum := nums1[i]
			for _,n := range nums2{
				sum += n
			}
			vetorResultante2 = append(vetorResultante2, sum)
		}
	}

	fmt.Println(vetorResultante1)
	fmt.Print(vetorResultante2)
}
