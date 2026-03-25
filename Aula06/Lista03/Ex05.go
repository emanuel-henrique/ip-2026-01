package main
import "fmt"

// DIVISIVEL POR 5
func main() {
  var num int
  fmt.Print("Digite o primeiro numero: ")
  fmt.Scan(&num)
  
  if num%5 == 0 {
      fmt.Print("Este numero é divisivel por 5")
  } else {
      fmt.Print("Este numero não é divisivel por 5")
  }
}
