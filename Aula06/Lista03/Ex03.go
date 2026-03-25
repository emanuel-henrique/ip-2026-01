package main
import "fmt"

// SOMA
func main() {
  var num1,num2 int
  fmt.Print("Digite o primeiro numero: ")
  fmt.Scan(&num1)
  
  fmt.Print("Digite o segundo numero: ")
  fmt.Scan(&num2)
  
  var sum int
  sum = num1 + num2
  
  if sum > 20{
      fmt.Print("Soma mais a adição de 8 é: ", sum+8)
  } else if sum >= 20{
      fmt.Print("Soma mais a subtração de 5 é: ", sum-5)
  }
}
