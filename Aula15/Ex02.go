package main
import "fmt"

func main() {
    nums := []int{10,20,36}
    fmt.Println(somar(nums, len(nums)-1))
}

func somar(arr []int, n int) int {
    if n == 0 {
        return arr[n]
    }   

    return arr[n] + somar(arr, n-1)
}
