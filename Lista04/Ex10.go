package main
import (
    "fmt"
)

func main() {
    var nums []int
   for i := range 51{
       if i > 0{
           num := fibonacci(i)
           nums = append(nums, num)
       }
   }
   fmt.Print(nums)
}

func fibonacci(n int) int {
    if n <= 1 {
        return n
    }
    a, b := 0, 1
    for i := 2; i <= n; i++ {
        a, b = b, a+b
    }
    return b
}
