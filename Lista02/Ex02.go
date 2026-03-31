package main
import "fmt"

// POSITIVO, NEGATIVO ou NULO
func main() {
  var num int
  fmt.Print("Digite o numero: ")
  fmt.Scan(&num)
  
  if num >= 0 {
    fmt.Printf("O número %d é positivo!", num)
  } else if num < 0{
    fmt.Printf("O número %d é negativo!", num) 
  } else {
      fmt.Printf("O número %d é nulo!", num) 
  }
}
