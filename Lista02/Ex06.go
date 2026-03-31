package main
import "fmt"

func main() {
  var a,b int
  fmt.Print("Digite o primeiro numero: ")
  fmt.Scan(&a)
  
  fmt.Print("Digite o segundo numero: ")
  fmt.Scan(&b)
  
  if a%b == 0 {
      fmt.Print("Estes numeros se dividem")
  } else {
      fmt.Print("Estes numeros não se dividem")
  }
}
