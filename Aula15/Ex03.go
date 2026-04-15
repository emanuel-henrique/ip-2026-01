package main

import "fmt"

func main() {
    nums := []int{10, 15, 5, 20, 36}
    
    inverter(nums, 0, len(nums)-1)
    
    fmt.Println(nums)
}

func inverter(arr []int, inicio int, fim int) {
    if inicio >= fim {
        return
    }

    arr[inicio], arr[fim] = arr[fim], arr[inicio]

    inverter(arr, inicio+1, fim-1)
}
