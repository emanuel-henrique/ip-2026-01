package main
import (
    "fmt"
    "math/rand/v2"
    "math"
)

func main() {
    var nums []float64
    var raizes []float64
    
    for range 15{
        num := float64(rand.IntN(100))
        nums = append(nums, num)
        raiz := math.Sqrt(num)
        raizes = append(raizes, raiz)
    }
    
    fmt.Println("Numeros: ", nums)
    fmt.Printf("Raizes: %.1f", raizes)
}
