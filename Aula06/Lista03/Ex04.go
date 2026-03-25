package main
import "fmt"
import "math"

// SOMA
func main() {
  var num float64
  fmt.Print("Digite o primeiro numero: ")
  fmt.Scan(&num)
  
  if num >= 0 {
      raiz := math.Sqrt(num)
      fmt.Printf(".2f", raiz)
  } else {
      fmt.Print(float64(num * num))
  }
}
