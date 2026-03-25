package main
import "fmt"

// IMPAR OU PAR?
func main() {
  var num int
  fmt.Print("Digite o numero: ")
  fmt.Scan(&num)
  
  if num%2 == 0 {
    fmt.Printf("O número %d é par!", num)
  } else{
    fmt.Printf("O número %d é impar!", num) 
  }
}
