package main
import "fmt"

func main() {
  var count float64
  
  fmt.Print("Quantas notas deseja cadastrar: ")
  fmt.Scan(&count)
  
  var somaNotas float64
  
  for i := 1.0; i <= count; i++ {
      var nota float64
      fmt.Printf("Digite a %.0fº nota: ", i)
      fmt.Scan(&nota)
      
      somaNotas = somaNotas + nota
  }
    
  media := somaNotas / count
  
  fmt.Print("Sua média é: ", media)
}
